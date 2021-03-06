package analyzer_test

import (
	bftexpectation "github.com/cloudfoundry-incubator/bosh-fuzz-tests/expectation"
	bftinput "github.com/cloudfoundry-incubator/bosh-fuzz-tests/input"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/cloudfoundry-incubator/bosh-fuzz-tests/analyzer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StemcellComparator", func() {
	var (
		stemcellComparator Comparator
		previousInputs     []bftinput.Input
		currentInput       bftinput.Input
	)

	BeforeEach(func() {
		logger := boshlog.NewLogger(boshlog.LevelNone)
		stemcellComparator = NewStemcellComparator(logger)
	})

	Context("when there are same instance groups that have different stemcell versions using vm types", func() {
		BeforeEach(func() {
			previousInputs = []bftinput.Input{
				{
					Stemcells: []bftinput.StemcellConfig{
						{
							Alias:   "fake-stemcell",
							Version: "1",
						},
					},
					InstanceGroups: []bftinput.InstanceGroup{
						{
							Name:     "foo-instance-group",
							Stemcell: "fake-stemcell",
						},
					},
				},
			}

			currentInput = bftinput.Input{
				Stemcells: []bftinput.StemcellConfig{
					{
						Alias:   "fake-stemcell",
						Version: "2",
					},
				},
				InstanceGroups: []bftinput.InstanceGroup{
					{
						Name:     "foo-instance-group",
						Stemcell: "fake-stemcell",
					},
				},
			}
		})

		It("returns debug log expectation", func() {
			expectations := stemcellComparator.Compare(previousInputs, currentInput)
			expectedDebugLogExpectation := bftexpectation.NewExistingInstanceDebugLog("stemcell_changed?", "foo-instance-group")
			Expect(expectations).To(ContainElement(expectedDebugLogExpectation))
		})
	})

	Context("when there are same instance groups that have different stemcell versions using resource pools", func() {
		BeforeEach(func() {
			previousInputs = []bftinput.Input{
				{
					CloudConfig: bftinput.CloudConfig{
						ResourcePools: []bftinput.ResourcePoolConfig{
							{
								Name: "fake-resource-pool",
								Stemcell: bftinput.StemcellConfig{
									Name:    "fake-stemcell",
									Version: "1",
								},
							},
							{
								Name: "fake-same-pool",
								Stemcell: bftinput.StemcellConfig{
									Name:    "fake-stemcell",
									Version: "1",
								},
							},
						},
					},
					InstanceGroups: []bftinput.InstanceGroup{
						{
							Name:         "foo-instance-group",
							ResourcePool: "fake-resource-pool",
						},
						{
							Name:         "another-instance-group",
							ResourcePool: "fake-same-pool",
						},
					},
				},
			}

			currentInput = bftinput.Input{
				CloudConfig: bftinput.CloudConfig{
					ResourcePools: []bftinput.ResourcePoolConfig{
						{
							Name: "fake-resource-pool",
							Stemcell: bftinput.StemcellConfig{
								Name:    "fake-stemcell",
								Version: "2",
							},
						},
						{
							Name: "fake-same-pool",
							Stemcell: bftinput.StemcellConfig{
								Name:    "fake-stemcell",
								Version: "1",
							},
						},
					},
				},
				InstanceGroups: []bftinput.InstanceGroup{
					{
						Name:         "foo-instance-group",
						ResourcePool: "fake-resource-pool",
					},
					{
						Name:         "another-instance-group",
						ResourcePool: "fake-same-pool",
					},
				},
			}
		})

		It("returns debug log expectation", func() {
			expectations := stemcellComparator.Compare(previousInputs, currentInput)
			expectedDebugLogExpectation := bftexpectation.NewExistingInstanceDebugLog("stemcell_changed?", "foo-instance-group")
			Expect(expectations).To(ContainElement(expectedDebugLogExpectation))
		})
	})

	Context("when switching from vm type (stemcell on InstanceGroup) to resource pool", func() {
		BeforeEach(func() {
			previousInputs = []bftinput.Input{
				{
					Stemcells: []bftinput.StemcellConfig{
						{
							Alias:   "fake-stemcell",
							Version: "1",
						},
					},
					InstanceGroups: []bftinput.InstanceGroup{
						{
							Name:     "foo-instance-group",
							Stemcell: "fake-stemcell",
						},
					},
				},
			}

			currentInput = bftinput.Input{
				CloudConfig: bftinput.CloudConfig{
					ResourcePools: []bftinput.ResourcePoolConfig{
						{
							Name: "fake-resource-pool",
							Stemcell: bftinput.StemcellConfig{
								OS:      "toronto-os",
								Version: "1",
							},
						},
					},
				},
				InstanceGroups: []bftinput.InstanceGroup{
					{
						Name:         "foo-instance-group",
						ResourcePool: "fake-resource-pool",
					},
				},
			}
		})

		It("returns debug log expectation", func() {
			expectations := stemcellComparator.Compare(previousInputs, currentInput)
			Expect(expectations).To(BeEmpty())
		})
	})
})
