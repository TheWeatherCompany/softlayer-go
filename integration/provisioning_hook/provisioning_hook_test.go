package provisioning_hook_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	datatypes "github.com/TheWeatherCompany/softlayer-go/data_types"
	"github.com/TheWeatherCompany/softlayer-go/softlayer"
	testhelpers "github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("SoftLayer Provisioning Hook", func() {
	var (
		err                     error
		provisioningHookService softlayer.SoftLayer_Provisioning_Hook_Service
	)

	BeforeEach(func() {
		provisioningHookService, err = testhelpers.CreateProvisioningHookService()
		Expect(err).ToNot(HaveOccurred())

		testhelpers.TIMEOUT = 30 * time.Second
		testhelpers.POLLING_INTERVAL = 10 * time.Second
	})

	Context("SoftLayer_Provisioning_Hook", func() {
		It("creates a SoftLayer Provisioning Hook, updates it and deletes it", func() {
			createdProvisioningHook, _ := testhelpers.CreateTestProvisioningHook()

			testhelpers.WaitForCreatedProvisioningHookToBePresent(createdProvisioningHook.Id)

			provisioningHookService, err := testhelpers.CreateProvisioningHookService()
			Expect(err).ToNot(HaveOccurred())

			result, err := provisioningHookService.GetObject(createdProvisioningHook.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(result.Uri).To(Equal("http://www.weather.com"))
			Expect(result.Name).To(Equal("TWCTestHook"))
			Expect(result.TypeId).To(BeNumerically("==", 1))

			provisioningHookEdit := datatypes.SoftLayer_Provisioning_Hook_Template{
				Name: "TWCEditTestHook",
				Uri:  "http://ww3.weather.com",
			}

			provisioningHookService.EditObject(createdProvisioningHook.Id, provisioningHookEdit)

			editResult, err := provisioningHookService.GetObject(createdProvisioningHook.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.CreateDate).To(Equal(editResult.CreateDate))
			Expect(editResult.Uri).To(Equal("http://ww3.weather.com"))
			Expect(editResult.Name).To(Equal("TWCEditTestHook"))
			Expect(editResult.ModifyDate).ToNot(BeNil())

			deleted, err := provisioningHookService.DeleteObject(createdProvisioningHook.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())
		})
	})
})
