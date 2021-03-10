package domain

// ApigeeTokenClaims represents the ApiGee token response
type ApigeeTokenClaims struct {
	AccessToken string `json:"access_token"`
	IssuedAt    int64  `json:"issued_at"`
	ExpiresIn   int64  `json:"expires_in"`
}
