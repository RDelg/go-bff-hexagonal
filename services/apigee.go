package services

import (
	"bff/domain"
	"bff/domain/apperrors"
	"log"
)

// ApigeeService contains the methods to acess apigee
type ApigeeService struct {
	ProxyAuthPort domain.ProxyAuthPort
	environment   string
}

// GetAccessToken validates the id token jwt string
// It returns the user extract from the IDTokenCustomClaims
func (s *ApigeeService) GetAccessToken() (string, error) {
	claims, err := s.ProxyAuthPort.GetAccessToken()
	if err != nil {
		log.Printf("Error getting the access token - Error: %v\n", err)
		return "", apperrors.NewAuthorization("Error getting the access token")
	}
	return claims.AccessToken, nil
}

//GetEnvironment returns the environment where to point the service
func (s *ApigeeService) GetEnvironment() (string, error) {
	return s.environment, nil
}

// NewApigeeService returns a new ApigeeService
func NewApigeeService(adapter domain.ProxyAuthPort, environment string) *ApigeeService {
	return &ApigeeService{ProxyAuthPort: adapter, environment: environment}
}
