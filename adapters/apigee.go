package adapters

import (
	"bff/domain"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", a.clientID)
	data.Set("client_secret", a.secret)
	data.Set("scope", "")

	req, err := http.NewRequest(http.MethodPost, a.authEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()

	// We read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	// Populate claims from body's data
	var claims domain.ApigeeTokenClaims
	json.Unmarshal([]byte(string(body)), &claims)
	return &claims, nil
}
