package data_types

type SoftLayer_Scale_LoadBalancer struct {
	HealthCheckId   int `json:"healthCheckId,omitempty"`
	Id              int `json:"id,omitempty"`
	Port            int `json:"port,omitempty"`
	VirtualServerId int `json:"virtualServerId,omitempty"`
}
