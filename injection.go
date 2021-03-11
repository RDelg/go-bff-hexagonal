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

	apigeeAdapter, err := adapters.NewApigeeAdapter(os.Getenv("APIGEE_CLIENT_ID"), os.Getenv("APIGEE_SECRET"))

	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %w", err)
	}
	apigeeService := services.NewApigeeService(apigeeAdapter)
	httpService := services.HTTPService{}

	router := gin.Default()

	controllers.ConfigRouter(&domain.RestConfig{
		R:                   router,
		InternalAuthService: apigeeService,
		HTTP:                &httpService,
	})
	return router, nil
}
