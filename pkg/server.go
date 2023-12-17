package pkg

import (
	"errors"
	"fmt"
	"net/http"
)

type ServerService struct {
	token string
}

//go:generate mockgen -destination=mocks/mock_server.go -package=mocks github.com/cloud-club/Aviator-service/pkg ServerInterface
type ServerInterface interface {
	Get(url string) error
	List(url string) error
	Create(url string, payload interface{}) error
	Update(url string) error
	Delete(url string) error
}

func NewServerService(token string) ServerInterface {
	return &ServerService{token: token}
}

func (server *ServerService) GetToken() string {
	return server.token
}

func (server *ServerService) Create(url string, payload interface{}) error {
	return nil
}

func (server *ServerService) Get(url string) error {
	return nil
}

func (server *ServerService) List(url string) error {
	if len(url) == 0 {
		return errors.New("please input url")
	}
	return nil
}

func (server *ServerService) Delete(url string) error {
	return nil
}

func (server *ServerService) Update(url string) error {
	// url += fmt.Sprintf("?regionCode=%s&serverInstanceNo=%s&serverProductCode=%s", updateParams.regionCode, updateParams.serverInstanceNo, updateParams.serverProductCode)
	url += "temp"

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// Set common headers
	GetCommonHeader(req)

	// Set authorization token
	SetAuthToken(req, server.token)

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}
