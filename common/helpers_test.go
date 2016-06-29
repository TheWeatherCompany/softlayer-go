package common_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/TheWeatherCompany/softlayer-go/client/fakes"
	"github.com/TheWeatherCompany/softlayer-go/common"
	testhelpers "github.com/TheWeatherCompany/softlayer-go/test_helpers"
)

var _ = Describe("SoftlayerLookupHelpers", func() {
	var (
		username, apiKey string

		fakeClient *slclientfakes.FakeSoftLayerClient
		err        error
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())
	})

	Context("#GetDatacenterByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetDatacenterByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of datacenter", func() {
			id, err := common.GetDatacenter(fakeClient, "ams01")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(265592))
		})
	})

	Context("#GetRoutingTypesByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetRoutingTypeByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of routing type", func() {
			id, err := common.GetRoutingType(fakeClient, "DNS")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(4))
		})
	})

	Context("#GetRoutingMethodsByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetRoutingMethodByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of routing type", func() {
			id, err := common.GetRoutingMethod(fakeClient, "ROUND_ROBIN")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(10))
		})
	})

	Context("#GetHealthCheckTypesByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetHealthCheckTypeByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of health check type", func() {
			id, err := common.GetHealthCheckType(fakeClient, "DNS")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(3))
		})
	})

	Context("#GetLocationGroupRegionalsByName", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetLocationGroupRegionalByName.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves ID of location group regional", func() {
			id, err := common.GetLocationGroupRegional(fakeClient, "na-usa-west-1")
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal(62))
		})
	})

	Context("#GetDatacenterById", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetDatacenterById.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves name of datacenter", func() {
			id, err := common.GetDatacenter(fakeClient, 265592)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal("ams01"))
		})
	})

	Context("#GetRoutingTypesById", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetRoutingTypeById.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves name of routing type", func() {
			id, err := common.GetRoutingType(fakeClient, 4)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal("DNS"))
		})
	})

	Context("#GetRoutingMethodsById", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetRoutingMethodById.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves name of routing type", func() {
			id, err := common.GetRoutingMethod(fakeClient, 10)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal("ROUND_ROBIN"))
		})
	})

	Context("#GetHealthCheckTypesById", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetHealthCheckTypeById.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves name of health check type", func() {
			id, err := common.GetHealthCheckType(fakeClient, 3)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal("DNS"))
		})
	})

	Context("#GetLocationGroupRegionalsById", func() {
		BeforeEach(func() {
			fakeClient.FakeHttpClient.DoRawHttpRequestResponse, err = testhelpers.ReadJsonTestFixtures("common", "GetLocationGroupRegionalById.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves name of Location group regional", func() {
			id, err := common.GetLocationGroupRegional(fakeClient, 62)
			Expect(err).ToNot(HaveOccurred())
			Expect(id).To(Equal("na-usa-west-1"))
		})
	})
})
