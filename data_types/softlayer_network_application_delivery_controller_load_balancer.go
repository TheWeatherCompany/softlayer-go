package data_types

type SoftLayer_Network_Application_Delivery_Controller_Load_Balancer struct {
	Id                    int                  `json:"id,omitempty"`
	ConnectionLimit       int                  `json:"connectionLimit,omitempty"`
	IpAddressId           int                  `json:"ipAddressId,omitempty"`
	SecurityCertificateId int                  `json:"securityCertificateId,omitempty"`
	SoftlayerHardware     []SoftLayer_Hardware `json:"loadBalancerHardware,omitempty"`
}