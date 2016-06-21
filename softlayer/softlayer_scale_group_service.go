package softlayer

import (
	"github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Scale_Group_Service interface {
	Service

	CreateObject(template data_types.SoftLayer_Scale_Group) (data_types.SoftLayer_Scale_Group, error)
	DeleteObject(scaleGroupId int) (bool, error)
}
