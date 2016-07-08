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

var _ = Describe("SoftLayer_Scale_Group_Service", func() {
	var (
		username, apiKey string
		err              error

		fakeClient *slclientfakes.FakeSoftLayerClient

		scaleGroupService softlayer.SoftLayer_Scale_Group_Service

		scaleGroup         datatypes.SoftLayer_Scale_Group
		scaleGroupTemplate datatypes.SoftLayer_Scale_Group
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		scaleGroupService, err = fakeClient.GetSoftLayer_Scale_Group_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(scaleGroupService).ToNot(BeNil())

		scaleGroup = datatypes.SoftLayer_Scale_Group{}
		scaleGroupTemplate = datatypes.SoftLayer_Scale_Group{}
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {

		})
	})

	Context("#CreateObject", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Scale_Group_Service_createAndGetObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new SoftLayer_Scale_Group instance", func() {

			scaleGroupTemplate = datatypes.SoftLayer_Scale_Group{
				Name:               "fake-name",
				RegionalGroupId:    1234,
				MinimumMemberCount: 1,
				MaximumMemberCount: 10,
				Cooldown:           30,
			}
			scaleGroup, err = scaleGroupService.CreateObject(scaleGroupTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(scaleGroup.Name).To(Equal("fake-name"))
			Expect(scaleGroup.RegionalGroupId).To(Equal(1234))
			Expect(scaleGroup.MinimumMemberCount).To(Equal(1))
			Expect(scaleGroup.Cooldown).To(Equal(30))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					scaleGroupTemplate = datatypes.SoftLayer_Scale_Group{
						Name:               "fake-name",
						RegionalGroupId:    1234,
						MinimumMemberCount: 1,
						MaximumMemberCount: 10,
						Cooldown:           30,
					}
					_, err = scaleGroupService.CreateObject(scaleGroupTemplate)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					scaleGroupTemplate = datatypes.SoftLayer_Scale_Group{
						Name:               "fake-name",
						RegionalGroupId:    1234,
						MinimumMemberCount: 1,
						MaximumMemberCount: 10,
						Cooldown:           30,
					}
					_, err = scaleGroupService.CreateObject(scaleGroupTemplate)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			scaleGroup.Id = 1234567
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Scale_Group_Service_createAndGetObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("successfully retrieves an auto scale group", func() {
			mask := []string{"cooldown"}
			scaleGroup, err := scaleGroupService.GetObject(scaleGroup.Id, mask)
			Expect(err).ToNot(HaveOccurred())
			Expect(scaleGroup.Name).To(Equal("fake-name"))
			Expect(scaleGroup.RegionalGroupId).To(Equal(1234))
			Expect(scaleGroup.MinimumMemberCount).To(Equal(1))
			Expect(scaleGroup.Cooldown).To(Equal(30))
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					
					mask := []string{"cooldown"}
					_, err := scaleGroupService.GetObject(scaleGroup.Id, mask)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode
					
					mask := []string{"cooldown"}
					_, err := scaleGroupService.GetObject(scaleGroup.Id, mask)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

	Context("#EditObject", func() {
		BeforeEach(func() {
			scaleGroup.Id = 1234567
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("services", "SoftLayer_Scale_Group_Service_editObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("edits an existing SoftLayer_Scale_Group instance", func() {
			edited := datatypes.SoftLayer_Scale_Group{
				Name: "edited-name",
			}
			result, err := scaleGroupService.EditObject(scaleGroup.Id, edited)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("when HTTP client returns error codes 40x or 50x", func() {
			It("fails for error code 40x", func() {
				errorCodes := []int{400, 401, 499}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					edited := datatypes.SoftLayer_Scale_Group{
						Name: "edited-name",
					}
					_, err := scaleGroupService.EditObject(scaleGroup.Id, edited)
					Expect(err).To(HaveOccurred())
				}
			})

			It("fails for error code 50x", func() {
				errorCodes := []int{500, 501, 599}
				for _, errorCode := range errorCodes {
					fakeClient.FakeHttpClient.DoRawHttpRequestInt = errorCode

					edited := datatypes.SoftLayer_Scale_Group{
						Name: "edited-name",
					}
					_, err := scaleGroupService.EditObject(scaleGroup.Id, edited)
					Expect(err).To(HaveOccurred())
				}
			})
		})
	})

})
