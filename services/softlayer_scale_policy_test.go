package services_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/TheWeatherCompany/softlayer-go/client/fakes"
	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	softlayer "github.com/TheWeatherCompany/softlayer-go/softlayer"
	testhelpers "github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer_Scale_Policy_Service", func() {
	var (
		username, apiKey string
		err              error

		fakeClient *slclientfakes.FakeSoftLayerClient

		scalePolicyService softlayer.SoftLayer_Scale_Policy_Service

		scalePolicy         datatypes.SoftLayer_Scale_Policy
		scalePolicyTemplate datatypes.SoftLayer_Scale_Policy
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		scalePolicyService, err = fakeClient.GetSoftLayer_Scale_Policy_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(scalePolicyService).ToNot(BeNil())

		scalePolicy = datatypes.SoftLayer_Scale_Policy{}
		scalePolicyTemplate = datatypes.SoftLayer_Scale_Policy{}
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := scalePolicyService.GetName()
			Expect(name).To(Equal("SoftLayer_Scale_Policy"))
		})
	})

	Context("#CreateObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Scale_Policy_Service_createAndGetObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

	})

	It("creates a new SoftLayer_Scale_Policy instance", func() {
		scalePolicyTemplate = datatypes.SoftLayer_Scale_Policy{
			Id:           4567,
			Name:         "fake-name",
			Cooldown:     30,
			ScaleGroupId: 1234567,
			ScaleActions: []datatypes.SoftLayer_Scale_Policy_Action{
				{
					TypeId:    1,
					Amount:    1,
					ScaleType: "RELATIVE",
				},
			},
			Triggers: []datatypes.SoftLayer_Scale_Policy_Trigger{
				{
					Id: 1,
					TypeId: 1,
				},
			},	
		}
		scalePolicy, err = scalePolicyService.CreateObject(scalePolicyTemplate)
		Expect(scalePolicy.Id).NotTo(BeNil())
		Expect(scalePolicy.Name).NotTo(BeNil())
		Expect(scalePolicy.Cooldown).NotTo(BeNil())
		Expect(scalePolicy.ScaleGroupId).NotTo(BeNil())
	})

	Context("when HTTP client returns error codes 40x or 50x", func() {
		It("fails for error code 40x", func() {
			errorCodes := []int{400, 401, 499}
			for _, errorCode := range errorCodes {
				fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

				scalePolicyTemplate = datatypes.SoftLayer_Scale_Policy{
					Name:         "fake-name",
					Cooldown:     30,
					ScaleGroupId: 1234567,
					ScaleActions: []datatypes.SoftLayer_Scale_Policy_Action{
						{
							TypeId:    1,
							Amount:    1,
							ScaleType: "RELATIVE",
						},
					},
					Triggers: []datatypes.SoftLayer_Scale_Policy_Trigger{
						{
							Id: 1,
							TypeId: 1,
						},
					},	
				}
			}
		})

		It("fails for error code 40x", func() {
			errorCodes := []int{500, 501, 599}
			for _, errorCode := range errorCodes {
				fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

				scalePolicyTemplate = datatypes.SoftLayer_Scale_Policy{
					Name:         "fake-name",
					Cooldown:     30,
					ScaleGroupId: 1234567,
					ScaleActions: []datatypes.SoftLayer_Scale_Policy_Action{
						{
							TypeId:    1,
							Amount:    1,
							ScaleType: "RELATIVE",
						},
					},
					Triggers: []datatypes.SoftLayer_Scale_Policy_Trigger{
						{
							Id: 1,
							TypeId: 1,
						},
					},	
				}
			}
		})
	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			scalePolicy.Id = 4567
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Scale_Policy_Service_createAndGetObject.json")
			Expect(err).ToNot(HaveOccurred())

		})

		It("successfully retrieves an auto scale policy", func() {
			scalePolicy, err := scalePolicyService.GetObject(scalePolicy.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(scalePolicy.ScaleGroupId).To(Equal(1234567))
			Expect(scalePolicy.Name).To(Equal("fake-name"))
			Expect(scalePolicy.Cooldown).To(Equal(30))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := scalePolicyService.GetObject(scalePolicy.Id)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := scalePolicyService.GetObject(scalePolicy.Id)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#DeleteObject", func() {
		BeforeEach(func() {
			scalePolicy.Id = 4567
		})

		It("successfully deletes the SoftLayer_Scale_Policy instance", func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse = []byte("true")

			deleted, err := scalePolicyService.DeleteObject(scalePolicy.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())
		})

		It("fails to delete the SoftLayer_Scale_Group instance", func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse = []byte("false")

			deleted, err := scalePolicyService.DeleteObject(scalePolicy.Id)
			Expect(err).To(HaveOccurred())
			Expect(deleted).To(BeFalse())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := scalePolicyService.DeleteObject(scalePolicy.Id)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					_, err := scalePolicyService.DeleteObject(scalePolicy.Id)
					Expect(err).To(HaveOccurred())
				}
			})
		})

	})

})
