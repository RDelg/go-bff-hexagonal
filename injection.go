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

	basePath := os.Getenv("API_ENDPOINT")
	authPath := basePath + "/authorization"

	apigeeAdapter, err := adapters.NewApigeeAdapter(authPath, os.Getenv("APIGEE_CLIENT_ID"), os.Getenv("APIGEE_SECRET"))

	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}
	apigeeService := services.NewApigeeService(apigeeAdapter)
	httpService := services.HTTPService{}

	router := gin.Default()

	controllers.ConfigRouter(&domain.Config{
		R:                   router,
		InternalAuthService: apigeeService,
		HTTPService:         &httpService,
		APIEndpoint:         basePath,
	})
	return router, nil
}
