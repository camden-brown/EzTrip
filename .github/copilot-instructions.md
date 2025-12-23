# GitHub Copilot Instructions - Travel App

## Project Overview

This is an Nx monorepo containing a Go GraphQL backend API for a travel application. The API uses modern patterns with separation of concerns, dependency injection, and idiomatic Go practices.

## Tech Stack

- **Monorepo**: Nx 22.3.1
- **Backend**: Go 1.24 with Gin web framework
- **API**: GraphQL via gqlgen v0.17.85
- **Database**: PostgreSQL 16 with GORM ORM
- **Containerization**: Docker and Docker Compose

## Architecture Principles

### Separation of Concerns

Each domain model lives in its own package with the following structure:

```
packages/api-go/
├── user/
│   ├── user.go          # Domain model with GORM tags
│   └── service.go       # Business logic (CRUD operations)
├── graph/
│   ├── schema.graphqls  # GraphQL schema definitions
│   ├── resolver.go      # Dependency injection container
│   └── schema.resolvers.go  # GraphQL resolver implementations
├── db/
│   └── db.go           # Database connection management
├── migrations/
│   └── migrations.go   # GORM AutoMigrate
└── seeds/
    ├── seeds.go        # Seed orchestration
    └── user_seeds.go   # Model-specific seeds
```

### Dependency Injection

- Resolvers receive dependencies through `NewResolver(db *gorm.DB)`
- Services are instantiated with database connection: `user.NewService(db)`
- No global state or singletons

## Code Conventions

### Domain Models

- Use GORM tags for database mapping
- Include soft delete support with `gorm.DeletedAt`
- Use `github.com/google/uuid` for IDs
- Example:

```go
type User struct {
    ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    FirstName string         `gorm:"column:first_name;not null"`
    Email     string         `gorm:"column:email;uniqueIndex;not null"`
    CreatedAt time.Time      `gorm:"column:created_at"`
    UpdatedAt time.Time      `gorm:"column:updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
```

### Services

- Name pattern: `type Service struct { db *gorm.DB }`
- Constructor: `func NewService(db *gorm.DB) *Service`
- Methods take `context.Context` as first parameter
- CRUD operations: `GetAll(ctx)`, `GetByID(ctx, id)`, `Create(ctx, entity)`, `Update(ctx, id, updates)`, `Delete(ctx, id)`
- Use GORM methods: `Find()`, `First()`, `Create()`, `Updates()`, `Delete()`

### GraphQL Schema

- Use custom domain models (not generated models)
- Configure in `gqlgen.yml`:

```yaml
models:
  User:
    model: eztrip/api-go/user.User
```

- Run `npx nx run api-go:generate` after schema changes
- Keep resolver implementations in `schema.resolvers.go`

### Database Migrations

- Use GORM AutoMigrate in `migrations/migrations.go`
- Add new models to `RunMigrations()`: `db.AutoMigrate(&user.User{}, &post.Post{})`
- AutoMigrate is safe - it won't delete columns

### Seeds

- Each model gets its own seed file: `seeds/{model}_seeds.go`
- Include idempotency checks (count existing records)
- Orchestrate in `seeds/seeds.go` with `RunSeeds(db)`
- Seeds run during Docker setup via `seed-init` service, not on server startup
- Run manually with: `npx nx run api-go:seed`

### Error Handling

- Return errors, don't panic
- Log errors with context: `log.Printf("Failed to create user: %v", err)`
- Use GORM error checking: `if err := db.First(&user, id).Error; err != nil`

### Naming Conventions

- Go structs: PascalCase
- Database columns: snake_case (via GORM tags)
- GraphQL fields: camelCase (gqlgen handles conversion)
- Package names: lowercase, singular (user, not users)
- Service receivers: single lowercase letter (s \*Service)

## Adding New Models

Follow this pattern when adding a new model (e.g., Post):

1. **Create domain model**: `packages/api-go/post/post.go`

   - Add GORM tags, UUID primary key, timestamps, soft delete

2. **Create service**: `packages/api-go/post/service.go`

   - Implement CRUD operations with GORM

3. **Update GraphQL schema**: `packages/api-go/graph/schema.graphqls`

   - Define types, queries, mutations

4. **Configure gqlgen**: `packages/api-go/gqlgen.yml`

   - Map GraphQL type to Go model: `Post: { model: eztrip/api-go/post.Post }`

5. **Generate and implement**:

   - Run: `npx nx run api-go:generate`
   - Implement resolvers in `schema.resolvers.go`
   - Inject service in `resolver.go`

6. **Add migration**: `packages/api-go/migrations/migrations.go`

   - Add to AutoMigrate: `db.AutoMigrate(&post.Post{})`

7. **Create seeds**: `packages/api-go/seeds/post_seeds.go`
   - Add orchestration call in `seeds/seeds.go`

## Environment Variables

Development (docker-compose.yml):

- `DB_HOST=postgres`
- `DB_PORT=5432`
- `DB_USER=postgres`
- `DB_PASSWORD=postgres`
- `DB_NAME=eztrip`
- `GIN_MODE=debug` (development mode)

Production: Use `GIN_MODE=release`

## Nx Commands

- `npx nx run api-go:serve` - Run development server
- `npx nx run api-go:build` - Build binary
- `npx nx run api-go:test` - Run tests
- `npx nx run api-go:lint` - Run go vet
- `npx nx run api-go:tidy` - Run go mod tidy
- `npx nx run api-go:generate` - Generate GraphQL code
- `npx nx run api-go:seed` - Run database seeds manually

## Docker

- Development: `docker-compose up -d` (seeds run automatically via seed-init service)
- Rebuild: `docker-compose up -d --build`
- Logs: `docker-compose logs -f api`
- Seed logs: `docker-compose logs seed-init`
- Stop: `docker-compose down`

Multi-stage Dockerfile builds optimized binary from source. The `seed-init` service runs once during startup to populate the database.

## Important Notes

- **Never use generated models** - always use custom domain models in your own packages
- **Always run gqlgen generate** after schema changes
- **Context is required** - all service methods take `context.Context`
- **Use GORM methods** - no raw SQL unless absolutely necessary
- **Idempotent seeds** - always check if data exists before inserting
- **Seeds run during Docker setup** - not on server startup (use `seed-init` service)
- **UUID primary keys** - use `type:uuid;default:gen_random_uuid()`
- **Soft deletes** - include `DeletedAt gorm.DeletedAt` in models

## Frontend (Web App) Guidelines

### Angular Component Conventions

- **Use OnPush change detection** - always add `changeDetection: ChangeDetectionStrategy.OnPush` for child components
  - Import: `import { ChangeDetectionStrategy } from '@angular/core';`
  - Improves performance by reducing change detection cycles
  - Required for all presentational/child components
- **Standalone components** - use standalone: true for all new components
- **Smart/Dumb pattern** - separate container (smart) from presentational (dumb) components

### CSS/SCSS Conventions

- **Use relative units** - prefer `rem`, `em`, `%`, `vh`, `vw` over `px`
  - `rem` for font sizes, spacing, and most dimensions
  - `em` for component-relative sizing
  - `%` or viewport units for responsive layouts
- **Use SCSS variables** - import from `styles/_variables.scss` for colors
- Base font size is 16px, so `1rem = 16px`

## GraphQL Playground

Access at: http://localhost:8080/graphql

Example query:

```graphql
query {
  users {
    id
    firstName
    lastName
    email
  }
}
```
