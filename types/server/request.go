package types

// 필수가 아닌 필드(필수 여부: No)는 주석 처리 해두었음.
// 필요할 때 주석 해제
// 필수가 아닌 필드 중 (필수 여부: Conditional)는 주석 처리 안 했음.
// 필요할 때 주석 처리
type CreateServerRequest struct {
	//RegionCode                        string               `json:"regionCode"`
	// ServerImageProductCode 와 MemberServerImageInstanceNo 둘 중 하나는 무조건 필수 기재
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

	//기본 네트워크 인터페이스로 설정하려면 0 입력
	// min:0, max:2
	NetworkInterfaceOrder int `json:"networkInterfaceList.N.networkInterfaceOrder"`

	// 사용자가 직접 생성한 네트워크 인터페이스를 추가하려고 하는 경우 해당 네트워크 인터페이스 번호 입력
	NetworkInterfaceNo string `json:"networkInterfaceNo"`

	// 새로 생성할 네트워크 인터페이스의 서브넷 또는 추가하려고 하는 기존 네트워크 인터페이스의 서브넷 결정
	// 기본 네트워크 인터페이스인 경우에는 자동으로 할당
	NetworkInterfaceSubnetNo string `json:"networkInterfaceSubnetNo"`

	//NetworkInterfaceIp 	  string               `json:"networkInterfaceIp"`

	// 네트워크 인터페이스를 새로 생성하는 경우 반드시 적용할 ACG 결정
	// 최대 3개의 ACG 적용 가능
	// accessControlGroupNo는 getAccessControlGroupList 액션을 통해 획득 가능
	AccessControlGroupNoListN []string `json:"accessControlGroupNoList"`
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
