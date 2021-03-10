package main

import (
	"bff/adapters"
	"bff/controllers"
	"bff/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func inject() (*gin.Engine, error) {
	log.Println("Injecting data sources")

	apigeeAdapter, err := adapters.CreateApigeeAdapter(os.Getenv("APIGEE_CLIENT_ID"), os.Getenv("APIGEE_SECRET"))

	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}
	apigeeService := services.NewApigeetokenService(apigeeAdapter)

	router := gin.Default()

	controllers.ConfigRouter(&controllers.RestConfig{
		R:             router,
		ApigeeService: apigeeService,
	})
	return router, nil
}
