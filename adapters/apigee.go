package adapters

import (
	"bff/domain"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ApigeeAdapter contains the values and server to communicate
// with the Apigee proxy
type ApigeeAdapter struct {
	endpoint     string
	clientID     string
	clientSecret string
	environment  string
	authPath     string
	httpClient   *http.Client
}

// NewApigeeAdapter returns a reference to a new ApigeeAdapter
func NewApigeeAdapter(endpoint, clientID, clientSecret, environment, authPath string) (*ApigeeAdapter, error) {
	return &ApigeeAdapter{
		endpoint:     endpoint,
		clientID:     clientID,
		clientSecret: clientSecret,
		environment:  environment,
		httpClient:   &http.Client{},
		authPath:     authPath,
	}, nil
}

// Auth authenticates againts the Apigee server and returns the access token
func (a *ApigeeAdapter) Auth() (*domain.TokenClaims, error) {

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", a.clientID)
	data.Set("client_secret", a.clientSecret)
	data.Set("scope", "")

	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, body, err := a.DoRequest(http.MethodPost, a.endpoint+a.authPath, &header, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var claims domain.TokenClaims
	json.Unmarshal([]byte(string(body)), &claims)
	return &claims, nil
}

// DoRequest method
func (a *ApigeeAdapter) DoRequest(method, url string, header *http.Header, body io.Reader) (int, []byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return 0, nil, err
	}
	if header != nil {
		req.Header = *header
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		log.Println("Error sending the request: ", err)
		return 0, nil, err
	}

	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, nil, err
	}
	return resp.StatusCode, rbody, nil
}

// GetEnv returns the adapter environment
func (a *ApigeeAdapter) GetEnv() string {
	return a.environment
}
