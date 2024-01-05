# Choose whatever you want, version >= 1.16
FROM golang:1.21-alpine

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest 
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
