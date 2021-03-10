package services

import (
	"bff/adapters"
	"bff/domain/apperrors"
	"log"
)

// ApigeetokenService contains the methods to acess apigee
type ApigeetokenService struct {
	apigeePort adapters.ApigeePort
}

// GetAccessToken validates the id token jwt string
// It returns the user extract from the IDTokenCustomClaims
func (s *ApigeetokenService) GetAccessToken() (string, error) {
	claims, err := s.apigeePort.GetAccessToken()
	if err != nil {
		log.Printf("Error getting the access token - Error: %v\n", err)
		return "", apperrors.NewAuthorization("Error getting the access token")
	}
	return claims.AccessToken, nil
}

// NewApigeetokenService returns a new ApigeetokenService
func NewApigeetokenService(adapter adapters.ApigeePort) *ApigeetokenService {
	return &ApigeetokenService{apigeePort: adapter}
}
