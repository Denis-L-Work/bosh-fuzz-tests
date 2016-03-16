package analyzer

import (
	bftexpectation "github.com/cloudfoundry-incubator/bosh-fuzz-tests/expectation"
	bftinput "github.com/cloudfoundry-incubator/bosh-fuzz-tests/input"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Analyzer interface {
	Analyze(inputs []bftinput.Input) []Case
}

type Case struct {
	Input        		bftinput.Input
	Expectations 		[]bftexpectation.Expectation
	DeploymentWillFail	bool
}

type analyzer struct {
	stemcellComparator       Comparator
	nothingChangedComparator Comparator
}

func NewAnalyzer(logger boshlog.Logger) Analyzer {
	return &analyzer{
		stemcellComparator:       NewStemcellComparator(logger),
		nothingChangedComparator: NewNothingChangedComparator(),
	}
}

func (a *analyzer) Analyze(inputs []bftinput.Input) []Case {
	cases := []Case{}
	for i := range inputs {
		expectations := []bftexpectation.Expectation{}
		deploymentWillFail := false

		if i != 0 {
			expectations = append(expectations, a.stemcellComparator.Compare(inputs[:i], inputs[i])...)
			expectations = append(expectations, a.nothingChangedComparator.Compare(inputs[:i], inputs[i])...)

			deploymentWillFail = a.isMigratingFromAzsToNoAzsAndReusingStaticIps(inputs[i - 1], inputs[i])
		}

		cases = append(cases, Case{
			Input:				inputs[i],
			Expectations: 		expectations,
			DeploymentWillFail: deploymentWillFail,
		})
	}

	return cases
}

func (a *analyzer) isMigratingFromAzsToNoAzsAndReusingStaticIps(previousInput bftinput.Input, currentInput bftinput.Input) bool {
	for _, job := range currentInput.Jobs {
		previousJob, found := previousInput.FindJobByName(job.Name)
		if found && (len(previousJob.AvailabilityZones) > 0 && len(job.AvailabilityZones) == 0) {
			for _, network := range job.Networks {
				previousNetwork, networkFound := previousJob.FindNetworkByName(network.Name)
				if networkFound {
					for _, currentIP := range network.StaticIps {
						for _, prevIP := range previousNetwork.StaticIps {
							if prevIP == currentIP {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}
