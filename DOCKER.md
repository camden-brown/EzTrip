# Docker Setup

## Quick Start

```bash
# Start all services (API + PostgreSQL)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Stop and remove volumes (fresh start)
docker-compose down -v
```

## Access

- **API**: http://localhost:8080
- **GraphQL Playground**: http://localhost:8080/graphql
- **PostgreSQL**: localhost:5432

## Development Workflow

### Build and Run
```bash
# Build and start
docker-compose up --build

# Rebuild only API
docker-compose up --build api

# Run in detached mode
docker-compose up -d
```

### Database Management

```bash
# Access PostgreSQL CLI
docker exec -it eztrip-db psql -U postgres -d eztrip

# Run migrations manually
docker exec -i eztrip-db psql -U postgres -d eztrip < packages/api-go/migrations/001_create_users.sql

# Run seeds
docker exec -i eztrip-db psql -U postgres -d eztrip < packages/api-go/seeds/users.sql

# Backup database
docker exec eztrip-db pg_dump -U postgres eztrip > backup.sql

# Restore database
docker exec -i eztrip-db psql -U postgres -d eztrip < backup.sql
```

### Logs

```bash
# All services
docker-compose logs -f

# API only
docker-compose logs -f api

# PostgreSQL only
docker-compose logs -f postgres
```

### Troubleshooting

```bash
# Restart services
docker-compose restart

# Rebuild from scratch
docker-compose down -v
docker-compose up --build

# Check service status
docker-compose ps

# Access API container shell
docker exec -it eztrip-api sh

# Check API health
curl http://localhost:8080/health
```

## Production Deployment

For production, update `docker-compose.yml`:

1. Set strong passwords in environment variables
2. Use `.env` file for secrets
3. Enable SSL/TLS
4. Configure proper logging
5. Set resource limits
6. Use production-grade PostgreSQL configuration

```yaml
# Example .env file
POSTGRES_PASSWORD=your_secure_password
DB_PASSWORD=your_secure_password
GIN_MODE=release
```

## Architecture

```
┌─────────────────┐
│   Browser       │
└────────┬────────┘
         │ :8080
┌────────▼────────┐      ┌──────────────┐
│   API (Go)      │─────▶│ PostgreSQL   │
│   - GraphQL     │ :5432 │ - Database   │
│   - Gin         │      │ - Migrations │
└─────────────────┘      └──────────────┘
```
