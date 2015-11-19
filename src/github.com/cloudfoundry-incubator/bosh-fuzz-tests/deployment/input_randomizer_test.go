package deployment_test

import (
	bftconfig "github.com/cloudfoundry-incubator/bosh-fuzz-tests/config"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	. "github.com/cloudfoundry-incubator/bosh-fuzz-tests/deployment"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InputRandomizer", func() {
	var (
		inputRandomizer InputRandomizer
	)

	It("generates extra input for migrated jobs", func() {
		parameters := bftconfig.Parameters{
			NameLength:               []int{5, 10},
			Instances:                []int{2, 4},
			AvailabilityZones:        [][]string{[]string{"z1"}, []string{"z1", "z2"}},
			PersistentDiskDefinition: []string{"disk_pool", "disk_type", "persistent_disk_size"},
			PersistentDiskSize:       []int{0, 100, 200},
			NumberOfJobs:             []int{1, 2},
			MigratedFromCount:        []int{0, 2},
		}
		logger := boshlog.NewLogger(boshlog.LevelNone)
		inputRandomizer = NewSeededInputRandomizer(parameters, 2, 64, logger)

		inputs, err := inputRandomizer.Generate()
		Expect(err).ToNot(HaveOccurred())

		Expect(inputs).To(Equal([]Input{
			{
				Jobs: []Job{
					{
						Name:              "joNAw",
						Instances:         4,
						AvailabilityZones: []string{"z1"},
						Network:           "default",
					},
					{
						Name:              "gQ8el",
						Instances:         4,
						AvailabilityZones: []string{"z1", "z2"},
						Network:           "default",
					},
				},
				CloudConfig: CloudConfig{
					AvailabilityZones: []string{"z1", "z2"},
				},
			},
			{
				Jobs: []Job{
					{
						Name:               "rU3YND0xNg",
						Instances:          4,
						AvailabilityZones:  []string{"z1"},
						PersistentDiskPool: "gBUnQKBYoE",
						Network:            "default",
					},
					{
						Name:               "pRWDsiO5Qu",
						Instances:          4,
						AvailabilityZones:  []string{"z1"},
						PersistentDiskPool: "a5gmsYqE7Y",
						Network:            "default",
					},
				},
				CloudConfig: CloudConfig{
					AvailabilityZones: []string{"z1"},
					PersistentDiskPools: []DiskConfig{
						{Name: "gBUnQKBYoE", Size: 100},
						{Name: "a5gmsYqE7Y", Size: 100},
					},
				},
			},
			{
				Jobs: []Job{
					{
						Name:               "joNAw",
						Instances:          4,
						AvailabilityZones:  []string{"z1"},
						PersistentDiskPool: "eagRjDTBs3",
						Network:            "default",
						MigratedFrom: []MigratedFromConfig{
							{Name: "rU3YND0xNg"},
							{Name: "pRWDsiO5Qu"},
						},
					},
				},
				CloudConfig: CloudConfig{
					AvailabilityZones: []string{"z1"},
					PersistentDiskPools: []DiskConfig{
						{Name: "eagRjDTBs3", Size: 100},
					},
				},
			},
		}))
	})

	// Pending until #108499370
	PIt("when migrated job does not have az it sets random az in migrated_from", func() {
		parameters := bftconfig.Parameters{
			NameLength:               []int{5},
			Instances:                []int{2},
			AvailabilityZones:        [][]string{[]string{"z1"}, nil},
			PersistentDiskDefinition: []string{"persistent_disk_size"},
			PersistentDiskSize:       []int{0},
			NumberOfJobs:             []int{1},
			MigratedFromCount:        []int{1},
		}
		logger := boshlog.NewLogger(boshlog.LevelNone)
		inputRandomizer = NewSeededInputRandomizer(parameters, 1, 64, logger)

		inputs, err := inputRandomizer.Generate()
		Expect(err).ToNot(HaveOccurred())

		Expect(inputs).To(Equal([]Input{
			{
				Jobs: []Job{
					{
						Name:      "vgrKicN3O2",
						Instances: 2,
						Network:   "no-az",
					},
				},
			},
			{
				Jobs: []Job{
					{
						Name:              "joNAw",
						Instances:         2,
						Network:           "default",
						AvailabilityZones: []string{"z1"},
						MigratedFrom: []MigratedFromConfig{
							{Name: "vgrKicN3O2", AvailabilityZone: "z1"},
						},
					},
				},
				CloudConfig: CloudConfig{
					AvailabilityZones: []string{"z1"},
				},
			},
		}))
	})
})
