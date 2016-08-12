package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TheWeatherCompany/softlayer-go/common"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	softlayer "github.com/TheWeatherCompany/softlayer-go/softlayer"
)

type softLayer_Location_Datacenter_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Location_Datacenter_Service(client softlayer.Client) *softLayer_Location_Datacenter_Service {
	return &softLayer_Location_Datacenter_Service{
		client: client,
	}
}

func (slds *softLayer_Location_Datacenter_Service) GetName() string {
	return "SoftLayer_Location_Datacenter"
}

func (slds *softLayer_Location_Datacenter_Service) GetDatacenters(objectMask []string, objectFilter string) ([]datatypes.SoftLayer_Location, error) {
	path := fmt.Sprintf("%s/%s", slds.GetName(), "getDatacenters.json")

	responseBytes, errorCode, err := slds.client.GetHttpClient().DoRawHttpRequestWithObjectFilterAndObjectMask(path, objectMask, objectFilter, "GET", &bytes.Buffer{})
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Location_Datacenter#getDatacenters, error message '%s'", err.Error())
		return []datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not SoftLayer_Location_Datacenter#getDatacenters, HTTP error code: '%d'", errorCode)
		return []datatypes.SoftLayer_Location{}, errors.New(errorMessage)
	}

	locations := []datatypes.SoftLayer_Location{}
	err = json.Unmarshal(responseBytes, &locations)
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: failed to decode JSON response, err message '%s'", err.Error())
		err := errors.New(errorMessage)
		return []datatypes.SoftLayer_Location{}, err
	}

	return locations, nil
}
