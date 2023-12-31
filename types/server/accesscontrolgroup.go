package types

type AccessControlGroup struct {
	AccessControlGroupNo          string     `xml:"accessControlGroupNo"`
	AccessControlGroupName        string     `xml:"accessControlGroupName"`
	IsDefault                     bool       `xml:"isDefault"`
	VpcNo                         string     `xml:"vpcNo"`
	AccessControlGroupStatus      CommonCode `xml:"accessControlGroupStatus"`
	AccessControlGroupDescription string     `xml:"accessControlGroupDescription"`
}

type AccessControlGroupList struct {
	TotalRows              int                  `xml:"totalRows"`
	AccessControlGroupList []AccessControlGroup `xml:"accessControlGroupList>accessControlGroup"`
}
