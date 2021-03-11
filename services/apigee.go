package services

import (
	"bff/adapters"
	"bff/domain/apperrors"
	"log"
)

// ApigeeService contains the methods to acess apigee
type ApigeeService struct {
	apigeePort adapters.ApigeePort
}

// GetAccessToken validates the id token jwt string
// It returns the user extract from the IDTokenCustomClaims
func (s *ApigeeService) GetAccessToken() (string, error) {
	claims, err := s.apigeePort.GetAccessToken()
	if err != nil {
		log.Printf("Error getting the access token - Error: %v\n", err)
		return "", apperrors.NewAuthorization("Error getting the access token")
	}
	return claims.AccessToken, nil
}

// NewApigeeService returns a new ApigeeService
func NewApigeeService(adapter adapters.ApigeePort) *ApigeeService {
	return &ApigeeService{apigeePort: adapter}
}
