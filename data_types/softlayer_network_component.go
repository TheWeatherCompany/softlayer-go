package data_types

type SoftLayer_Network_Component struct {
	NetworkVlanId int                              `json:"networkVlanId,omitempty"`
	NetworkVlan   *SoftLayer_Network_Vlan_Template `json:"networkVlan,omitempty"`
}