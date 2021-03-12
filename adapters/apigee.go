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

// ApigeeProxyAuthAdapter represents an apigee operator
type ApigeeProxyAuthAdapter struct {
	authEndpoint string
	clientID     string
	secret       string
}

// NewApigeeProxyAuthAdapter returns a reference to a new ApigeeProxyAuthAdapter
func NewApigeeProxyAuthAdapter(authEndpoint, clientID, secret string) (*ApigeeProxyAuthAdapter, error) {
	// TODO: Add validation
	return &ApigeeProxyAuthAdapter{
		authEndpoint: authEndpoint,
		clientID:     clientID,
		secret:       secret,
	}, nil
}

//GetAccessToken returns an Apigee token and its duration
func (a *ApigeeProxyAuthAdapter) GetAccessToken() (*domain.ApigeeTokenClaims, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", a.clientID)
	data.Set("client_secret", a.secret)
	data.Set("scope", "")
	req, err := http.NewRequest(http.MethodPost, a.authEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	// We read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Populate claims from body's data
	var claims domain.ApigeeTokenClaims
	json.Unmarshal([]byte(string(body)), &claims)
	return &claims, nil
}
