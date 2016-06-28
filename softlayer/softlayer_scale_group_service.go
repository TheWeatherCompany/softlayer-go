package softlayer

import (
	"github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Scale_Group_Service interface {
	Service

	CreateObject(template data_types.SoftLayer_Scale_Group) (data_types.SoftLayer_Scale_Group, error)
	GetObject(groupId int) (data_types.SoftLayer_Scale_Group, error)
	ForceDeleteObject(scaleGroupId int) (bool, error)
}
