# Database Migrations & Seeds

## GORM Auto-Migrations

This project uses GORM's AutoMigrate feature for database schema management. Migrations run automatically when the API starts.

### How It Works

**Migrations** - `migrations/migrations.go`
- Runs on every API startup
- Creates tables, columns, and indexes based on model structs
- **Safe**: Won't delete existing columns or change types
- Add new models to `RunMigrations()` function

**Seeds** - `seeds/seeds.go`
- Controlled by `SEED_DATA` environment variable
- Only runs if database is empty
- Creates sample data for development

### Local Development

```bash
# Migrations run automatically on startup
npx nx serve api-go

# Enable seeding
SEED_DATA=true npx nx serve api-go
```

### Docker

```bash
# With seeds (development)
docker-compose up

# Without seeds (set SEED_DATA=false in docker-compose.yml)
docker-compose up
```

### Adding New Models

1. Create your model struct with GORM tags
2. Add it to `migrations/migrations.go`:
   ```go
   err := db.AutoMigrate(
       &user.User{},
       &yourpackage.YourModel{}, // Add here
   )
   ```
3. Restart the API - table will be created automatically

