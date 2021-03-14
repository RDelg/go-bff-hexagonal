package main

import (
	"bff/adapters"
	"bff/controllers"
	"bff/domain"
	"bff/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func inject() (*gin.Engine, error) {
	log.Println("Injecting data sources")

	endpoint := os.Getenv("API_ENDPOINT")
	apigeeClient := os.Getenv("APIGEE_CLIENT_ID")
	apigeeSecret := os.Getenv("APIGEE_SECRET")
	environment := os.Getenv("API_ENV")

	apigeeProxy, err := adapters.NewApigeeAdapter(endpoint, apigeeClient, apigeeSecret, environment, "/authorization")

	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}

	ProxyService := services.ProxyService{Proxy: apigeeProxy}

	router := gin.Default()

	controllers.ConfigRouter(&domain.Config{
		R:            router,
		ProxyService: &ProxyService,
	})
	return router, nil
}

// have func(authPath string) github.com/gin-gonic/gin.HandlerFunc,
// want func() github.com/gin-gonic/gin.HandlerFunc
