package domain

// ApigeeTokenClaims represents the Apigee token response
type ApigeeTokenClaims struct {
	AccessToken string `json:"access_token"`
	IssuedAt    string `json:"issued_at"`
	ExpiresIn   string `json:"expires_in"`
}
