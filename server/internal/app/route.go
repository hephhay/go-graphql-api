package app

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/hephhay/go-graphql/graphql/generated"
	"github.com/hephhay/go-graphql/internal/resolver"
)

func NewRoute(resolver *resolver.UserResolver) *gin.Engine {
	r := gin.Default()
	r.Use(GinContextToContextMiddleware())
	r.Use(LoggerToFile())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/query", graphqlHandler(resolver))
	r.GET("/", playgroundHandler())

	return r
}

func graphqlHandler(resolver *resolver.UserResolver) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	h.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return errors.New("Internal server error!")
	})

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
