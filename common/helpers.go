package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
	"strconv"
)

const (
	DATACENTER_TYPE_NAME   = "SoftLayer_Location_Datacenter"
	ROUTING_TYPE_NAME      = "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type"
	ROUTING_METHOD_NAME    = "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method"
	HEALTH_CHECK_TYPE_NAME = "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type"
)

type lookupId func([]byte) (interface{}, error)

func isInt(key interface{}) (bool, error) {
	switch v := key.(type) {
	case int:
		return true, nil
	case string:
		return false, nil
	default:
		return false, fmt.Errorf("Expected type string or int. Recieved type %s", v)
	}
}

func getIdByName(client softlayer.Client, nameMask string, nameType string, nameTypeGet string, key interface{}, getById bool, lookupFunc lookupId) (interface{}, error) {
	var ObjectFilter string
	if getById {
		ObjectFilter = string(`{"id":{"operation":"` + strconv.Itoa(key.(int)) + `"}}`)
	} else {
		ObjectFilter = string(`{"` + nameMask + `":{"operation":"` + key.(string) + `"}}`)
	}
	ObjectMasks := []string{"id", nameMask}

	response, errorCode, err := client.GetHttpClient().DoRawHttpRequestWithObjectFilterAndObjectMask(fmt.Sprintf("%s/%s", nameType, nameTypeGet), ObjectMasks, ObjectFilter, "GET", new(bytes.Buffer))
	if err != nil {
		return -1, err
	}

	if IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not retrieve %s, HTTP error code: '%d'", nameType, errorCode)
		return -1, errors.New(errorMessage)
	}

	return lookupFunc(response)
}

func GetDatacenter(client softlayer.Client, key interface{}) (interface{}, error) {
	getById, err := isInt(key)

	if err != nil {
		return -1, err
	}

	return getIdByName(client, "name", DATACENTER_TYPE_NAME, "getDatacenters.json", key, getById,
		func(response []byte) (interface{}, error) {
			locations := []datatypes.SoftLayer_Location{}

			err := json.Unmarshal(response, &locations)
			if err != nil {
				return -1, err
			}

			for _, location := range locations {
				if getById && location.Id == key.(int) {
					return location.Name, nil
				} else if !getById && location.Name == key.(string) {
					return location.Id, nil
				}
			}

			return -1, fmt.Errorf("Datacenter %s not found", key)
		})
}

func GetRoutingType(client softlayer.Client, key interface{}) (interface{}, error) {
	getById, err := isInt(key)

	if err != nil {
		return -1, err
	}

	return getIdByName(client, "keyname", ROUTING_TYPE_NAME, "getAllObjects.json", key, getById,
		func(response []byte) (interface{}, error) {
			routingTypes := []datatypes.SoftLayer_Routing_Type{}

			err := json.Unmarshal(response, &routingTypes)
			if err != nil {
				return -1, err
			}

			for _, routingType := range routingTypes {
				if getById && routingType.Id == key.(int) {
					return routingType.KeyName, nil
				} else if !getById && routingType.KeyName == key.(string) {
					return routingType.Id, nil
				}
			}

			return -1, fmt.Errorf("Routing type %s not found", key)
		})
}

func GetRoutingMethod(client softlayer.Client, key interface{}) (interface{}, error) {
	getById, err := isInt(key)

	if err != nil {
		return -1, err
	}

	return getIdByName(client, "keyname", ROUTING_METHOD_NAME, "getAllObjects.json", key, getById,
		func(response []byte) (interface{}, error) {
			routingMethods := []datatypes.SoftLayer_Routing_Method{}

			err := json.Unmarshal(response, &routingMethods)
			if err != nil {
				return -1, err
			}

			for _, routingMethod := range routingMethods {
				if getById && routingMethod.Id == key.(int) {
					return routingMethod.KeyName, nil
				} else if !getById && routingMethod.KeyName == key.(string) {
					return routingMethod.Id, nil
				}
			}

			return -1, fmt.Errorf("Routing method %s not found", key)
		})
}

func GetHealthCheckType(client softlayer.Client, key interface{}) (interface{}, error) {
	getById, err := isInt(key)

	if err != nil {
		return -1, err
	}

	return getIdByName(client, "keyname", HEALTH_CHECK_TYPE_NAME, "getAllObjects.json", key, getById,
		func(response []byte) (interface{}, error) {
			healthCheckTypes := []datatypes.SoftLayer_Health_Check_Type{}

			err := json.Unmarshal(response, &healthCheckTypes)
			if err != nil {
				return -1, err
			}

			for _, healthCheckType := range healthCheckTypes {
				if getById && healthCheckType.Id == key.(int) {
					return healthCheckType.KeyName, nil
				} else if !getById && healthCheckType.KeyName == key.(string) {
					return healthCheckType.Id, nil
				}
			}

			return -1, fmt.Errorf("Health check type %s not found", key)
		})
}
