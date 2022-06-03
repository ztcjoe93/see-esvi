package main

import (
	"fmt"
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

func (s *CliTestSuite) TestCliArgParsePathProvided() {
	os.Args = []string{"golang_script", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
	})
}

func (s *CliTestSuite) TestCliArgParsePathRecursive() {
	os.Args = []string{"golang_script", "-r=true", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.True(s.T(), *isRecursive)
	})
}

func (s *CliTestSuite) TestCliArgParseNoPathProvided() {
	fmt.Println(os.Args)
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}
