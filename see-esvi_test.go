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
	s.initialArgs = os.Args
}

func (s *CliTestSuite) BeforeTest() {
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

func (s *CliTestSuite) AfterTest() {
	os.Args = s.initialArgs
}

func TestCliArgParsePathProvided(t *testing.T) {
	os.Args = append(os.Args, "test")
	fmt.Println(os.Args)
	assert.NotPanics(t, func() {
		cliArgParse()
	})
}

func TestCliArgParseNoPathProvided(t *testing.T) {
	assert.Panics(t, func() {
		cliArgParse()
	})
}
