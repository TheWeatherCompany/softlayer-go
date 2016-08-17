package data_types

import (
	"time"
)

type SoftLayer_Network_Vlan struct {
	AccountId       int                        `json:"accountId,omitempty"`
	Id              int                        `json:"id,omitempty"`
	ModifyDate      *time.Time                 `json:"modifyDate,omitempty"`
	Name            string                     `json:"name,omitempty"`
	NetworkVrfId    int                        `json:"networkVrfId,omitempty"`
	Note            string                     `json:"note,omitempty"`
	PrimaryRouter   *SoftLayer_Hardware_Router `json:"primaryRouter,omitempty"`
	PrimarySubnetId int                        `json:"primarySubnetId,omitempty"`
	VlanNumber      int                        `json:"vlanNumber,omitempty"`
	PrimarySubnets  []SoftLayer_Network_Subnet `json:"primarySubnets,omitempty"`
}

type SoftLayer_Network_Vlan_Template struct {
	AccountId       int                        `json:"accountId,omitempty"`
	Id              int                        `json:"id,omitempty"`
	ModifyDate      *time.Time                 `json:"modifyDate,omitempty"`
	Name            string                     `json:"name,omitempty"`
	NetworkVrfId    int                        `json:"networkVrfId,omitempty"`
	Note            string                     `json:"note,omitempty"`
	PrimarySubnetId int                        `json:"primarySubnetId,omitempty"`
	VlanNumber      int                        `json:"vlanNumber,omitempty"`
	PrimaryRouter   *SoftLayer_Hardware_Router `json:"primaryRouter,omitempty"`
	PrimarySubnets  []SoftLayer_Network_Subnet `json:"primarySubnets,omitempty"`
}
