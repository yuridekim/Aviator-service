package types

type CreateServerRequest struct {
	RegionCode                        string               `json:"regionCode"`
	MemberServerImageInstanceNo       string               `json:"memberServerImageInstanceNo"`
	ServerImageProductCode            string               `json:"serverImageProductCode"`
	ServerImageNo                     string               `json:"serverImageNo"`
	VpcNo                             string               `json:"vpcNo"`
	SubnetNo                          string               `json:"subnetNo"`
	ServerProductCode                 string               `json:"serverProductCode"`
	ServerSpecCode                    string               `json:"serverSpecCode"`
	IsEncryptedBaseBlockStorageVolume bool                 `json:"isEncryptedBaseBlockStorageVolume"`
	FeeSystemTypeCode                 string               `json:"feeSystemTypeCode"`
	ServerCreateCount                 int                  `json:"serverCreateCount"`
	ServerCreateStartNo               int                  `json:"serverCreateStartNo"`
	ServerName                        string               `json:"serverName"`
	NetworkInterfaceList              NetworkInterface     `json:"networkInterfaceList"`
	PlacementGroupNo                  string               `json:"placementGroupNo"`
	IsProtectServerTermination        bool                 `json:"isProtectServerTermination"`
	ServerDescription                 string               `json:"serverDescription"`
	InitScriptNo                      string               `json:"initScriptNo"`
	LoginKeyName                      string               `json:"loginKeyName"`
	AssociateWithPublicIp             bool                 `json:"associateWithPublicIp"`
	RaidTypeName                      string               `json:"raidTypeName"`
	BlockDevicePartitionList          BlockDevicePartition `json:"blockDevicePartitionList"`
	BlockStorageMappingList           BlockStorageMapping  `json:"blockStorageMappingList"`
	ResponseFormatType                string               `json:"responseFormatType"`
}

type GetServerRequest struct{}

type ListServerRequest struct{}

type UpdateServerRequest struct{}
