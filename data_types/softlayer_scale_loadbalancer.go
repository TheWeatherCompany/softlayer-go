package data_types

type SoftLayer_Scale_LoadBalancer struct {
	HealthCheck     *SoftLayer_Health_Check `json:"healthCheck,omitempty"`
	Id              int                     `json:"id,omitempty"`
	Port            int                     `json:"port,omitempty"`
	VirtualServerId int                     `json:"virtualServerId,omitempty"`
}

type SoftLayer_Health_Check struct {
	Attributes        []SoftLayer_Health_Check_Attribute `json:"attributes,omitempty"`
	Id                int                                `json:"id,omitempty"`
	HealthCheckTypeId int                                `json:"healthCheckTypeId,omitempty"`
	Name              string                             `json:"name,omitempty"`
}

type SoftLayer_Health_Check_Attribute struct {
	Type  *SoftLayer_Health_Check_Attribute_Type `json:"type,omitempty"`
	Value string                                 `json:"value,omitempty"`
}

type SoftLayer_Health_Check_Attribute_Type struct {
	Id      int    `json:"id,omitempty"`
	Keyname string `json:"keyname,omitempty"`
}
