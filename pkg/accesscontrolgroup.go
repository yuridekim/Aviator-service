package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cloud-club/Aviator-service/types/auth"
	serverType "github.com/cloud-club/Aviator-service/types/server"
)

type AccessControlGroupService struct {
	KeyService *auth.KeyService
}

func NewAccessControlGroupService(keyService *auth.KeyService) AccessControlGroupInterface {
	return &AccessControlGroupService{KeyService: keyService}
}

type AccessControlGroupInterface interface {
	Get(url string) (*serverType.AccessControlGroupList, error)
}

func (acg *AccessControlGroupService) Get(url string) (*serverType.AccessControlGroupList, error) {
	// Set url with query parameters
	// However, there is no must required query parameters for this API, so we will comment this line right now.
	//requestParams := serverType.CreateRequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set HTTP header for NCP authorization
	SetNCPHeader(req, acg.KeyService.GetAccessKey(), acg.KeyService.GetSecretKey())

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
	var acgl *serverType.AccessControlGroupList
	responseInterface, err := serverType.MapResponse(responseByteData, &acgl)
	if err != nil {
		return nil, err
	}

	// interface{} to *serverType.AccessControlGroupList
	responseStruct := responseInterface.(**serverType.AccessControlGroupList)

	return *responseStruct, nil
}
