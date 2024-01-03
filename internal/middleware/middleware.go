package middleware

import (
	"github.com/gin-gonic/gin"
	"go_di_template/config"
)

type Middleware struct {
	Config *config.Config
}

func NewMiddleware(c *config.Config) *Middleware {
	return &Middleware{
		Config: c,
	}
}

func (mdw *Middleware) AuthMiddleware(ctx *gin.Context) {
	//TODO implement me
	ctx.Next()
}
