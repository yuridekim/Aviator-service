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

type ServerInterface interface {
	List(url string, request *types.ListServerRequest) (*types.ListServerResponse, error)
	Create(url string, request *types.CreateServerRequest, params []int) (*types.CreateServerResponse, error)
	Update(url string, request *types.UpdateServerRequest) (*types.UpdateServerResponse, error)
	Stop(url string, request *types.StopServerRequest) (*types.StopServerResponse, error)
	Delete(url string, request *types.DeleteServerRequest) (*types.DeleteServerResponse, error)
}

func (server *ServerService) List(url string, request *types.ListServerRequest) (*types.ListServerResponse, error) {
	requestParams := types.RequestString(request)

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
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	println(string(responseByteData))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.ListServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.ListServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Create(url string, request *types.CreateServerRequest, params []int) (*types.CreateServerResponse, error) {
	// Set url with query parameters
	networkInterfaceIndex := params[0]
	acgNoList := params[1]
	requestParams := types.CreateRequestString(request, networkInterfaceIndex, acgNoList)

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
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	println(string(responseByteData))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// fmt.Println("request:", requestParams)
	// fmt.Println(string(responseByteData))
	var csr *types.CreateServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// interface{} 타입으로 변환된 responseInterface를 다시 CreateServerResponse 타입으로 변환
	responseStruct := responseInterface.(**types.CreateServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Update(url string, request *types.UpdateServerRequest) (*types.UpdateServerResponse, error) {
	requestParams := types.RequestString(request)

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
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	println(string(responseByteData))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.UpdateServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.UpdateServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Stop(url string, request *types.StopServerRequest) (*types.StopServerResponse, error) {
	requestParams := types.RequestString(request)

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
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	println(string(responseByteData))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.StopServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.StopServerResponse)

	return *responseStruct, err
}

func (server *ServerService) Delete(url string, request *types.DeleteServerRequest) (*types.DeleteServerResponse, error) {
	requestParams := types.RequestString(request)

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
		// Read the response body and show the body message in error.
		responseByteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", responseByteData)
	}

	responseByteData, err := io.ReadAll(resp.Body)
	println(string(responseByteData))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr *types.DeleteServerResponse
	responseInterface, err := types.MapResponse(responseByteData, &csr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseStruct := responseInterface.(**types.DeleteServerResponse)

	return *responseStruct, err
}
