package data_types

type SoftLayer_Scale_Group struct {
	Cooldown                   int                              `json:"cooldown,omitempty"`
	Id                         int                              `json:"id,omitempty"`
	LoadBalancers              []SoftLayer_Scale_LoadBalancer   `json:"loadBalancers,omitempty"`
	MaximumMemberCount         int                              `json:"maximumMemberCount,omitempty"`
	MinimumMemberCount         int                              `json:"minimumMemberCount,omitempty"`
	Name                       string                           `json:"name,omitempty"`
	NetworkVlans               []SoftLayer_Scale_Network_Vlan   `json:"networkVlans,omitempty"`
	Policies                   []SoftLayer_Scale_Policy         `json:"policies,omitempty"`
	RegionalGroupId            int                              `json:"regionalGroupId,omitempty"`
	SuspendedFlag              bool                             `json:"suspendedFlag,omitempty"`
	TerminationPolicyId        int                              `json:"terminationPolicyId,omitempty"`
	VirtualGuestMemberTemplate SoftLayer_Virtual_Guest_Template `json:"virtualGuestMemberTemplate,omitempty"`
}
