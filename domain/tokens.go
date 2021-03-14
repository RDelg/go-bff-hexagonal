package domain

// TokenClaims represents a token response
type TokenClaims struct {
	AccessToken string `json:"access_token"`
	IssuedAt    string `json:"issued_at"`
	ExpiresIn   string `json:"expires_in"`
}
