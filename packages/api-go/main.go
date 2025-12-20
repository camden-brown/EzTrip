package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"travel-app/api-go/db"
	"travel-app/api-go/graph"
	"travel-app/api-go/migrations"
	"travel-app/api-go/seeds"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
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

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	
	// Get allowed origins from environment variable
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins != "" {
		corsConfig.AllowOrigins = strings.Split(allowedOrigins, ",")
	} else {
		// Default: allow all origins in development (use specific origins in production)
		corsConfig.AllowAllOrigins = true
	}
	
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	
	router.Use(cors.New(corsConfig))

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
