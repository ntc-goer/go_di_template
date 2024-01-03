package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
	"go_di_template/config"
	"go_di_template/internal/middleware"
	"go_di_template/internal/product"
	"go_di_template/internal/server"
	"go_di_template/pkg/database"
)

func serverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run Server API Service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancelFn := context.WithCancel(context.Background())
			defer cancelFn()

			container := provideCoreDependencies()
			// Invoke API Server
			err := container.Invoke(func(c *server.CoreHTTPServer) {
				c.AddCoreRouter()
				err := c.Start()
				if err != nil {
					return
				}
			})
			if err != nil {
				return err
			}

			// Invoke MongoDB Client -> Connect
			var client *mongo.Client
			err = container.Invoke(func(db *mongo.Database) {
				client = db.Client()
			})
			if err != nil {
				return err
			}
			defer func() {
				if err = client.Disconnect(ctx); err != nil {
					panic(err)
				}
			}()
			return nil
		},
	}
}

func provideCoreDependencies() *dig.Container {
	c := dig.New()

	err := c.Provide(gin.New)
	if err != nil {
		return nil
	}
	c.Provide(server.NewHTTPServer)
	c.Provide(server.NewCoreHTTPServer)
	_ = c.Provide(context.Background)
	c.Provide(func() *config.Config {
		return cfg
	})
	// Package configuration
	c.Provide(middleware.NewMiddleware)
	c.Provide(database.NewMongoDB)
	c.Provide(InitMongoDB)
	c.Provide(config.NewCors)

	// API Configuration
	// Sample
	c.Provide(product.NewRepository)
	c.Provide(product.NewService)
	c.Provide(product.NewHandler)

	return c
}
