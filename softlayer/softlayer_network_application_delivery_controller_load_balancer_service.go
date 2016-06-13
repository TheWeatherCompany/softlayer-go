package softlayer

import (
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
)

type SoftLayer_Network_Application_Delivery_Controller_Load_Balancer_Service_CreateOptions struct {
	Connections int
	Location    string
}

type SoftLayer_Network_Application_Delivery_Controller_Load_Balancer_Service interface {
	Service

	CreateLoadBalancer(createOptions *SoftLayer_Network_Application_Delivery_Controller_Load_Balancer_Service_CreateOptions) (datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer, error)

        GetObject(id int) (datatypes.SoftLayer_Network_Application_Delivery_Controller_Load_Balancer, error)

        DeleteObject(id int) (bool, error)

	FindCreatePriceItems(createOptions *SoftLayer_Network_Application_Delivery_Controller_Load_Balancer_Service_CreateOptions) ([]datatypes.SoftLayer_Product_Item_Price, error)
}
