package routes

import (
	"bff/domain"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type routes struct {
	Method string
	Path   string
}

func loadRoutes(file string) ([]routes, error) {
	csvFile, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening the file: %v\n", err)
		return nil, err
	}
	log.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Printf("Error reading the file: %v\n", err)
		return nil, err
	}
	routesArr := make([]routes, len(csvLines))

	for i, line := range csvLines {
		routesArr[i] = routes{
			Method: line[0],
			Path:   line[1],
		}
	}
	return routesArr, nil
}

// AddRoutes adds all the routes
func AddRoutes(config *domain.Config, rg *gin.RouterGroup) {
	routes, err := loadRoutes("./resources/routes.csv")
	if err != nil {
		log.Fatalf("Error loading the routes: %v\n", err)
	}
	for _, route := range routes {
		url := route.Path
		log.Println(route.Method, ": ", url)
		switch route.Method {
		case "POST":
			rg.POST(url, config.ProxyService.Do(http.MethodPost))
			log.Println("OK")
		case "GET":
			rg.GET(url, config.ProxyService.Do(http.MethodGet))
			log.Println("OK")
		default:
			log.Fatalln("METHOD NOT IMPLEMENTED")
		}

	}
}
