package adapters

import "bff/domain"

// ApigeePort contains all the methods to interactuate with ApiGee
type ApigeePort interface {
	GetAccessToken() (*domain.ApigeeTokenClaims, error)
}

// ApigeeAdapter represents an apigee operator
type ApigeeAdapter struct {
	clientID string
	secret   string
}

// NewApigeeAdapter returns a reference to a new ApigeeAdapter
func NewApigeeAdapter(clientID, secret string) (*ApigeeAdapter, error) {
	// TODO: Add validation
	return &ApigeeAdapter{
		clientID: clientID,
		secret:   secret,
	}, nil
}

//GetAccessToken returns an Apigee token and its duration
func (a *ApigeeAdapter) GetAccessToken() (*domain.ApigeeTokenClaims, error) {
	return &domain.ApigeeTokenClaims{
		AccessToken: "test_token asd qwe",
		IssuedAt:    1,
		ExpiresIn:   2,
	}, nil
}
