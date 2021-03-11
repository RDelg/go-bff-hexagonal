package adapters

import (
	"bff/domain"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ApigeePort contains all the methods to interactuate with ApiGee
type ApigeePort interface {
	GetAccessToken() (*domain.ApigeeTokenClaims, error)
}

// ApigeeAdapter represents an apigee operator
type ApigeeAdapter struct {
	authEndpoint string
	clientID     string
	secret       string
}

// NewApigeeAdapter returns a reference to a new ApigeeAdapter
func NewApigeeAdapter(authEndpoint, clientID, secret string) (*ApigeeAdapter, error) {
	// TODO: Add validation
	return &ApigeeAdapter{
		authEndpoint: authEndpoint,
		clientID:     clientID,
		secret:       secret,
	}, nil
}

type foo struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

//GetAccessToken returns an Apigee token and its duration
func (a *ApigeeAdapter) GetAccessToken() (*domain.ApigeeTokenClaims, error) {
	data := domain.ApigeeTokenRequest{
		GrandType:    "client_credentials",
		ClientID:     a.clientID,
		ClientSecret: a.secret,
		Scope:        "",
	}

	dataJSON, _ := json.Marshal(&data)

	// fmt.Println("ASDASD", bytes.NewBuffer(dataJSON)) // write response to ResponseWriter (w)

	req, _ := http.NewRequest("POST", a.authEndpoint, bytes.NewBuffer(dataJSON))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Println("ASDASDASDASDDAS", sb)

	return &domain.ApigeeTokenClaims{
		AccessToken: "test_token asd qwe",
		IssuedAt:    1,
		ExpiresIn:   2,
	}, nil
}
