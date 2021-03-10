package services

// TokenService defines methods the handler layer expects to interact
// with in regards to producing the Apigee access_token
type TokenService interface {
	GetAccessToken() (string, error)
}
