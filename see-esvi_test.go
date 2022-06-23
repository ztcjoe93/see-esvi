package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type CliTestSuite struct {
	suite.Suite
}

func (s *CliTestSuite) SetupTest() {
	// to reset cli args
	os.Args = []string{"golang_script"}

	// to reset valFlag
	val := ""
	valFlag = &val
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

func TestCliTestSuite(t *testing.T) {
	suite.Run(t, new(CliTestSuite))
}

func (s *CliTestSuite) TestParseCommandRead() {
	cmd, err := parseCommand("read")
	assert.Nil(s.T(), err)

	// https://github.com/stretchr/testify/issues/182
	cmdFn := runtime.FuncForPC(reflect.ValueOf(cmd).Pointer()).Name()
	readFn := runtime.FuncForPC(reflect.ValueOf(readData).Pointer()).Name()

	assert.Equal(s.T(), cmdFn, readFn)
}

func (s *CliTestSuite) TestParseCommandModify() {
	str := "someVal"
	valFlag = &str
	cmd, err := parseCommand("modify")
	assert.Nil(s.T(), err)

	cmdFn := runtime.FuncForPC(reflect.ValueOf(cmd).Pointer()).Name()
	modifyFn := runtime.FuncForPC(reflect.ValueOf(modifyData).Pointer()).Name()

	assert.Equal(s.T(), cmdFn, modifyFn)
}

func (s *CliTestSuite) TestParseCommandModifyNoValFlag() {
	_, err := parseCommand("modify")
	fmt.Println(err)
	if assert.Error(s.T(), err) {
		assert.Equal(s.T(), errors.New("no value provided to valFlag"), err)
	}
}

func (s *CliTestSuite) TestParseCommandInvalid() {
	_, err := parseCommand("asdf12!@#")
	assert.NotNil(s.T(), err)
}

func (s *CliTestSuite) TestCliArgParseNoPathProvided() {
	os.Args = []string{"golang_script", "read"}
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}

func (s *CliTestSuite) TestCliArgParsePathProvided() {
	os.Args = []string{"golang_script", "read", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		_, path := cliArgParse()

		assert.Equal(s.T(), "someFilePath", path)
		assert.False(s.T(), *isRecursive)
		assert.Equal(s.T(), 0, targetField)
	})
}

func (s *CliTestSuite) TestCliArgMultipleParsePathProvided() {
	os.Args = []string{"golang_script", "read", "someFilePath", "anotherFilePath"}
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}

func (s *CliTestSuite) TestCliArgParseNoArgs() {
	os.Args = []string{"golang_script"}
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}

func (s *CliTestSuite) TestCliArgParseInvalidCommand() {
	os.Args = []string{"golang_script", "boomBoomPow", "someFilePath"}
	assert.Panics(s.T(), func() {
		cliArgParse()
	})
}

func (s *CliTestSuite) TestCliArgParsePathRecursive() {
	os.Args = []string{"golang_script", "-r", "read", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.True(s.T(), *isRecursive)
	})
}

func (s *CliTestSuite) TestCliArgParseTargetFieldInt() {
	os.Args = []string{"golang_script", "-tf=5", "read", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.Equal(s.T(), 5, targetField)
	})
}

func (s *CliTestSuite) TestCliArgParseTargetFieldString() {
	os.Args = []string{"golang_script", "-tf=someHeader", "read", "someFilePath"}
	assert.NotPanics(s.T(), func() {
		cliArgParse()
		assert.Equal(s.T(), "someHeader", targetField)
	})
}

func (s *CliTestSuite) TestInitializeNewData() {

	path := "./somePath"
	csvRecord := [][]string{{"header_1", "header_2"},
		{"someVal", "somePal"}, {"cat", "dog"}}

	dataStructure := newData(path, csvRecord)

	assert.Equal(s.T(), "somePath", dataStructure.name)
	assert.Equal(s.T(), []string{"header_1", "header_2"}, dataStructure.headers)
	assert.Equal(s.T(), [][]string{{"someVal", "somePal"}, {"cat", "dog"}},
		dataStructure.values)
}
