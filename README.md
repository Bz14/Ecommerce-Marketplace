# ECOMMERCE MARKETPLACE

# Database Schema Migrations – guzoye-db

We manage PostgreSQL schema changes (tables, functions, procedures, sequences, indexes, etc.) using **dbmate** inside **Docker**.  
This ensures:

- Consistent behavior across developers and environments
- No local installation of dbmate, pg_dump, or PostgreSQL client tools required
- Clean, incremental Git history
- Easy rollback in development
- Reproducible schema from any commit

## Folder Structure

migrations-repo/
├── db/
│ ├── migrations/ # ← all versioned .sql migration files live here
│ │ ├── 20260202100000_initial_baseline.sql
│ │ ├── 20260202114500_create_test_table.sql
│ │ └── ...
│ └── schema.sql # ← auto-generated current full schema
├── docker-compose.yml # ← defines the dbmate service
├── .env # ← database credentials
├── .env.example # ← template for DATABASE_URL
├── .gitignore
└── README.md

## Prerequisites

1. **Docker** must be installed and running
   - Windows/macOS → [Docker Desktop](https://www.docker.com/products/docker-desktop/)
   - Linux → Docker Engine + docker-compose

2. Git clone the repo

```bash
git clone https://github.com/your-org/migrations.git
cd migrations
```

## One-time Setup

1. Copy the example environment file and fill in with the database credentials i.e username, password, host, port and databasename

```bash
  cp .env.example .env
```

- Open .env and replace the placeholder:

```bash
  DATABASE_URL=postgresql://USER_NAME:PASSWORD@HOST:PORT/DATABASE_NAME?sslmode=require
```

2. Verify Docker Compose works

```bash
  docker compose version
```

## Common Commands

All commands are run from the repository root using docker compose run.

- Check connection & migration status: Shows applied vs pending migrations

```bash
  docker compose run --rm dbmate status
```

- Dump current schema: Creates or updates `./db/schema.sql`

```bash
 docker compose run --rm dbmate dump
```

- Create new migration: Creates a new migration file in ./db/migrations/ on your machine

```bash
 docker compose run --rm dbmate new migration_name
```

- Apply all pending migration: Runs `--migrate:up` sections

```bash
 docker compose run --rm dbmate up
```

- Rollback the last migration: Runs `--migrate:down`

```bash
docker compose run --rm dbmate rollback
```

```bash
docker compose run --rm dbmate down
```

- See all dbmate options

```bash
docker compose run --rm dbmate --help
```

## Steps for migration

1. Create a new migration: A new file appears: db/migrations/20260202....../create_test_table.sql

```bash
docker compose run --rm dbmate new create_test_table
```

2. Edit the migration file: Open it in your editor and write the change

- migrate:up: The new migration to be applied
- migrate:down: The rollback point

```bash
- migrate:up
CREATE TABLE public.test(id SERIAL PRIMARY_KEY)

-- migrate:down
DROP TABLE public.test
```

3. Apply the migration to the remote database

```bash
docker compose run --rm dbmate up
```

4. (Optional) Refresh the full schema reference

```bash
docker compose run --rm dbmate dump
```

5. Commit the change

```bash
git checkout -b your branch
git add db/migrations/
git add db/schema.sql
git commit -m "Your commit message"
git push
```

6. Test rollback (only in development environment)

```bash
docker compose run --rm dbmate rollback
```

Then re-apply if needed:

```bash
docker compose run --rm dbmate up
```
