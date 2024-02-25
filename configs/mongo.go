package configs

import (
	"os"
	"strconv"
)

type mongoConfig struct {
	mongoUrl     string
	mongoDB      string
	mongoTimeout int
}

func MongoConfigure() *mongoConfig {
	mongoURL := os.Getenv("MONGO_URL")
	mongoDB := os.Getenv("MONGO_DB")
	if mongoURL == "" {
		panic("Please specify a Mongo URL")
	}
	if mongoDB == "" {
		panic("Please specify a Mongo DB")
	}
	mongoTimeout, err := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	if err != nil {
		panic(err)
	}
	return &mongoConfig{
		mongoUrl:     mongoURL,
		mongoDB:      mongoDB,
		mongoTimeout: mongoTimeout,
	}
}

func (c *mongoConfig) URL() string {
	return c.mongoUrl
}

func (c *mongoConfig) Database() string {
	return c.mongoDB
}

func (c *mongoConfig) Timeout() int {
	return c.mongoTimeout
}
