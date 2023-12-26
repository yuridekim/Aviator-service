package pkg_test

import (
	"testing"

	"github.com/cloud-club/Aviator-service/mocks"
	"github.com/cloud-club/Aviator-service/pkg"
	server "github.com/cloud-club/Aviator-service/types/server"
	"github.com/stretchr/testify/assert"
)

func TestGetNetworkInterfaceList(t *testing.T) {
	testCases := []struct {
		testName string
		url      string
		response *server.NetworkInterfaceList
		err      error
	}{
		{
			testName: "Success - 성공",
			url:      pkg.API_URL + pkg.GET_NETWORKINTERFACE_LIST_PATH,
			response: &server.NetworkInterfaceList{
				TotalRows: 2,
				NetworkInterfaceList: []server.NetworkInterface{
					{
						NetworkInterfaceNo:     "1",
						NetworkInterfaceName:   "test-01",
						SubnetNo:               "1",
						DeleteOnTermination:    true,
						IsDefault:              true,
						NetworkInterfaceStatus: server.CommonCode{Code: "code", CodeName: "codeName"},
						IP:                     "10.0.0.0",
						MacAddress:             "00:00:00:00:00:00",
					},
					{
						NetworkInterfaceNo:     "2",
						NetworkInterfaceName:   "test-02",
						SubnetNo:               "2",
						DeleteOnTermination:    true,
						IsDefault:              true,
						NetworkInterfaceStatus: server.CommonCode{Code: "code", CodeName: "codeName"},
						IP:                     "10.0.0.1",
						MacAddress:             "00:00:00:00:00:00",
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Helper()
		t.Run(tc.testName, func(t *testing.T) {
			mockServer := &mocks.MockNetworkInterface{}
			mockServer.On("Get", tc.url).Return(tc.response, tc.err).Once()

			response, err := mockServer.Get(tc.url)

			assert.Equal(t, tc.response, response)
			assert.Nil(t, err)
			mockServer.AssertExpectations(t)
		})
	}
}
