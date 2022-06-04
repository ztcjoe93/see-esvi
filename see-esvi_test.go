package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type CliTestSuite struct {
	suite.Suite
}

func (s *CliTestSuite) SetupTest() {
	os.Args = []string{"golang_script"}
	cfg := zap.NewDevelopmentConfig()
	_, err := os.Stat(".")
	if err != nil {
		panic(err)
	}

	cfg.OutputPaths = []string{"stderr"}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar = logger.Sugar()
}

func TestCliTestSuite(t *testing.T) {
	suite.Run(t, new(CliTestSuite))
}

func (s *CliTestSuite) TestCliArgParseNoPathProvided() {
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}

func (s *CliTestSuite) TestCliArgParsePathProvided() {
	os.Args = []string{"golang_script", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.False(s.T(), *isRecursive)
		assert.Equal(s.T(), 0, targetField)
	})
}

func (s *CliTestSuite) TestCliArgMultipleParsePathProvided() {
	os.Args = []string{"golang_script", "someFilePath", "anotherFilePath"}
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}

func (s *CliTestSuite) TestCliArgParsePathRecursive() {
	os.Args = []string{"golang_script", "-r", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.True(s.T(), *isRecursive)
	})
}

func (s *CliTestSuite) TestCliArgParseTargetFieldInt() {
	os.Args = []string{"golang_script", "-tf=5", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.Equal(s.T(), 5, targetField)
	})
}

func (s *CliTestSuite) TestCliArgParseTargetFieldString() {
	os.Args = []string{"golang_script", "-tf=someHeader", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.Equal(s.T(), "someHeader", targetField)
	})
}
