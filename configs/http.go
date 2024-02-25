package configs

import (
	"os"
	"strconv"
)

type httpConfig struct {
	port int
}

func HTTPConfigure() *httpConfig {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8000
	}
	return &httpConfig{port: port}
}

func (h *httpConfig) Port() int {
	return h.port
}
