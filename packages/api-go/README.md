# Travel API - Go GraphQL Backend

A GraphQL API built with Go, Gin framework, and gqlgen.

## Prerequisites

- Go 1.23 or higher
- Nx CLI
- Docker & Docker Compose (for containerized setup)

## Quick Start

### With Docker (Recommended)

```bash
# From workspace root - starts API + PostgreSQL
docker-compose up -d

# View logs
docker-compose logs -f api

# API available at http://localhost:8080
```

See [DOCKER.md](../../DOCKER.md) for full Docker documentation.

### Local Development

```bash
# Start PostgreSQL (if not using Docker)
# See migrations/README.md for setup

# From the workspace root
npx nx serve api-go

# Or directly in the package
cd packages/api-go
go run main.go
```

The API will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /health` - Check API status

### GraphQL Endpoint
- `POST /graphql` - GraphQL API endpoint
- `GET /graphql` - GraphQL Playground (interactive UI)

## GraphQL Schema

Define your schema in `graph/schema.graphqls`, then regenerate code:

```bash
cd packages/api-go
go run github.com/99designs/gqlgen generate
```

### Current Schema

```graphql
type Query {
  hello: String!
}
```

### Example Query

```graphql
query {
  hello
}
```

## Using GraphQL Playground

1. Open your browser to `http://localhost:8080/graphql`
2. Use the interactive UI to explore the schema and run queries
3. Documentation is auto-generated from your schema

## Development Commands

```bash
# Run the server
npx nx serve api-go

# Build the application
npx nx build api-go

# Run tests
npx nx test api-go

# Run linter
npx nx lint api-go

# Tidy dependencies
npx nx run api-go:tidy

# Regenerate GraphQL code after schema changes
cd packages/api-go
go run github.com/99designs/gqlgen generate
```

## Project Structure

```
api-go/
├── graph/
│   ├── model/              # Generated GraphQL models
│   ├── generated.go        # Generated GraphQL server code
│   ├── resolver.go         # Resolver with data storage
│   ├── schema.resolvers.go # Resolver implementations
│   └── schema.graphqls     # GraphQL schema definition
├── main.go                 # Server entry point with Gin
├── go.mod                  # Go module dependencies
└── project.json            # Nx project configuration
```

## Adding Your Schema

1. Edit `graph/schema.graphqls` to define your types
2. Run `go run github.com/99designs/gqlgen generate`
3. Implement the generated resolver methods in `graph/schema.resolvers.go`
4. Add any data/dependencies to `graph/resolver.go`
