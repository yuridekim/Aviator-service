package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"

	types "github.com/cloud-club/Aviator-service/types/server"
)

type ServerService struct {
	accessKey string
	secretKey string
}

//go:generate mockgen -destination=mocks/mock_server.go -package=mocks github.com/cloud-club/Aviator-service/pkg ServerInterface
func NewServerService(accessKey, secretKey string) *ServerService {
	return &ServerService{accessKey: accessKey, secretKey: secretKey}
}

func (server *ServerService) CallApi(url string, request types.RequestInterface) (interface{}, error) {
	// Set url with query parameters
	requestParams := request.RequestString()

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url+requestParams, nil)
	if err != nil {
		return nil, err
	}
	// Set HTTP header for NCP authorization
	SetNCPHeader(req, server.accessKey, server.secretKey)

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

	// fmt.Println("request:", requestParams)
	// fmt.Println(string(responseByteData))

	responseStruct, err := request.MapResponse(responseByteData)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return responseStruct, err
}
