package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_di_template/config"
	"golang.org/x/net/context"
	"os"
	"time"
)

const (
	connectionStringTemplate    = "mongodb://%s:%s@%s/%s"
	connectionStringSrvTemplate = "mongodb+srv://%s:%s@%s/%s"
)

type MongoDB struct {
	ConnectTimeout int      `json:"connect_timeout"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Hosts          []string `json:"hosts"`
	Options        string   `json:"options"`
	Database       string   `json:"database"`
}

func NewMongoDB(c *config.Config) *MongoDB {
	return &MongoDB{
		Username: c.MongoDB.UserName,
		Password: c.MongoDB.Password,
		Hosts:    []string{c.MongoDB.Host},
		Database: c.MongoDB.Database,
	}
}

func (m *MongoDB) Connect() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	env := os.Getenv("APP_ENV")
	connectionStr := connectionStringSrvTemplate
	if env == "development.local" {
		connectionStr = connectionStringTemplate
	}
	uri := fmt.Sprintf(connectionStr, m.Username, m.Password, m.Hosts[0], m.Database)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	client.Database(m.Database)
	return client, err
}
