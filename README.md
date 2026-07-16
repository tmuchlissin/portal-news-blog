# Portal News Blog API

**Portal News Blog API** is a backend service for a news/blog portal built with **Go**.

- 🧱 **Architecture:** uses **Hexagonal Architecture** / **Ports and Adapters**.
- 🎯 **Core idea:** keep **business logic** independent from external details.
- 🔌 **Adapters:** HTTP handlers, database repositories, configuration, storage, and third-party integrations stay outside the core layer.
- 🧪 **Benefit:** easier to test, extend, and refactor as the application grows.

## Project Structure

```text
portal-news-blog/
├── cmd/                         # CLI command layer using Cobra.
│   ├── root.go                  # Root Cobra command, config flag, and Viper initialization.
│   └── start.go                 # Start command that calls the application bootstrap.
├── config/                      # Application configuration and database setup.
│   ├── config.go                # Maps environment values into application config structs.
│   └── database.go              # Opens the PostgreSQL connection with GORM.
├── database/                    # Database-related files.
│   └── migrations/              # SQL migration files for schema changes.
├── internal/                    # Private application code.
│   ├── adapter/                 # External interface implementations.
│   │   ├── cloudflare/          # Adapter for Cloudflare/storage integration.
│   │   ├── handler/             # HTTP handlers/controllers.
│   │   └── repository/          # Database access layer.
│   ├── app/                     # Application bootstrap and dependency initialization.
│   │   └── app.go               # Initializes application dependencies.
│   └── core/                    # Core business logic.
│       ├── domain/              # Domain layer.
│       │   └── model/           # Domain/entity models.
│       └── service/             # Business logic layer.
├── lib/                         # Shared helper packages.
│   ├── conf/                    # Configuration helper package.
│   │   └── conf.go              # Config helper utilities.
│   └── jwt/                     # JWT helper package.
├── .env                         # Local environment file. Do not commit real secrets.
├── main.go                      # Application entry point.
├── go.mod
├── go.sum
└── README.md
```

Folder details:

- `cmd/` contains CLI command definitions. `root.go` owns the root command and config loading, while `start.go` runs the app.
- `config/` contains the app configuration structs and PostgreSQL connection setup.
- `database/migrations/` contains versioned SQL files. Each migration has an `.up.sql` file for applying changes and a `.down.sql` file for rollback.
- `internal/adapter/` is for outer-layer implementations such as HTTP handlers, database repositories, and Cloudflare/storage integration.
- `internal/app/` is the current application bootstrap layer.
- `internal/core/` is for application core code: domain models in `domain/model/` and business logic in `service/`.
- `lib/` contains shared helper packages such as config helpers and JWT utilities.
- `main.go` is the executable entry point used by `go run main.go`.

## Requirements

- Go `1.23` or a compatible version.
- PostgreSQL database. Neon works as long as SSL is enabled.
- `golang-migrate` for SQL migrations.

## Environment

Create a `.env` file in the project root.

```env
APP_ENV=development
APP_PORT=8080

DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=password
DATABASE_NAME=portal_news_blog
DATABASE_SSL_MODE=require
DATABASE_MAX_OPEN_CONNECTIONS=10
DATABASE_MAX_IDLE_CONNECTIONS=5

JWT_SECRET_KEY=change-this-secret
JWT_ISSUER=portal-news-blog
```

Notes:

- `DATABASE_SSL_MODE` defaults to `require` in code when it is not set.
- Neon requires `sslmode=require`.
- The config loader still supports the old typo `DATABSE_NAME` as a fallback, but `DATABASE_NAME` is the recommended key.
- Do not commit real credentials from `.env`.

## Local Setup

Clone the repository and enter the project directory:

```bash
git clone <repository-url>
cd portal-news-blog
```

Install Go dependencies:

```bash
go mod tidy
```

Create `.env` manually from the Environment section above, or copy it from `.env.example` if that file is added later.

## Run The Application

Run the default command:

```bash
go run main.go
```

You can also run the explicit Cobra command:

```bash
go run main.go start
```

The app currently initializes configuration and opens a PostgreSQL connection. If the database connection succeeds, the command exits without starting an HTTP server yet.

## Database Migrations

This project uses `golang-migrate` to create and run SQL migration files in `database/migrations`.

Install `golang-migrate`:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Make sure the `migrate` binary is available:

```bash
export PATH="$PATH:$HOME/go/bin"
migrate -version
```

To make this permanent for new terminal sessions:

```bash
echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.bashrc
source ~/.bashrc
```

Create a new migration:

```bash
migrate create -ext sql -dir database/migrations -seq create_users_table
```

That command creates two files:

```text
database/migrations/000001_create_users_table.up.sql
database/migrations/000001_create_users_table.down.sql
```

Load `.env` into the current shell:

```bash
set -a
source .env
set +a
```

If your `.env` still uses the old typo `DATABSE_NAME`, map it to the correct variable name before running migrations:

```bash
export DATABASE_NAME="$DATABSE_NAME"
```

Run all pending migrations:

```bash
migrate -path database/migrations \
  -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=require" \
  up
```

Check the current migration version:

```bash
migrate -path database/migrations \
  -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=require" \
  version
```

Rollback the latest migration:

```bash
migrate -path database/migrations \
  -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=require" \
  down 1
```

## Dirty Migration Recovery

If a migration fails, `golang-migrate` may mark the database as dirty, for example:

```text
Dirty database version 3
```

Fix the SQL file first, then force the database version back to the last successful migration. For example, if migration `000003` failed, the last successful version is `2`:

```bash
migrate -path database/migrations \
  -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=require" \
  force 2
```

Then rerun migrations:

```bash
migrate -path database/migrations \
  -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=require" \
  up
```

`force` does not run SQL and does not rollback data. It only updates the migration version stored in the database, so use it only after you understand which migration failed and the SQL file has been fixed.

## Notes

- 🔐 **JWT** is used for stateless authentication by carrying signed user claims.
- ⚙️ **[Viper](https://github.com/spf13/viper)** is used to load configuration from `.env` and environment variables.
- 🐍 **[Cobra](https://github.com/spf13/cobra)** is used to define CLI commands such as the root command and `start`.
- 🗄️ **[golang-migrate](https://github.com/golang-migrate/migrate)** is used to create, apply, rollback, and track SQL migrations.

## License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE).
