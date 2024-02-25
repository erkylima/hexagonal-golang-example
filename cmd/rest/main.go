package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/erkylima/hexagonal/hexagonal/configs"
	"github.com/erkylima/hexagonal/hexagonal/internal/beneficiary"
	"github.com/erkylima/hexagonal/hexagonal/pkg/api"
	"github.com/erkylima/hexagonal/hexagonal/pkg/database/mongodb"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// repo <- service -> serializer  -> http

func main() {
	repo := repo()
	service := beneficiary.NewBeneficiaryService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{name}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)

	go func() {
		port := httpPort()
		fmt.Printf("Listening on port :%s\n", port)
		errs <- http.ListenAndServe(port, r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	HttPConfig := configs.HTTPConfigure()
	port := HttPConfig.Port()
	return fmt.Sprintf(":%d", port)
}

func repo() beneficiary.BeneficiaryRepository {
	mongoConfig := configs.MongoConfigure()
	repo, err := mongodb.NewMongoRepository(mongoConfig.URL(), mongoConfig.Database(), mongoConfig.Timeout())
	if err != nil {
		print("Error creating", err.Error())
	}
	return repo
}
