package main

import (
	"log"
	"net/http"
	"os"

	"travel-app/api-go/db"
	"travel-app/api-go/graph"
	"travel-app/api-go/migrations"
	"travel-app/api-go/seeds"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// Initialize database connection
	dbConfig := db.GetConfigFromEnv()
	database, err := db.NewGormDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL with GORM")

	// Run migrations
	if err := migrations.RunMigrations(database); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Run seeds (only in development)
	if os.Getenv("SEED_DATA") == "true" {
		if err := seeds.RunSeeds(database); err != nil {
			log.Printf("Warning: Failed to run seeds: %v", err)
		}
	}

	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, HealthResponse{
			Status:  "ok",
			Message: "Travel GraphQL API is running",
		})
	})

	// Initialize GraphQL resolver with database
	resolver := graph.NewResolver(database)

	// GraphQL handler
	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	playgroundHandler := playground.Handler("GraphQL Playground", "/graphql")

	// GraphQL endpoint
	router.POST("/graphql", func(c *gin.Context) {
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	// GraphQL playground (development UI)
	router.GET("/graphql", func(c *gin.Context) {
		playgroundHandler.ServeHTTP(c.Writer, c.Request)
	})

	router.Run(":8080")
}
