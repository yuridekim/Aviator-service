package types

// 필수가 아닌 필드(필수 여부: No)는 주석 처리 해두었음.
// 필요할 때 주석 해제
// 필수가 아닌 필드 중 (필수 여부: Conditional)는 주석 처리 안 했음.
// 필요할 때 주석 처리
type CreateServerRequest struct {
	//RegionCode                        string               `json:"regionCode"`
	MemberServerImageInstanceNo string `json:"memberServerImageInstanceNo"` // Conditional
	ServerImageProductCode      string `json:"serverImageProductCode"`      // Conditional
	ServerImageNo               string `json:"serverImageNo"`               // Conditional
	VpcNo                       string `json:"vpcNo"`
	SubnetNo                    string `json:"subnetNo"`
	//ServerProductCode                 string               `json:"serverProductCode"`
	ServerSpecCode string `json:"serverSpecCode"` // Conditional
	//IsEncryptedBaseBlockStorageVolume bool                 `json:"isEncryptedBaseBlockStorageVolume"`
	//FeeSystemTypeCode                 string               `json:"feeSystemTypeCode"`
	//ServerCreateCount                 int                  `json:"serverCreateCount"`
	//ServerCreateStartNo               int                  `json:"serverCreateStartNo"`
	//ServerName                        string               `json:"serverName"`

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// NetworkInterfaceList.N의 값들
	// NetworkInterfaceList를 먼저 호출한 뒤, N 번째 NetworkInterface 정보에서 필요한 부분들 추출해서 사용
	NetworkInterfaceOrder int `json:"networkInterfaceList"`
	//NetworkInterfaceNo         string               `json:"networkInterfaceNo"`
	//NetworkInterfaceSubnetNo  string               `json:"networkInterfaceSubnetNo"`
	//NetworkInterfaceIp 	  string               `json:"networkInterfaceIp"`
	//accessControlGroupNoListN []string             `json:"accessControlGroupNoList"`
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	// PlacementGroupNo           string               `json:"placementGroupNo"`
	// IsProtectServerTermination bool                 `json:"isProtectServerTermination"`
	// ServerDescription          string               `json:"serverDescription"`
	// InitScriptNo               string               `json:"initScriptNo"`
	// LoginKeyName               string               `json:"loginKeyName"`
	// AssociateWithPublicIp      bool                 `json:"associateWithPublicIp"`
	RaidTypeName string `json:"raidTypeName"` // Conditional

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// BlockDevicePartitionList.N의 값들
	// BlockDevicePartitionList를 먼저 호출한 뒤, N 번째 BlockDevicePartition 정보에서 필요한 부분들 추출해서 사용
	//BlockDevicePartitionMountPoint string               `json:"blockDevicePartitionMountPoint"`
	//BlockDevicePartitionSize       string               `json:"blockDevicePartitionSize"`
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// BlockStorageMappingList.N의 값들
	// BlockStorageMappingList를 먼저 호출한 뒤, N 번째 BlockStorageMapping 정보에서 필요한 부분들 추출해서 사용
	//BlockStorageMappingOrder                      int                  `json:"blockStorageMappingList"`
	//BlockStorageMappingSnapshotInstanceNo         string               `json:"blockStorageMappingSnapshotInstanceNo"`
	//BlockStorageMappingBlockStorageSize           string               `json:"blockStorageMappingBlockStorageSize"`
	//BlockStorageMappingBlockStorageName           string               `json:"blockStorageMappingBlockStorageName"`
	//BlockStorageMappingBlockStorageVolumeTypeCode string               `json:"blockStorageMappingBlockStorageVolumeTypeCode"`
	//BlockStorageMappingEncrypted                  string               `json:"blockStorageMappingEncrypted"`
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	//ResponseFormatType string `json:"responseFormatType"`
}

type GetServerRequest struct{}

type ListServerRequest struct{}

type UpdateServerRequest struct{}
