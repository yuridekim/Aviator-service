package types

type NetworkInterface struct {
	NetworkInterfaceNo          string
	NetworkInterfaceName        string
	SubnetNo                    string
	DeleteOnTermination         bool
	IsDefault                   bool
	DeviceName                  string
	NetworkInterfaceStatus      CommonCode
	InstanceType                CommonCode
	InstanceNo                  string
	IP                          string
	MacAddress                  string
	AccessControlGroupNoList    []AccessControlGroup
	NetworkInterfaceDescription string
	SecondaryIPList             []SecondaryIP
}

type NetworkInterfaceList struct {
	TotalRows            int
	NetworkInterfaceList []NetworkInterface
}

type NetworkInterfaceNoList struct {
	NetworkInterfaceNoList []string
}

type AccessControlGroup struct {
}

type SecondaryIP struct {
}
