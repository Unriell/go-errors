// +build !acceptance

package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (suite *ErrorHandlerTestSuite) TestUndefinedError() {

	// invoke
	result := suite.errHandler.Error(errors.New("UNDEFINED|my fake msg|mockID"), nil)

	// assert
	assert.Equal(suite.T(), "BQ0000000", result.ID)
	assert.Equal(suite.T(), "Error Not found", result.Msg)
	assert.Equal(suite.T(), 500, result.Status)
	assert.Equal(suite.T(), "this error doesn't exist. Msg UNDEFINED|my fake msg|mockID", result.ComponentMsg)
}

func (suite *ErrorHandlerTestSuite) TestEmptyError() {

	// invoke
	result := suite.errHandler.Error(errors.New(""), nil)

	// assert
	assert.Equal(suite.T(), "BQ0000000", result.ID)
	assert.Equal(suite.T(), "Error Not found", result.Msg)
	assert.Equal(suite.T(), 500, result.Status)
	assert.Equal(suite.T(), "", "")
	assert.Equal(suite.T(), "Undefined error code. No error code found", result.ComponentMsg)
}

func (suite *ErrorHandlerTestSuite) TestSuccessError() {

	dst := make(map[string]string, 3)
	dst["URL"] = "www.google.com"
	dst["os"] = "android"
	dst["tagTest"] = "test tag"
	dst["extraTag"] = "extra test tag"

	// invoke
	result := suite.errHandler.Error(errors.New("BQ0000001|my error msg|mockID"), dst)

	// assert
	assert.Equal(suite.T(), "BQ0000001", result.ID)
	assert.Equal(suite.T(), "Invalid inbound entity", result.Msg)
	assert.Equal(suite.T(), 400, result.Status)
	assert.Equal(suite.T(), "mockID", result.SentryCode)
	assert.Equal(suite.T(), "my error msg", result.ComponentMsg)
}

func (suite *ErrorHandlerTestSuite) TestErrorWithoutID() {

	dst := make(map[string]string, 3)
	dst["URL"] = "www.google.com"
	dst["os"] = "android"
	dst["tagTest"] = "test tag"
	dst["extraTag"] = "extra test tag"

	// invoke
	result := suite.errHandler.Error(errors.New("BQ0000001|my error msg"), dst)

	// assert
	assert.Equal(suite.T(), "BQ0000001", result.ID)
	assert.Equal(suite.T(), "Invalid inbound entity", result.Msg)
	assert.Equal(suite.T(), 400, result.Status)
	assert.Equal(suite.T(), "", result.SentryCode)
	assert.Equal(suite.T(), "my error msg", result.ComponentMsg)
}

type ErrorHandlerTestSuite struct {
	suite.Suite
	errHandler *ErrorHandler
}

func (suite *ErrorHandlerTestSuite) SetupTest() {
	suite.errHandler = NewErrorHandler("QA", "https://1eecf0d5795b4f9bbc5f4210c1513f4c:35ae8d2ab04144d8a01127612ed810d9@sentry.bq.com/132", "mock", true)
}

func TestErrorHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorHandlerTestSuite))
}
