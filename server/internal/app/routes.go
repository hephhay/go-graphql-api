package app

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/hephhay/go-graphql/graphql/generated"
	"github.com/hephhay/go-graphql/internal/resolver"
	"net/http"
	"os"
	"io"
)

func NewRoute(resolver *resolver.UserResolver) *gin.Engine {
	f, _ := os.Create("logs/app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.Use(GinContextToContextMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/query", graphqlHandler(resolver))
	r.GET("/", playgroundHandler())

	return r
}

func graphqlHandler(resolver *resolver.UserResolver) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
