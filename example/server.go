package example

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	server "github.com/cloud-club/Aviator-service/types/server"
)

type MyServer struct {
	Server ExampleServerInterface
}

func NewMyServer(server ExampleServerInterface) *MyServer {
	if server == nil {
		panic("server is nil")
	}
	return &MyServer{Server: server}

}

func (ms *MyServer) CreateServerInstance(url string, request *server.CreateServerRequest) (*server.CreateServerResponse, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}

	response, err := ms.Server.Create(url, request)
	fmt.Printf("MyServer.CreateServerInstance() called")
	return response, err
}

type ExampleServer struct {
}

func NewExampleServer() ExampleServerInterface {
	return &ExampleServer{}
}

type ExampleServerInterface interface {
	Create(url string, request *server.CreateServerRequest) (*server.CreateServerResponse, error)
	Get(url string) error
	List(url string) error
	Update(url string) error
	Delete(url string) error
}

func (client *ExampleServer) Create(url string, request *server.CreateServerRequest) (*server.CreateServerResponse, error) {
	// 테스트용으로 시각화하기 위해 짠 코드라서 실제로는 사용하지 않음!
	// Header에 Content-Type & Token 값 추가했다고 가정
	requestString := createRequestString(request)
	response, err := http.Get(url + requestString)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal(response.Status)
		return nil, err
	}

	responseByteData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var csr server.CreateServerResponse
	responseInterface, err := server.MapResponse(responseByteData, &csr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// interface{} 타입으로 변환된 responseInterface를 다시 CreateServerResponse 타입으로 변환
	responseStruct := responseInterface.(*server.CreateServerResponse)

	return responseStruct, err
}

func (client *ExampleServer) Get(url string) error {
	return nil
}

func (client *ExampleServer) List(url string) error {
	return nil
}

func (client *ExampleServer) Update(url string) error {
	return nil
}

func (client *ExampleServer) Delete(url string) error {
	return nil
}
