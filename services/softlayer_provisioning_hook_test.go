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

var _ = Describe("SoftLayer_Provisioning_Hook_Service", func() {
	var (
		username, apiKey string
		err              error

		fakeClient *slclientfakes.FakeSoftLayerClient

		provisioningHookService softlayer.SoftLayer_Provisioning_Hook_Service

		provisioningHook         datatypes.SoftLayer_Provisioning_Hook
		provisioningHookTemplate datatypes.SoftLayer_Provisioning_Hook_Template
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		provisioningHookService, err = fakeClient.GetSoftLayer_Provisioning_Hook_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(provisioningHookService).ToNot(BeNil())

		provisioningHook = datatypes.SoftLayer_Provisioning_Hook{}
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := provisioningHookService.GetName()
			Expect(name).To(Equal("SoftLayer_Provisioning_Hook"))
		})
	})

	Context("#CreateObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Provisioning_Hook_Service_createObject.json")
			Expect(err).ToNot(HaveOccurred())

			provisioningHookTemplate = datatypes.SoftLayer_Provisioning_Hook_Template{
				Name:   "TWC-PostInstallScript",
				TypeId: 1,
				Uri:    "http://www.weather.com",
			}
		})

		It("creates a new SoftLayer_Provisioning_Hook", func() {
			provisioningHook, err = provisioningHookService.CreateObject(provisioningHookTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(provisioningHook.Name).To(Equal("TWC-PostInstallScript"))
			Expect(provisioningHook.TypeId).To(Equal(1))
			Expect(provisioningHook.Uri).To(Equal("http://www.weather.com"))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err = provisioningHookService.CreateObject(provisioningHookTemplate)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					_, err = provisioningHookService.CreateObject(provisioningHookTemplate)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#EditObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Provisioning_Hook_Service_editObject.json")

		})

		It("edits an existing provisioning hook", func() {
			provisioningHookTemplate = datatypes.SoftLayer_Provisioning_Hook_Template{
				Name: "edited-name",
				Uri:  "http://test.provision-script.net/provision.sh",
			}
			result, err := provisioningHookService.EditObject(provisioningHook.Id, provisioningHookTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					edited := datatypes.SoftLayer_Provisioning_Hook_Template{
						Name: "edited-name",
					}
					_, err := provisioningHookService.EditObject(provisioningHook.Id, edited)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					edited := datatypes.SoftLayer_Provisioning_Hook_Template{
						Name: "edited-name",
					}
					_, err := provisioningHookService.EditObject(provisioningHook.Id, edited)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#DeleteObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Provisioning_Hook_Service_deleteObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully removes the SoftLayer_Provisoining_Hook object", func() {
			result, err := provisioningHookService.DeleteObject(12345)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

	})

})
