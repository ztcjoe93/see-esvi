package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilitiesTestSuite struct {
	suite.Suite
}

func (s *UtilitiesTestSuite) SetupTest() {
}

func TestUtilitiesTestSuite(t *testing.T) {
	suite.Run(t, new(UtilitiesTestSuite))
}

func (s *UtilitiesTestSuite) TestTypeOfUtilityStringArg() {
	var data interface{} = "someString"

	returnVal := typeof(data)
	assert.Equal(s.T(), "string", returnVal)
}

func (s *UtilitiesTestSuite) TestTypeOfUtilityIntArg() {
	var data interface{} = 123

	returnVal := typeof(data)
	assert.Equal(s.T(), "int", returnVal)
}

func (s *UtilitiesTestSuite) TestTypeOfUtilityOtherArg() {
	var data interface{} = make(map[int]string)

	returnVal := typeof(data)
	assert.Equal(s.T(), "unknown", returnVal)
}
