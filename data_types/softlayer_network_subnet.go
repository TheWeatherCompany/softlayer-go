package data_types

type SoftLayer_Network_Subnet struct {
	Cidr              int                                  `json:"cidr,omitempty"`
	Id                int                                  `json:"Id,omitempty"`
	NetworkIdentifier string                               `json:"NetworkIdentifier,omitempty"`
	IpAddresses       []SoftLayer_Network_Subnet_IpAddress `json:"IpAddresses,omitempty"`
}
