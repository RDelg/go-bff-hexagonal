package services

import (
	"bff/domain"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProxyService contains the methods to acess apigee
type ProxyService struct {
	APIEndpoint string
	ProxyAuth   domain.ProxyAuthPort
}

//Get gets data
func (s *ProxyService) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Hello from http service. url: %v\n", c.Request.URL)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "test": "123"})
		new_url := s.APIEndpoint + c.Request.URL.Path
		log.Println("new_url: ", new_url)
		req, err := http.NewRequest(http.MethodGet, new_url, nil)
		req.Header = c.Request.Header.Clone()

		log.Println("HEADER CLONED: ", req.Header.Clone())
		if err != nil {
			log.Println(err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error sending the request: ", err)
		}

		defer resp.Body.Close()
		rbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		c.String(resp.StatusCode, string(rbody))
	}
}

//Get gets data
func (s *ProxyService) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		new_url := s.APIEndpoint + c.Request.URL.Path
		req, err := http.NewRequest(http.MethodPost, new_url, c.Request.Body)
		if err != nil {
			log.Println(err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error sending the request: ", err)
		}
		defer resp.Body.Close()
		rbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		c.String(resp.StatusCode, string(rbody))
	}
}
