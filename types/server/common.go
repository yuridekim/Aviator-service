package types

type CommonCode struct {
	Code     string // server status : int | creat | run | nstop
	CodeName string
}

type NetworkInterface struct {
	networkInterfaceOrder    int
	networkInterfaceNo       string
	subnetNo                 string
	ip                       string
	accessControlGroupNoList []accessControlGroup
}

type accessControlGroup struct {
}

// BlockDevicePartition means block disk parition info (such as AWS EBS)
type BlockDevicePartition struct {
	MountPoint    string
	PartitionSize string
}

type BlockStorageMapping struct {
	order                      int
	snapshotInstanceNo         string
	blockStorageSize           string
	blockStorageName           string
	blockStorageVolumeTypeCode string
	encrypted                  string
}

// NetworkInterfaceNoList is same as NIC number list
type NetworkInterfaceNoList struct {
	NetworkInterfaceNoList []NetworkInterface
}

type BlockDevicePartitionList struct {
	BlockDevicePartitionList []BlockDevicePartition
}

type BlockStorageMappingList struct {
	BlockStorageMappingList []BlockStorageMapping
}
