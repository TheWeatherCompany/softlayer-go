package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/TheWeatherCompany/softlayer-go/client/fakes"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
	testhelpers "github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("Softlayer_Load_Balancer", func() {
	var (
		username, apiKey string

		fakeClient *slclientfakes.FakeSoftLayerClient

		lbService softlayer.SoftLayer_Load_Balancer_Service
		err       error

		lb datatypes.SoftLayer_Load_Balancer

		loadBalancerCreateOptions softlayer.SoftLayer_Load_Balancer_CreateOptions
		loadBalancerUpdateOptions datatypes.SoftLayer_Load_Balancer_Update

		virtualServerCreateOptions softlayer.SoftLayer_Load_Balancer_Service_Group_CreateOptions
		virtualServerUpdateOptions softlayer.SoftLayer_Load_Balancer_Service_Group_CreateOptions

		serviceCreateOptions softlayer.SoftLayer_Load_Balancer_Service_CreateOptions
		serviceUpdateOptions softlayer.SoftLayer_Load_Balancer_Service_CreateOptions
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		fakeClient.SoftLayerServices["SoftLayer_Product_Package"] = &testhelpers.MockProductPackageService{}

		lbService, err = fakeClient.GetSoftLayer_Load_Balancer_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(lbService).ToNot(BeNil())

		lb = datatypes.SoftLayer_Load_Balancer{}
	})

	Context("#CreateLoadBalancer", func() {
		BeforeEach(func() {
			loadBalancerCreateOptions = softlayer.SoftLayer_Load_Balancer_CreateOptions{
				Location:    "ams01",
				Connections: 15000,
				HaEnabled:   false,
			}
			datacentersResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetDatacenterByName.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, datacentersResponse)
			Expect(err).ToNot(HaveOccurred())
			responseOrder, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Product_Order_LoadBalancer_placeOrder.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, responseOrder)
			Expect(err).ToNot(HaveOccurred())
			responseLbList, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Account_Service_getLoadBalancers.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, responseLbList)
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new load balancer", func() {
			lb, err := lbService.CreateLoadBalancer(&loadBalancerCreateOptions)
			Expect(err).ToNot(HaveOccurred())
			Expect(lb.Id).To(Equal(150585))
			Expect(lb.IpAddressId).To(Equal(28772116))
			Expect(lb.IpAddress.IpAddress).To(Equal("161.202.117.116"))
			Expect(lb.HaEnabled).To(Equal(false))
			Expect(lb.ConnectionLimit).Should(BeNumerically(">=", 15000))
			Expect(lb.ConnectionLimit).Should(BeNumerically("<=", 150000))
			Expect(lb.SecurityCertificateId).To(Equal(0))
			Expect(lb.SoftlayerHardware[0].Datacenter.Name).To(Equal("ams01"))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancer(&loadBalancerCreateOptions)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancer(&loadBalancerCreateOptions)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#UpdateLoadBalancer", func() {
		BeforeEach(func() {
			lb.Id = 150585
			var securityCertificateId = new(int)
			*securityCertificateId = 1000
			loadBalancerUpdateOptions = datatypes.SoftLayer_Load_Balancer_Update{
				SecurityCertificateId: securityCertificateId,
			}
			createLoadBalancerResponse, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Load_Balancer_CreateLoadBalancer.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, createLoadBalancerResponse)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Updates the load balancer", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			result, err := lbService.UpdateLoadBalancer(lb.Id, &loadBalancerUpdateOptions)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err = lbService.UpdateLoadBalancer(lb.Id, &loadBalancerUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.UpdateLoadBalancer(lb.Id, &loadBalancerUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#DeleteLoadBalancer", func() {
		BeforeEach(func() {
			lb.Id = 150585
			getBillingItemResponse, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Load_Balancer_GetBillingItem.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, getBillingItemResponse)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Deletes the load balancer", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			_, err := lbService.DeleteObject(lb.Id)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err = lbService.UpdateLoadBalancer(lb.Id, &loadBalancerUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.UpdateLoadBalancer(lb.Id, &loadBalancerUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#CreateLoadBalancerVirtualServer", func() {
		BeforeEach(func() {
			lb.Id = 150585
			virtualServerCreateOptions = softlayer.SoftLayer_Load_Balancer_Service_Group_CreateOptions{
				Allocation:    1000,
				Port:          80,
				RoutingMethod: "CONSISTENT_HASH_IP",
				RoutingType:   "HTTP",
			}
			createLoadBalancerResponse, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Load_Balancer_CreateLoadBalancer.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, createLoadBalancerResponse)
			Expect(err).ToNot(HaveOccurred())
			routingMethodResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingMethodByName.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingMethodResponse)
			Expect(err).ToNot(HaveOccurred())
			routingTypeResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingTypeByName.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingTypeResponse)
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new load balancer virtual server", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			result, err := lbService.CreateLoadBalancerVirtualServer(lb.Id, &virtualServerCreateOptions)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancerVirtualServer(lb.Id, &virtualServerCreateOptions)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancerVirtualServer(lb.Id, &virtualServerCreateOptions)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#UpdateLoadBalancerVirtualServer", func() {
		BeforeEach(func() {
			lb.Id = 150585
			virtualServerUpdateOptions = softlayer.SoftLayer_Load_Balancer_Service_Group_CreateOptions{
				Allocation:    5000,
				Port:          443,
				RoutingMethod: "ROUND_ROBIN",
				RoutingType:   "HTTPS",
			}
			createLoadBalancerResponse, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Load_Balancer_CreateLoadBalancer.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, createLoadBalancerResponse)
			Expect(err).ToNot(HaveOccurred())
			routingMethodResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingMethodByName.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingMethodResponse)
			Expect(err).ToNot(HaveOccurred())
			routingTypeResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingTypeByName.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingTypeResponse)
			Expect(err).ToNot(HaveOccurred())
		})

		It("updates load balancer virtual server", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			result, err := lbService.CreateLoadBalancerVirtualServer(lb.Id, &virtualServerUpdateOptions)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancerVirtualServer(lb.Id, &virtualServerUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancerVirtualServer(lb.Id, &virtualServerUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#DeleteLoadBalancerVirtualServer", func() {
		BeforeEach(func() {
			lb.Id = 150585
			lb.VirtualServers = []*datatypes.Softlayer_Load_Balancer_Virtual_Server{{
				Id: 264313,
			}}
		})

		It("Deletes the load balancer virtual server", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			_, err := lbService.DeleteLoadBalancerVirtualServer(lb.VirtualServers[0].Id)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.DeleteLoadBalancerVirtualServer(lb.VirtualServers[0].Id)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.DeleteLoadBalancerVirtualServer(lb.VirtualServers[0].Id)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#CreateLoadBalancerService", func() {
		BeforeEach(func() {
			lb.Id = 150585
			serviceCreateOptions = softlayer.SoftLayer_Load_Balancer_Service_CreateOptions{
				ServiceGroupId:  260113,
				Enabled:         1,
				Port:            82,
				IpAddressId:     12345,
				HealthCheckType: "DNS",
				Weight:          1000,
			}
			createLoadBalancerResponse, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Load_Balancer_GetLoadBalancer_with_VirtualServer.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, createLoadBalancerResponse)
			Expect(err).ToNot(HaveOccurred())
			routingTypeResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingTypeById.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingTypeResponse)
			Expect(err).ToNot(HaveOccurred())
			routingMethodResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingMethodById.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingMethodResponse)
			Expect(err).ToNot(HaveOccurred())
			healthCheckTypeResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetHealthCheckTypeByName.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, healthCheckTypeResponse)
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new load balancer service", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			result, err := lbService.CreateLoadBalancerService(lb.Id, &serviceCreateOptions)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancerService(lb.Id, &serviceCreateOptions)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.CreateLoadBalancerService(lb.Id, &serviceCreateOptions)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#UpdateLoadBalancerService", func() {
		BeforeEach(func() {
			lb.Id = 150585
			lb.VirtualServers = []*datatypes.Softlayer_Load_Balancer_Virtual_Server{{
				Id:         264313,
				Allocation: 100,
				Port:       82,
				ServiceGroups: []*datatypes.Softlayer_Service_Group{{
					Id:              260113,
					RoutingMethodId: 21,
					RoutingTypeId:   5,
					Services: []*datatypes.Softlayer_Service{{
						Id: 529393,
						HealthChecks: []*datatypes.Softlayer_Health_Check{{
							HealthCheckTypeId: 3,
						}},
					}},
				}},
			}}
			serviceUpdateOptions = softlayer.SoftLayer_Load_Balancer_Service_CreateOptions{
				Port:            443,
				HealthCheckType: "DNS",
			}
			createLoadBalancerResponse, err := testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Load_Balancer_GetLoadBalancer_with_Service.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, createLoadBalancerResponse)
			Expect(err).ToNot(HaveOccurred())
			routingTypeResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingTypeById.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingTypeResponse)
			Expect(err).ToNot(HaveOccurred())
			routingMethodResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetRoutingMethodById.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, routingMethodResponse)
			Expect(err).ToNot(HaveOccurred())
			healthCheckTypeIdResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetHealthCheckTypeById.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, healthCheckTypeIdResponse)
			Expect(err).ToNot(HaveOccurred())
			healthCheckTypeResponse, err := testhelpers.ReadJsonTestFixtures("common", "GetHealthCheckTypeByName.json")
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, healthCheckTypeResponse)
			Expect(err).ToNot(HaveOccurred())
		})

		It("update load balancer service", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			result, err := lbService.UpdateLoadBalancerService(lb.Id, lb.VirtualServers[0].ServiceGroups[0].Id,
				lb.VirtualServers[0].ServiceGroups[0].Services[0].Id, &serviceUpdateOptions)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.UpdateLoadBalancerService(lb.Id, lb.VirtualServers[0].ServiceGroups[0].Id,
						lb.VirtualServers[0].ServiceGroups[0].Services[0].Id, &serviceUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.UpdateLoadBalancerService(lb.Id, lb.VirtualServers[0].ServiceGroups[0].Id,
						lb.VirtualServers[0].ServiceGroups[0].Services[0].Id, &serviceUpdateOptions)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#DeleteLoadBalancerService", func() {
		BeforeEach(func() {
			lb.Id = 150585
			lb.VirtualServers = []*datatypes.Softlayer_Load_Balancer_Virtual_Server{{
				Id:         264313,
				Allocation: 100,
				Port:       82,
				ServiceGroups: []*datatypes.Softlayer_Service_Group{{
					Id:              260113,
					RoutingMethodId: 21,
					RoutingTypeId:   5,
					Services: []*datatypes.Softlayer_Service{{
						Id: 529393,
						HealthChecks: []*datatypes.Softlayer_Health_Check{{
							HealthCheckTypeId: 3,
						}},
					}},
				}},
			}}
		})

		It("Deletes the load balancer service", func() {
			fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte("true"))
			Expect(err).ToNot(HaveOccurred())

			_, err := lbService.DeleteLoadBalancerService(lb.VirtualServers[0].ServiceGroups[0].Services[0].Id)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.DeleteLoadBalancerService(lb.VirtualServers[0].ServiceGroups[0].Services[0].Id)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.DoRawHttpRequestResponses = append(fakeClient.DoRawHttpRequestResponses, []byte{})
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err := lbService.DeleteLoadBalancerService(lb.VirtualServers[0].ServiceGroups[0].Services[0].Id)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})
})
