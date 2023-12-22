# MiniWeb Documentation

## Overview

MiniWeb is a lightweight web application built using Go. It follows a modular structure with a focus on simplicity and flexibility.

## Table of Contents

- [Dependencies](#dependencies)
- [Directory Structure](#directory-structure)
- [Routing](#routing)
- [Database](#database)
- [Environment Variables](#environment-variables)
- [Template Dependencies](#template-dependencies)
- [Getting Started](#getting-started)

## Dependencies

- [Chi](https://github.com/go-chi/chi): A lightweight, idiomatic and composable router for building Go HTTP services.
- [sql-migrate](https://github.com/rubenv/sql-migrate): A database migration tool, inspired by Ruby on Rails migrations.
- [SQLC](https://github.com/kyleconroy/sqlc): Generate type safe Go from SQL.
- [Alpine.js](https://alpinejs.dev/): A rugged, minimal JavaScript framework.
- [Tailwind CSS](https://tailwindcss.com/): A utility-first CSS framework.
- [HTMX](https://htmx.org/): A library for high-level interactions between the browser and server.



## Routing

The routing in MiniWeb is implemented using [Chi](https://github.com/go-chi/chi). For detailed information on routing, refer to the [Chi documentation](https://pkg.go.dev/github.com/go-chi/chi/v5).

## Database

MiniWeb uses [sql-migrate](https://github.com/rubenv/sql-migrate) for database migrations. Additionally, [SQLC](https://github.com/kyleconroy/sqlc) is employed to generate type-safe Go code from SQL.

## Environment Variables

The application uses environment variables for configuration. Key environment variables include:

- `PORT`: The port on which the server listens.
- (Custom variables based on your application needs.)

Refer to the [Env Helper Documentation](#) for more details.

## Template Dependencies

The templates in MiniWeb rely on the following front-end dependencies:

- [Alpine.js](https://alpinejs.dev/): A minimal JavaScript framework for enhanced interactivity.
- [Tailwind CSS](https://tailwindcss.com/): A utility-first CSS framework for styling.
- [HTMX](https://htmx.org/): A library for dynamic HTML updates and interactions.

Ensure you have these dependencies installed and configured in your development environment. Run the following commands for setup:

```bash
npm install
```

## Running the server

```bash
npm run watch
```
Separate terminal
```bash
air
```
