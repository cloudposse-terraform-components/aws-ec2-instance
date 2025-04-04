package test

import (
	"context"
	"testing"

	"github.com/cloudposse/test-helpers/pkg/atmos"
	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
	awshelper "github.com/cloudposse/test-helpers/pkg/aws"
	"github.com/stretchr/testify/assert"
)

type ComponentSuite struct {
	helper.TestSuite
}

func (s *ComponentSuite) TestBasic() {
	const component = "ec2-instance/basic"
	const stack = "default-test"
	const awsRegion = "us-east-2"

	defer s.DestroyAtmosComponent(s.T(), component, stack, nil)
	options, _ := s.DeployAtmosComponent(s.T(), component, stack, nil)
	assert.NotNil(s.T(), options)

	instanceIds := atmos.OutputList(s.T(), options, "instance_id")
	assert.NotEmpty(s.T(), instanceIds)

	privateIp := atmos.OutputList(s.T(), options, "private_ip")
	assert.EqualValues(s.T(), 1, len(privateIp))

	instance := awshelper.GetEc2Instances(s.T(), context.Background(), instanceIds[0], awsRegion)
	assert.EqualValues(s.T(), "t3a.micro", instance.InstanceType)
	assert.EqualValues(s.T(), "running", instance.State.Name)

	s.DriftTest(component, stack, nil)
}

func (s *ComponentSuite) TestEnabledFlag() {
	const component = "ec2-instance/disabled"
	const stack = "default-test"
	s.VerifyEnabledFlag(component, stack, nil)
}

func TestRunSuite(t *testing.T) {
	suite := new(ComponentSuite)

	suite.AddDependency(t, "vpc", "default-test", nil)

	helper.Run(t, suite)
}
