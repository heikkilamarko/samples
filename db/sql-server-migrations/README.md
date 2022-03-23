# SQL Server Migrations

## Create `.env` file

**File contents:**

```text
SA_PASSWORD={SA_PASSWORD}
```

SQL Server password policy requirements: **The password must be at least 8 characters.**

## Start SQL Server

```bash
> docker compose up --build -d sql-server
```

## Create database

```bash
> SA_PASSWORD={SA_PASSWORD} ./create_database.sh
```

## Run migrations

```bash
> docker compose build migrate
> docker compose run migrate
```
