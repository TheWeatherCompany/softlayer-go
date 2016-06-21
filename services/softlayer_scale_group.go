package services

import (
	"github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
)

type softlayer_Scale_Group_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Scale_Group_Service(client softlayer.Client) *softlayer_Scale_Group_Service {
	return &softlayer_Scale_Group_Service{
		client: client,
	}
}

func (slsgs *softlayer_Scale_Group_Service) GetName() string {
	return "SoftLayer_Scale_Group"
}

func (slsgs *softlayer_Scale_Group_Service) CreateObject(template data_types.SoftLayer_Scale_Group) (data_types.SoftLayer_Scale_Group, error) {
	return data_types.SoftLayer_Scale_Group{}, nil
}

func (slsgs *softlayer_Scale_Group_Service) DeleteObject(userid int) (bool, error) {
	return true, nil
}
