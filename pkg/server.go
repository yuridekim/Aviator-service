package pkg

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	serverType "github.com/cloud-club/Aviator-service/types/server"
)

type ServerService struct {
	token string
}

//go:generate mockgen -destination=mocks/mock_server.go -package=mocks github.com/cloud-club/Aviator-service/pkg ServerInterface
type ServerInterface interface {
	Get(url string) error
	List(url string) error
	Create(url string, request *serverType.CreateServerRequest) (*serverType.CreateServerResponse, error)
	Update(url string) error
	Delete(url string) error
}

func NewServerService(token string) ServerInterface {
	return &ServerService{token: token}
}

func (server *ServerService) GetToken() string {
	return server.token
}

func (server *ServerService) Create(url string, request *serverType.CreateServerRequest) (*serverType.CreateServerResponse, error) {
	// Set url with query parameters
	requestParams := serverType.CreateRequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, "45x3qDmooHFxwJywHbbK", "xUFTKEw2POsYl5AgBSxf4K2ZJm1JHJ51KHN5BDK8")

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *serverType.CreateServerResponse
	responseInterface, err := serverType.MapResponse(responseByteData, &csr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// interface{} 타입으로 변환된 responseInterface를 다시 CreateServerResponse 타입으로 변환
	responseStruct := responseInterface.(*serverType.CreateServerResponse)

	return responseStruct, err
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
