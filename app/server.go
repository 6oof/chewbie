package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/6oof/miniweb-base/app/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/securecookie"
)

type MiniWeb struct {
	Router *chi.Mux
}

func MbinInit() *MiniWeb {

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(compressResponseMiddleware)

	// CSRF protection
	csrfKey := securecookie.GenerateRandomKey(32)
	csrfMiddleware := csrf.Protect(csrfKey, csrf.Secure(false), csrf.CookieName("ccsrf")) // Set Secure to true in production
	r.Use(csrfMiddleware)

	// Add your routes as needed
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	RegisterRoutes(r)
	// Apply CSRF protection to all routes
	miniWeb := &MiniWeb{
		Router: r,
	}

	return miniWeb
}

func compressResponseMiddleware(next http.Handler) http.Handler {
	return handlers.CompressHandler(next)
}

func MbinServe(port string) {
	c := MbinInit()

	// Create a server with timeouts
	server := &http.Server{
		Addr:         port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      c.Router,
	}

	appEnv := helpers.Env("APP_ENV", "development")
	appPort := helpers.Env("PORT", "8080")

	if appEnv == "production" {
		// Listen for syscall signals for process to interrupt/quit
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sig

			// Create a deadline to wait for.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			// Trigger graceful shutdown
			err := server.Shutdown(ctx)
			if err != nil {
				log.Println(err)
			}
		}()

		// Start the server in a goroutine so that it doesn't block.
		go func() {
			err := server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()

		log.Printf("Server running in graceful mode on port %s\n", appPort)

		// Block until an interrupt signal is received.
		<-sig
		log.Println("Shutting down...")

		// Optionally, you could run server.Shutdown in a goroutine
		// and block on <-ctx.Done() if your application should wait
		// for other services to finalize based on context cancellation.
		log.Println("Server gracefully stopped.")
	} else {
		log.Printf("Server running on port %s\n", appPort)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}
