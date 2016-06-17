package data_types

type SoftLayer_Load_Balancer_Virtual_Server_Update_Parameters struct {
	Parameters []Softlayer_Load_Balancer_Virtual_Server_Parameters `json:"parameters"`
}

type Softlayer_Load_Balancer_Virtual_Server_Parameters struct {
	VirtualServers []*Softlayer_Load_Balancer_Virtual_Server `json:"virtualServers"`
}

type Softlayer_Load_Balancer_Virtual_Server struct {
	Id            int                        `json:"id,omitempty"`
	Allocation    int		         `json:"allocation"`
	Port          int		         `json:"port"`
	ServiceGroups []*Softlayer_Service_Group `json:"serviceGroups"`
}

type Softlayer_Service_Group struct {
	Id              int                  `json:"id,omitempty"`
	RoutingMethodId int                  `json:"routingMethodId"`
	RoutingTypeId   int                  `json:"routingTypeId"`
	Services        []*Softlayer_Service `json:"services"`
}

type Softlayer_Service struct {
	Id              int                          `json:"id,omitempty"`
	Enabled         int		             `json:"enabled"`
	Port            int		             `json:"port"`
	IpAddressId     int		             `json:"ipAddressId"`
	HealthChecks    []*Softlayer_Health_Check    `json:"healthChecks"`
	GroupReferences []*Softlayer_Group_Reference `json:"groupReferences"`
}

type Softlayer_Health_Check struct {
	HealthCheckTypeId int `json:"healthCheckTypeId"`
}

type Softlayer_Group_Reference struct {
	Weight int `json:"weight"`
}