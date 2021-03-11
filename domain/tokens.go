package domain

// ApigeeTokenRequest represents the Apigee data that is required
// when getting the access token
type ApigeeTokenRequest struct {
	GrandType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

// ApigeeTokenClaims represents the Apigee token response
type ApigeeTokenClaims struct {
	AccessToken string `json:"access_token"`
	IssuedAt    int64  `json:"issued_at"`
	ExpiresIn   int64  `json:"expires_in"`
}
