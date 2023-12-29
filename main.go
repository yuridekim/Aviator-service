package main

import (
	"fmt"

	pkg "github.com/cloud-club/Aviator-service/pkg"
	types "github.com/cloud-club/Aviator-service/types/server"
)

func main() {
	vserver := pkg.NewServerService("b7e6Eq3fmVMGKBCCSLbi", "S6ewbCjNSCk5kQLRDHvqXDGPqTUDwDn2LLhmIKma")
	serverNo := test_list(0, vserver)
	test_stop(serverNo, vserver)
	// fmt.Println("return code of response: %s", response.)
	// ssr := types.StopServerRequest{ServerNo: "blabla"}
	// vserver.CallApi(pkg.API_URL+pkg.STOP_SERVER_INSTANCE_PATH, ssr)
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

func test_stop(serverNo string, vserver *pkg.ServerService) { // server index for choosing which server to target
	ssr := types.StopServerRequest{ServerNo: serverNo}
	response, err := vserver.CallApi(pkg.API_URL+pkg.STOP_SERVER_INSTANCE_PATH, ssr)
	responseStruct := response.(*types.StopServerResponse)
	fmt.Println("Error for stopping: ", err)
	fmt.Println("responsestruct for stopping:", responseStruct)
}
