package main

import (
	"fmt"
	"time"

	pkg "github.com/cloud-club/Aviator-service/pkg"
	types "github.com/cloud-club/Aviator-service/types/server"
)

func main() {
	vserver := pkg.NewServerService("access key", "secret key")
	serverNo := test_list(0, vserver)

	stopReturn := test_stop(serverNo, vserver)
	fmt.Println("Is stop successful: ", stopReturn)

	time.Sleep(10) //wait for stop to take place

	deleteReturn := test_delete(serverNo, vserver)
	fmt.Println("Is deletion(termination) successful: ", deleteReturn)
}

func test_list(server_index int, vserver *pkg.ServerService) string { // server index for choosing which server to target
	gsr := types.ListServerRequest{RegionCode: "KR"}
	response, err := vserver.CallApi(pkg.API_URL+pkg.GET_SERVER_INSTANCE_PATH, gsr)
	responseStruct := response.(*types.ListServerResponse)
	serverNo := responseStruct.ServerInstanceList[server_index].ServerInstanceNo
	fmt.Println("Error for listing:", err)
	fmt.Println("responsestruct for listing: ", responseStruct)
	fmt.Println("serverno: ", serverNo)

	return serverNo
}

func test_stop(serverNo string, vserver *pkg.ServerService) string { // server index for choosing which server to target
	ssr := types.StopServerRequest{ServerNo: serverNo}
	response, err := vserver.CallApi(pkg.API_URL+pkg.STOP_SERVER_INSTANCE_PATH, ssr)
	responseStruct := response.(*types.StopServerResponse)
	fmt.Println("Error for stopping: ", err)
	fmt.Println("responsestruct for stopping:", responseStruct)

	return responseStruct.ReturnMessage
}

func test_delete(serverNo string, vserver *pkg.ServerService) string { // server index for choosing which server to target
	dsr := types.DeleteServerRequest{ServerNo: serverNo}
	response, err := vserver.CallApi(pkg.API_URL+pkg.DELETE_SERVER_INSTANCE_PATH, dsr)
	responseStruct := response.(*types.DeleteServerResponse)
	fmt.Println("Error for stopping: ", err)
	fmt.Println("responsestruct for stopping:", responseStruct)

	return responseStruct.ReturnMessage
}

func test_update(serverNo, serverProductCode string, vserver *pkg.ServerService) string { // server index for choosing which server to target
	dsr := types.UpdateServerRequest{ServerInstanceNo: serverNo, ServerProductCode: serverProductCode}
	response, err := vserver.CallApi(pkg.API_URL+pkg.UPDATE_SERVER_INSTANCE_PATH, dsr)
	responseStruct := response.(*types.UpdateServerResponse)
	fmt.Println("Error for stopping: ", err)
	fmt.Println("responsestruct for stopping:", responseStruct)

	return responseStruct.ReturnMessage
}
