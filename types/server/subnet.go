package types

import "time"

type Subnet struct {
	SubnetNo     string     `xml:"subnetNo"`
	VpcNo        string     `xml:"vpcNo"`
	ZoneCode     string     `xml:"zoneCode"`
	SubnetName   string     `xml:"subnetName"`
	Subnet       string     `xml:"subnet"`
	SubnetStatus CommonCode `xml:"subnetStatus"`
	CreateDate   time.Time  `xml:"createDate"`
	SubnetType   CommonCode `xml:"subnetType"`
	UsageType    CommonCode `xml:"usageType"`
	NetworkAclNo string     `xml:"networkAclNo"`
}

type SubnetList struct {
	TotalRows int      `xml:"totalRows"`
	Subnet    []Subnet `xml:"subnetList>subnet"`
}
