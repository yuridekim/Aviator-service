package pkg_test

import (
	"testing"
	"time"

	"github.com/cloud-club/Aviator-service/mocks"
	"github.com/cloud-club/Aviator-service/pkg"
	server "github.com/cloud-club/Aviator-service/types/server"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	ncp := pkg.NewNcpService("ncp service token")
	ncp.Server = pkg.NewServerService("ncp server token")

	// given
	tests := []struct {
		name          string
		url           string
		payload       interface{}
		expectedError string
		expectedData  interface{}
	}{
		{
			name:          "성공",
			url:           "http://localhost:8080",
			payload:       nil,
			expectedError: "",
			expectedData:  nil,
		},
		{
			name:          "(실패) url을 입력 안함",
			url:           "",
			payload:       nil,
			expectedError: "please input url",
			expectedData:  nil,
		},
	}

	for i := range tests {
		t.Logf("%s : running scenario %d", tests[i].name, i)
		t.Run(tests[i].name, func(t *testing.T) {
			err := ncp.Server.List(tests[i].url)
			if err != nil {
				if err.Error() != tests[i].expectedError {
					t.Fatalf("expected error : %v, got : %v", tests[i].expectedError, err)
				}
			} else {
				if tests[i].expectedError != "" {
					t.Fatalf("expected error : %v, got : %v", tests[i].expectedError, err)
				}
			}
		})
	}

}

func TestCreate(t *testing.T) {

	var path = "createServerInstances"

	testCases := []struct {
		testName string
		url      string
		request  *server.CreateServerRequest
		response *server.CreateServerResponse
		err      error
	}{
		{
			testName: "Success - 성공",
			url:      pkg.API_URL + path,
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
			url:      pkg.API_URL + path,
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
			mockServer := &mocks.MockServerInterface{}
			mockServer.On("Create", tc.url, tc.request).
				Return(tc.response, tc.err).
				Once()

			response, err := mockServer.Create(pkg.API_URL+path, tc.request)

			assert.Nil(t, err, "The error should be nil")
			assert.Equal(t, tc.response, response, "The responses should be equal")
			mockServer.AssertExpectations(t)
		})
	}
}
