package pkg

import (
	"errors"
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
	return nil
}
