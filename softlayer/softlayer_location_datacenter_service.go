package softlayer

import (
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Location_Datacenter_Service interface {
	Service

	GetDatacenters(objectMask []string, objectFilter string) ([]datatypes.SoftLayer_Location, error)
}
