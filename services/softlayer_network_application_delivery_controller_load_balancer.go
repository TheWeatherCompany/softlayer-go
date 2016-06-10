package services

import (
	"errors"
	"fmt"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	softlayer "github.com/TheWeatherCompany/softlayer-go/softlayer"
	"github.com/hashicorp/terraform/helper/resource"
	"strconv"
	"strings"
	"time"
	"bytes"
	"github.com/TheWeatherCompany/softlayer-go/common"
	"encoding/json"
)

const (
	PACKAGE_TYPE_APPLICATION_DELIVERY_CONTROLLER_LOAD_BALANCER = "ADDITIONAL_SERVICES_LOAD_BALANCER"
	ORDER_TYPE_APPLICATION_DELIVERY_CONTROLLER_LOAD_BALANCER   = "SoftLayer_Container_Product_Order_Network_LoadBalancer"
	PACKAGE_ID_APPLICATION_DELIVERY_CONTROLLER_LOAD_BALANCER   = 194
)

type softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service struct {
	client softlayer.Client
}

func NewSoftLayer_Network_Application_Delivery_Controller_Load_Balancer_Service(client softlayer.Client) *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service {
	return &softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service{
		client: client,
	}
}

func (slnadcs *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service) GetName() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}

func (slnadclbs *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service) CreateLoadBalancer(createOptions *softlayer.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer_Service_CreateOptions) (datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer, error) {

	orderService, err := slnadclbs.client.GetSoftLayer_Product_Order_Service()
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, err
	}

	items, err := slnadclbs.FindCreatePriceItems(createOptions)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, err
	}

	order := datatypes.SoftLayer_Container_Product_Order_Network_Application_Delivery_Controller{
		PackageId:   PACKAGE_ID_APPLICATION_DELIVERY_CONTROLLER_LOAD_BALANCER,
		ComplexType: ORDER_TYPE_APPLICATION_DELIVERY_CONTROLLER_LOAD_BALANCER,
		Location:    createOptions.Location,
		Prices:      items,
		Quantity:    1,
	}

	receipt, err := orderService.PlaceContainerOrderApplicationDeliveryController(order)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, err
	}

	vpx, err := slnadclbs.findLoadBalancerByOrderId(receipt.OrderId)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, err
	}

	return vpx, nil
}

func (slnadclbs *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service) GetObject(id int) (datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer, error) {
	response, errorCode, err := slnadclbs.client.GetHttpClient().DoRawHttpRequest(fmt.Sprintf("%s/%d/getObject.json", slnadclbs.GetName(), id), "GET", new(bytes.Buffer))
	if err != nil {
		errorMessage := fmt.Sprintf("softlayer-go: could not perform SoftLayer_Network_Application_Delivery_Controller_Load_Balancer#getObject, error message '%s'", err.Error())
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, errors.New(errorMessage)
	}

	if common.IsHttpErrorCode(errorCode) {
		errorMessage := fmt.Sprintf("softlayer-go: could not perform SoftLayer_Network_Application_Delivery_Controller_Load_Balancer#getObject, HTTP error code: '%d'", errorCode)
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, errors.New(errorMessage)
	}

	lb := datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}
	err = json.Unmarshal(response, &lb)
	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, err
	}

	return lb, nil
}

func (slnadclbs *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service) FindCreatePriceItems(createOptions *softlayer.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer_Service_CreateOptions) ([]datatypes.SoftLayer_Product_Item_Price, error) {
	items, err := slnadclbs.getApplicationDeliveryControllerLoadBalancerItems()
	if err != nil {
		return []datatypes.SoftLayer_Product_Item_Price{}, err
	}

	lbKey := slnadclbs.getLoadBalancerPriceItemKeyName(createOptions.Connections)

	var lbItemPrice datatypes.SoftLayer_Product_Item_Price

	for _, item := range items {
		itemKey := item.Key
		if itemKey == lbKey {
			lbItemPrice = item.Prices[0]
		}
	}

	var errorMessages []string

	if lbItemPrice.Id == 0 {
		errorMessages = append(errorMessages, fmt.Sprintf("LB Connections field has an incorrect value"))
	}

	if len(errorMessages) > 0 {
		err = errors.New(strings.Join(errorMessages, "\n"))
		return []datatypes.SoftLayer_Product_Item_Price{}, err
	}

	return []datatypes.SoftLayer_Product_Item_Price{lbItemPrice}, nil
}

func (slnadclbs *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service) findLoadBalancerByOrderId(orderId int) (datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer, error) {
	ObjectFilter := string(`{"adcLoadBalancers":{"dedicatedBillingItem":{"orderItem":{"order":{"id":{"operation":` + strconv.Itoa(orderId) + `}}}}}}`)

	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"complete"},
		Refresh: func() (interface{}, string, error) {
			accountService, err := slnadclbs.client.GetSoftLayer_Account_Service()
			if err != nil {
				return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, "", err
			}
			lbs, err := accountService.GetApplicationDeliveryControllerLoadBalancersWithFilter(ObjectFilter)
			if err != nil {
				return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, "", err
			}

			if len(lbs) == 1 {
				return lbs[0], "complete", nil
			} else if len(lbs) == 0 {
				return nil, "pending", nil
			} else {
				return nil, "", fmt.Errorf("Expected one load balancer: %s", err)
			}
		},
		Timeout:    10 * time.Minute,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	pendingResult, err := stateConf.WaitForState()

	if err != nil {
		return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{}, err
	}

	var result, ok = pendingResult.(datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer)

	if ok {
		return result, nil
	}

	return datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer{},
		fmt.Errorf("Cannot find Application Delivery Controller Load Balancer with order id '%d'", orderId)
}

func (slnadclbs *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service) getApplicationDeliveryControllerLoadBalancerItems() ([]datatypes.SoftLayer_Product_Item, error) {
	productPackageService, err := slnadclbs.client.GetSoftLayer_Product_Package_Service()
	if err != nil {
		return []datatypes.SoftLayer_Product_Item{}, err
	}

	return productPackageService.GetItemsByType(PACKAGE_TYPE_APPLICATION_DELIVERY_CONTROLLER_LOAD_BALANCER)
}

func (slnadclbs *softLayer_Network_Application_Delivery_Controller_Load_Balancer_Service) getLoadBalancerPriceItemKeyName(connections int) string {
	name := "DEDICATED_LOAD_BALANCER_WITH_HIGH_AVAILABILITY_AND_SSL"
	return strings.Join([]string{name, strconv.Itoa(connections), "CONNECTIONS"}, DELIMITER)
}
