package app

import (
	"net/http"

	"eztrip/api-go/graph"
	"eztrip/api-go/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SetupRoutes(router *gin.Engine, database *gorm.DB) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, HealthResponse{
			Status:  "ok",
			Message: "Travel GraphQL API is running",
		})
	})

	resolver := graph.NewResolver(database)
	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	router.POST("/graphql", func(c *gin.Context) {
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	if gin.Mode() != gin.ReleaseMode {
		playgroundHandler := playground.Handler("GraphQL Playground", "/graphql")
		router.GET("/graphql", func(c *gin.Context) {
			playgroundHandler.ServeHTTP(c.Writer, c.Request)
		})
		logger.Log.WithFields(map[string]interface{}{
			"component": "graphql",
			"path":      "/graphql",
		}).Info("GraphQL Playground enabled")
	}
}
