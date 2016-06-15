package data_types

type SoftLayer_Network_Application_Delivery_Controller_Load_Balancer struct {
	Id                    int                                       `json:"id,omitempty"`
	ConnectionLimit       int                                       `json:"connectionLimit,omitempty"`
	IpAddressId           int                                       `json:"ipAddressId,omitempty"`
	SecurityCertificateId int                                       `json:"securityCertificateId,omitempty"`
	IpAddress             SoftLayer_Network_Ip_Address              `json:"ipAddress,omitempty"`
	SoftlayerHardware     []SoftLayer_Hardware                      `json:"loadBalancerHardware,omitempty"`
	VirtualServers        []*Softlayer_Load_Balancer_Virtual_Server `json:"virtualServers,omitempty"`
}

type SoftLayer_Network_Ip_Address struct {
	IpAddress             string                      `json:"ipAddress,omitempty"`
	SubnetId              int                         `json:"subnetId,omitempty"`
}