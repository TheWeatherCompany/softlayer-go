package softlayer

import (
	"github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Scale_Policy_Service interface {
	Service

	CreateObject(template data_types.SoftLayer_Scale_Policy) (data_types.SoftLayer_Scale_Policy, error)
	DeleteObject(scalePolicyId int) (bool, error)
}
