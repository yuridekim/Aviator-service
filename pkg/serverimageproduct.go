package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cloud-club/Aviator-service/types/auth"
	serverType "github.com/cloud-club/Aviator-service/types/server"
)

type ImageProductService struct {
	KeyService *auth.KeyService
}

func NewImageProductService(keyService *auth.KeyService) ImageProductInterface {
	return &ImageProductService{KeyService: keyService}
}

type ImageProductInterface interface {
	Get(url string) (*serverType.ProductList, error)
}

func (product *ImageProductService) Get(url string) (*serverType.ProductList, error) {
	// Set url with query parameters
	// However, there is no must required query parameters for this API, so we will comment this line right now.
	//requestParams := serverType.CreateRequestString(request)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set HTTP header for NCP authorization
	SetNCPHeader(req, product.KeyService.GetAccessKey(), product.KeyService.GetSecretKey())

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
	var sl *serverType.ProductList
	responseInterface, err := serverType.MapResponse(responseByteData, &sl)
	if err != nil {
		return nil, err
	}

	// interface{} to *serverType.SubnetList
	responseStruct := responseInterface.(**serverType.ProductList)

	return *responseStruct, nil
}
