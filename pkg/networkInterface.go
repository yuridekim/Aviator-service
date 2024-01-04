package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cloud-club/Aviator-service/types/auth"
	serverType "github.com/cloud-club/Aviator-service/types/server"
)

type NetworkInterfaceService struct {
	KeyService *auth.KeyService
}

func NewNetworkInterfaceService(keyService *auth.KeyService) NetworkInterface {
	return &NetworkInterfaceService{KeyService: keyService}
}

type NetworkInterface interface {
	Get(url string) (*serverType.NetworkInterfaceList, error)
}

func (networkInterface *NetworkInterfaceService) Get(url string) (*serverType.NetworkInterfaceList, error) {
	// Set url with query parameters
	// However, there is no must required query parameters for this API, so we will comment this line right now.
	//requestParams := serverType.CreateRequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set HTTP header for NCP authorization
	SetNCPHeader(req, networkInterface.KeyService.GetAccessKey(), networkInterface.KeyService.GetSecretKey())

	// Make the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// Read the response body
	responseByteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Convert the response body to the struct
	var nifl *serverType.NetworkInterfaceList
	responseInterface, err := serverType.MapResponse(responseByteData, &nifl)
	if err != nil {
		return nil, err
	}

	// interface{} to *serverType.NetworkInterfaceList
	responseStruct := responseInterface.(**serverType.NetworkInterfaceList)

	return *responseStruct, nil
}
