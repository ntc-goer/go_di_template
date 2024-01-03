package server

import (
	"github.com/gin-gonic/gin"
	"go_di_template/config"
	"go_di_template/internal/product"
	"net/http"
)

type CoreHTTPServer struct {
	*HTTPServer
	Cors    *gin.HandlerFunc
	Config  *config.Config
	Product *product.Handler
}

func NewCoreHTTPServer(httpServer *HTTPServer,
	cors *gin.HandlerFunc,
	pd *product.Handler,
) *CoreHTTPServer {
	return &CoreHTTPServer{
		HTTPServer: httpServer,
		Cors:       cors,
		Product:    pd,
	}
}

func (c *CoreHTTPServer) AddCoreRouter() {
	c.Engine.Use(*c.Cors)

	c.Engine.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Source Code")
	})

	v1 := c.Engine.Group("/v1")

	// Sample
	v1.GET("/product", c.Product.GetAll)
}
