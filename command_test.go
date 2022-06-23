package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type CommandTestSuite struct {
	suite.Suite
}

func (s *CommandTestSuite) SetupTest() {
	// to reset cli args
	os.Args = []string{"golang_script"}
	targetField = 0

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

func TestCommandTestSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}

func (s *CommandTestSuite) TestGetTargetFieldString() {
	os.Args = []string{"golang_script", "-tf=someHeader", "read", "someFilePath"}
	cliArgParse()

	dataSlice = []*Data{
		{
			name: "test",
			headers: []string{
				"header1", "header2", "someHeader",
			},
			values: [][]string{
				{"val1", "val2", "val3"},
			},
		},
	}
	targetIndex := getTargetField()
	assert.Equal(s.T(), 2, targetIndex)
}

func (s *CommandTestSuite) TestGetTargetFieldInt() {
	os.Args = []string{"golang_script", "-tf=1", "read", "someFilePath"}
	cliArgParse()

	targetIndex := getTargetField()
	assert.Equal(s.T(), 1, targetIndex)
}
