package types

type NetworkInterface struct {
	NetworkInterfaceNo   string `xml:"networkInterfaceNo"`
	NetworkInterfaceName string `xml:"networkInterfaceName"`
	SubnetNo             string `xml:"subnetNo"`
	DeleteOnTermination  bool   `xml:"deleteOnTermination"`
	IsDefault            bool   `xml:"isDefault"`
	//DeviceName                  string               `xml:"deviceName"` // Conditional
	NetworkInterfaceStatus CommonCode `xml:"networkInterfaceStatus"`
	//InstanceType                CommonCode           `xml:"instanceType"` // Conditional
	//InstanceNo                  string               `xml:"instanceNo"` // Conditional
	IP         string `xml:"ip"`
	MacAddress string `xml:"macAddress"`
	//AccessControlGroupNoList    []AccessControlGroup `xml:"accessControlGroupNoList"` // Conditional
	//NetworkInterfaceDescription string               `xml:"networkInterfaceDescription"` // Conditional
	//SecondaryIPList             []SecondaryIP        `xml:"secondaryIPList"` // Conditional
}

type NetworkInterfaceList struct {
	TotalRows            int                `xml:"totalRows"`
	NetworkInterfaceList []NetworkInterface `xml:"networkInterfaceList>networkInterface"`
}

type NetworkInterfaceNoList struct {
	NetworkInterfaceNoList []string
}

type SecondaryIP struct {
}
