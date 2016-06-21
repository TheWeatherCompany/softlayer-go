package data_types

import "time"

type SoftLayer_Scale_Policy struct {
	Actions      []SoftLayer_Scale_Policy_Action  `json:"actions,omitempty"`
	Cooldown     int                              `json:"cooldown,omitempty"`
	Id           int                              `json:"id,omitempty"`
	Name         string                           `json:"name,omitempty"`
	ScaleGroupId int                              `json:"scaleGroupId,omitempty"`
	Triggers     []SoftLayer_Scale_Policy_Trigger `json:"triggers,omitempty"`
}

type SoftLayer_Scale_Policy_Action struct {
	Id     int `json:"id,omitempty"`
	TypeId int `json:"typeId,omitempty"`
}

type SoftLayer_Scale_Policy_Action_Scale struct {
	SoftLayer_Scale_Policy_Action
	Amount    int    `json:"amount,omitempty"`
	ScaleType string `json:"scaleType,omitempty"`
}

type SoftLayer_Scale_Policy_Trigger struct {
	Id     int `json:"id,omitempty"`
	TypeId int `json:"typeId,omitempty"`
}

type SoftLayer_Scale_Policy_Trigger_OneTime struct {
	SoftLayer_Scale_Policy_Trigger
	Date *time.Time `json:"date,omitempty"`
}

type SoftLayer_Scale_Policy_Trigger_Repeating struct {
	SoftLayer_Scale_Policy_Trigger
	Schedule string `json:"schedule,omitempty"`
}

type SoftLayer_Scale_Policy_Trigger_ResourceUse struct {
	SoftLayer_Scale_Policy_Trigger
	Watches []SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch `json:"watches,omitempty"`
}

type SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch struct {
	Id       int    `json:"id,omitempty"`
	Metric   string `json:"metric,omitempty"`
	Operator string `json:"operator,omitempty"`
	Period   string `json:"period,omitempty"`
	Value    string `json:"value,omitempty"`
}
