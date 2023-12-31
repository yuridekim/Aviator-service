package pkg

import (
	"errors"
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
	Get(url string) error
	List(url string) error
	Create(url string, request *types.CreateServerRequest, params []int) (*types.CreateServerResponse, error)
	Update(url string) error
	Delete(url string) error
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
	// url 정의
	url += ""

	// httpRequest 생성
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("Error creating request: ", err)
	}
	/// NCP 헤더 설정
	SetNCPHeader(request, "", "")

	// httpRequest token 설정
	SetAuthToken(request, server.token)

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// 요청 보내기
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error sending request:", err)
	}
	defer response.Body.Close()

	// 결과 반환
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response status code:", err)
	}

	return nil
}

func (server *ServerService) Stop(url string) error {
	// url 정의
	url += ""

	// httpRequest 생성
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("Error creating request: ", err)
	}
	/// NCP 헤더 설정
	SetNCPHeader(request, "", "")

	// httpRequest token 설정
	SetAuthToken(request, server.token)

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// 요청 보내기
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("Error sending request:", err)
	}
	defer response.Body.Close()

	// 결과 반환
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response status code:", err)
	}

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
