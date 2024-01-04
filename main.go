package main

import (
	"fmt"

	//"time"

	pkg "github.com/cloud-club/Aviator-service/pkg"
	auth "github.com/cloud-club/Aviator-service/types/auth"
	types "github.com/cloud-club/Aviator-service/types/server"
)

func getImageProductList(ncp *pkg.NcpService) {
	ncp.ServerImageProduct = pkg.NewImageProductService(&ncp.Key)
	imageProductList, err := ncp.ServerImageProduct.Get(pkg.API_URL + pkg.GET_SERVER_IMAGE_PRODUCT_LIST_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(imageProductList)
}

func getProductList(ncp *pkg.NcpService) {
	ncp.ServerProduct = pkg.NewProductService(&ncp.Key)

	gpr := &types.GetProductRequest{ServerImageProductCode: "SW.VSVR.OS.LNX64.CNTOS.0703.B050"}

	productList, err := ncp.ServerProduct.Get(pkg.API_URL+pkg.GET_SERVER_PRODUCT_LIST_PATH, gpr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(productList)
}

func getNetworkInterfaceList(ncp *pkg.NcpService) {
	ncp.Network = pkg.NewNetworkInterfaceService(&ncp.Key)
	networkInterfaceList, err := ncp.Network.Get(pkg.API_URL + pkg.GET_NETWORKINTERFACE_LIST_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(networkInterfaceList)
}

func getAccessControlGroupList(ncp *pkg.NcpService) {
	ncp.AccessControlGroup = pkg.NewAccessControlGroupService(&ncp.Key)
	accessControlGroupList, err := ncp.AccessControlGroup.Get(pkg.API_URL + pkg.GET_ACG_LIST_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(accessControlGroupList)
}

func getSubnetList(ncp *pkg.NcpService) {
	ncp.Subnet = pkg.NewSubnetService(&ncp.Key)
	subnetList, err := ncp.Subnet.Get(pkg.VPC_API_URL + pkg.GET_SUBNET_LIST_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(subnetList)
}

func getVpcList(ncp *pkg.NcpService) {
	ncp.Vpc = pkg.NewVpcService(&ncp.Key)
	vpcList, err := ncp.Vpc.Get(pkg.VPC_API_URL + pkg.GET_VPC_LIST)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(vpcList)
}

func create(ncp *pkg.NcpService) {
	csr := &types.CreateServerRequest{ServerImageProductCode: "SW.VSVR.OS.LNX64.CNTOS.0703.B050", VpcNo: "52833", SubnetNo: "120320", NetworkInterfaceOrder: 0, AccessControlGroupNoListN: "148207", ServerProductCode: "SVR.VSVR.HICPU.C002.M004.NET.HDD.B050.G002"}

	response, err := ncp.Server.Create(pkg.API_URL+pkg.CREATE_SERVER_INSTANCE_PATH, csr, []int{1, 1})
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print response
	fmt.Println(response)
}

func stop(ncp *pkg.NcpService) {
	ssr := &types.StopServerRequest{ServerNo: "21741451"}

	stopServerResponse, err := ncp.Server.Stop(pkg.API_URL+pkg.STOP_SERVER_INSTANCE_PATH, ssr)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print response
	fmt.Println(stopServerResponse)
}

func list(ncp *pkg.NcpService) {
	lsr := &types.ListServerRequest{RegionCode: "KR"}

	serverListResponse, err := ncp.Server.List(pkg.API_URL+pkg.GET_SERVER_INSTANCE_PATH, lsr)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print response
	fmt.Println(serverListResponse)
	// fmt.Println("server: ", serverListResponse.ServerInstanceList[0])
	// fmt.Println("status: ", serverListResponse.ServerInstanceList[0].ServerInstanceStatus.Code)
}

func update(ncp *pkg.NcpService) {
	usr := &types.UpdateServerRequest{ServerInstanceNo: "21741451", ServerProductCode: "SVR.VSVR.STAND.C032.M128.NET.HDD.B050.G002"}

	updateServerResponse, err := ncp.Server.Update(pkg.API_URL+pkg.UPDATE_SERVER_INSTANCE_PATH, usr)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print response
	fmt.Println(updateServerResponse)
}

func delete(ncp *pkg.NcpService) {
	dsr := &types.DeleteServerRequest{ServerNo: "21741451"}

	deleteServerResponse, err := ncp.Server.Delete(pkg.API_URL+pkg.DELETE_SERVER_INSTANCE_PATH, dsr)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print response
	fmt.Println(deleteServerResponse)
}

func start(ncp *pkg.NcpService) {
	ssr := &types.StartServerRequest{ServerNo: "21763788"}
	_, err := ncp.Server.Start(pkg.API_URL+pkg.START_SERVER_INSTANCE_PATH, ssr)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	ncp := pkg.NewNcpService("ncp service token")
	ncp.Key = *auth.NewKeyService("6CmrDJ4KaswJ10g25GEP", "OvZ7QHH0Bi3AwGn5rlsD7xoC986bEOiIjdbwMFCo")
	ncp.Server = pkg.NewServerService(&ncp.Key)

	// Check image product list
	getImageProductList(ncp)

	// Check product list
	getProductList(ncp)

	// Get network interface list
	getNetworkInterfaceList(ncp)

	// Get access control group list
	getAccessControlGroupList(ncp)

	// Get subnet list
	getSubnetList(ncp)

	// Get VPC list
	getVpcList(ncp)

	list(ncp)

	// Create server instance
	//create(ncp)

	//6. Get server instance list
	//list(ncp)

	// 7. Stop server instance
	//stop(ncp)
	// list(ncp)

	// Start server instance
	// start(ncp)

	// list(ncp)

	// Update server instance
	// update(ncp)
	// time.Sleep(20 * time.Second)
	// list(ncp)

	// // Terminate server instance
	// delete(ncp)

}
