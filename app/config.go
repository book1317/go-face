package app

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	vn    *viper.Viper
	State string

	Users   Users   `mapstructure:"users"`
	MongoDB MongoDB `mapstructure:"mongo"`
}

type Users struct {
	Enable bool `mapstructure:"enable"`
}

type MongoDB struct {
	ConnectionString string `mapstructure:"connection_string"`
	Client           *mongo.Client
}

func (m *MongoDB) binding() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.ConnectionString))
	if err != nil {
		return err
	}
	if err := client.Connect(context.Background()); err != nil {
		return err
	}

	// check connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	m.Client = client
	return nil
}

func NewConfig(state string) *Config {
	return &Config{State: state}
}

func (c *Config) Init() error {
	name := fmt.Sprintf("config.%s", c.State)

	vn := viper.New()
	vn.AddConfigPath("./configs")
	vn.SetConfigName(name)
	c.vn = vn

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := c.binding(); err != nil {
		return err
	}
	return nil
}

func (c *Config) Init_Test(configPath string) error {
	name := fmt.Sprintf("config.%s", c.State)

	vn := viper.New()
	vn.AddConfigPath(configPath)
	vn.SetConfigName(name)
	c.vn = vn

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := c.binding(); err != nil {
		return err
	}
	return nil
}

func (c *Config) binding() error {
	if err := c.vn.Unmarshal(&c); err != nil {
		return err
	}

	if err := c.MongoDB.binding(); err != nil {
		return err
	}

	return nil
}
