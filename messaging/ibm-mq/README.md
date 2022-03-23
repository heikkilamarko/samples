# IBM MQ Docker Compose

## Create `.env` file

**File contents:**

```text
MQ_ADMIN_PASSWORD={MQ_ADMIN_PASSWORD}
```

## Start IBM MQ

```bash
> docker compose up --build -d
```

**IBM MQ Web UI:** https://localhost:9443

- username: `admin`
- password: `{MQ_ADMIN_PASSWORD}`
