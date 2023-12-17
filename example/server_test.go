package example

import (
	"testing"
	"time"

	"github.com/cloud-club/Aviator-service/mocks"
	server "github.com/cloud-club/Aviator-service/types/server"
	"github.com/stretchr/testify/assert"
)

func TestExampleServerCreate(t *testing.T) {
	testCases := []struct {
		testName string
		url      string
		request  *server.CreateServerRequest
		response *server.CreateServerResponse
		err      error
	}{
		{
			testName: "Success - 성공",
			url:      "https://ncloud.apigw.ntruss.com/vserver/v2/",
			request: &server.CreateServerRequest{
				ServerImageProductCode: "SW.VSVR.OS.LNX64.CNTOS.0703.B050",
				VpcNo:                  "***04",
				SubnetNo:               "***43",
				NetworkInterfaceOrder:  1,
			},
			response: &server.CreateServerResponse{
				RequestId:     "e7e7e7e7-7e7e-7e7e-7e7e-7e7e7e7e7e7e",
				ReturnCode:    0,
				ReturnMessage: "success",
				TotalRows:     1,
				ServerInstanceList: []server.ServerInstance{
					{
						ServerInstanceNo:               "***4299",
						ServerName:                     "test-***",
						CpuCount:                       2,
						MemorySize:                     4294967296,
						PlatformType:                   server.CommonCode{Code: "LNX64", CodeName: "Linux 64 Bit"},
						LoginKeyName:                   "test-***",
						ServerInstanceStatus:           server.CommonCode{Code: "INIT", CodeName: "Server initializing"},
						ServerInstanceOperation:        server.CommonCode{Code: "NULL", CodeName: "Server operation null"},
						ServerInstanceStatusName:       "init",
						CreateDate:                     time.Time{},
						Uptime:                         time.Time{},
						ServerImageProductCode:         "SW.VSVR.OS.LNX64.CNTOS.0703.B050",
						ServerProductCode:              "SVR.VSVR.STAND.C002.M004.NET.SSD.B050.G001",
						IsProtectServerTermination:     false,
						ZoneCode:                       "KR-1",
						RegionCode:                     "KR",
						VpcNo:                          "***04",
						SubnetNo:                       "***43",
						NetworkInterfaceNoList:         server.NetworkInterfaceNoList{},
						ServerInstanceType:             server.CommonCode{Code: "SVRSTAND", CodeName: "Server Standard"},
						BaseBlockStorageDiskType:       server.CommonCode{Code: "NET", CodeName: "Network Storage"},
						BaseBlockStorageDiskDetailType: server.CommonCode{Code: "SSD", CodeName: "SSD"},
					},
				},
			},
			err: nil,
		},
		{
			testName: "Failed - 필수 파라미터 누락",
			url:      "https://ncloud.apigw.ntruss.com/vserver/v2/",
			request: &server.CreateServerRequest{
				VpcNo:    "***04",
				SubnetNo: "***43",
			},
			response: &server.CreateServerResponse{
				RequestId:     "e7e7e7e7-7e7e-7e7e-7e7e-7e7e7e7e7e7e",
				ReturnCode:    0,
				ReturnMessage: "Failed",
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Helper()
		t.Run(tc.testName, func(t *testing.T) {
			mockServer := &mocks.MockExampleServerInterface{}
			mockServer.On("Create", tc.url, tc.request).
				Return(tc.response, tc.err).
				Once()

			response, err := mockServer.Create("https://ncloud.apigw.ntruss.com/vserver/v2/", tc.request)

			assert.Nil(t, err, "The error should be nil")
			assert.Equal(t, tc.response, response, "The responses should be equal")
			mockServer.AssertExpectations(t)
		})
	}
}
