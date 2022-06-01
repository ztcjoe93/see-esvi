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
	initialArgs []string
}

func (s *CliTestSuite) SetupTest() {
	if s.initialArgs == nil {
		s.initialArgs = os.Args
	}
	os.Args = s.initialArgs
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
	os.Args = append(os.Args, "somefilepath")
	cliArgParse()
	assert.NotPanics(s.T(), func() {
	})
}

func (s *CliTestSuite) TestCliArgParseNoPathProvided() {
	fmt.Println(os.Args)
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}
