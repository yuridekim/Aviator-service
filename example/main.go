package example

import (
	"fmt"
	"log"

	server "github.com/cloud-club/Aviator-service/types/server"
)

// main function
func Main() {
	// Create MyServer
	exampleServer := NewExampleServer()
	myserver := NewMyServer(exampleServer)

	myserver_response, myserver_err := myserver.Server.Create(
		"https://ncloud.apigw.ntruss.com/vserver/v2/",
		&server.CreateServerRequest{
			ServerImageProductCode: "SW.VSVR.OS.LNX64.CNTOS.0703.B050",
			VpcNo:                  "***04",
			SubnetNo:               "***43",
			NetworkInterfaceOrder:  1,
		},
	)

	if myserver_err != nil {
		log.Fatalf("Failed to create server instance: %v", myserver_err)
	}

	fmt.Println(myserver_response)
}
