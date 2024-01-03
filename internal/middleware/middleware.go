package middleware

import "go_di_template/config"

type Middleware struct {
	Config *config.Config
}

func NewMiddleware(c *config.Config) *Middleware {
	return &Middleware{
		Config: c,
	}
}
