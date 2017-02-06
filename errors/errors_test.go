// +build !acceptance

package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (suite *ErrorTestSuite) TestResourceNotFoundError() {

	// invoke
	var result = suite.resourceNotFoundError.Error()

	// assert
	assert.Equal(suite.T(), ErrResourceNotFound+"|"+suite.resourceNotFoundError.Message+"|"+suite.resourceNotFoundError.RequestID, result)
}

func (suite *ErrorTestSuite) TestUndefinedError() {

	// invoke
	var result = suite.undefinedError.Error()

	// assert
	assert.Equal(suite.T(), ErrUndefined+"|"+suite.undefinedError.Err.Error(), result)
}

func (suite *ErrorTestSuite) TestBindError() {

	// invoke
	var result = suite.bindError.Error()

	// assert
	assert.Equal(suite.T(), ErrBind+"|"+suite.bindError.Err.Error()+". "+suite.bindError.Message+"|"+suite.bindError.RequestID, result)
}

func (suite *ErrorTestSuite) TestLdapServerFatalError() {

	// invoke
	var result = suite.ldapServerFatalError.Error()

	// assert
	assert.Equal(suite.T(), ErrLDAPServerFatal+"|"+suite.ldapServerFatalError.Err.Error()+". "+suite.ldapServerFatalError.Message+"|"+suite.ldapServerFatalError.RequestID, result)
}

func (suite *ErrorTestSuite) TestAuthenticationError() {

	// invoke
	var result = suite.authenticationError.Error()

	// assert
	assert.Equal(suite.T(), ErrAuthentication+"|"+suite.authenticationError.Err.Error()+". "+suite.authenticationError.Message+"|"+suite.authenticationError.RequestID, result)
}

func (suite *ErrorTestSuite) TestPrivateKeyError() {

	// invoke
	var result = suite.privateKeyError.Error()

	// assert
	assert.Equal(suite.T(), ErrPrivateKey+"|"+suite.privateKeyError.Err.Error()+". "+suite.privateKeyError.Message+"|"+suite.privateKeyError.RequestID, result)
}

func (suite *ErrorTestSuite) TestPublicKeyError() {

	// invoke
	var result = suite.publicKeyError.Error()

	// assert
	assert.Equal(suite.T(), ErrPublicKey+"|"+suite.publicKeyError.Err.Error()+". "+suite.publicKeyError.Message+"|"+suite.publicKeyError.RequestID, result)
}

func (suite *ErrorTestSuite) TestSignError() {

	// invoke
	var result = suite.signError.Error()

	// assert
	assert.Equal(suite.T(), ErrSign+"|"+suite.signError.Err.Error()+". "+suite.signError.Message+"|"+suite.signError.RequestID, result)
}

func (suite *ErrorTestSuite) TestContentTypeError() {

	// invoke
	var result = suite.contentTypeError.Error()

	// assert
	assert.Equal(suite.T(), ErrUnexpectedContentType+"|"+suite.contentTypeError.Err.Error()+"|"+suite.contentTypeError.RequestID, result)
}

func (suite *ErrorTestSuite) TestMysqlFatalError() {

	// invoke
	var result = suite.mysqlFatalError.Error()

	// assert
	assert.Equal(suite.T(), ErrMySQLFatal+"|"+suite.mysqlFatalError.Err.Error()+"|"+suite.mysqlFatalError.RequestID, result)
}

func (suite *ErrorTestSuite) TestUnssuportedAlgorithmError() {

	// invoke
	var result = suite.unssuportedAlgorithmError.Error()

	// assert
	assert.Equal(suite.T(), ErrUnsupportedAuthAlgorithm+"|"+suite.unssuportedAlgorithmError.Err.Error()+"|"+suite.unssuportedAlgorithmError.RequestID, result)
}

type ErrorTestSuite struct {
	suite.Suite
	resourceNotFoundError     *ResourceNotFoundError
	undefinedError            *UndefinedError
	bindError                 *BindError
	ldapServerFatalError      *LDAPServerFatalError
	authenticationError       *AuthenticationError
	privateKeyError           *PrivateKeyError
	publicKeyError            *PublicKeyError
	signError                 *SignError
	contentTypeError          *ContentTypeError
	mysqlFatalError           *MySQLFatalError
	unssuportedAlgorithmError *UnsupportedAlgorithmError
}

func (suite *ErrorTestSuite) SetupTest() {
	suite.resourceNotFoundError = &ResourceNotFoundError{Err: errors.New("db is down"), Message: "datastore is down", RequestID: "MockID"}
	suite.undefinedError = &UndefinedError{errors.New("undefined")}
	suite.bindError = &BindError{Err: errors.New("bind error"), Message: "my custom bind msg", RequestID: "MockID"}
	suite.ldapServerFatalError = &LDAPServerFatalError{Err: errors.New("ldap server error"), Message: "ldap is down", RequestID: "MockID"}
	suite.authenticationError = &AuthenticationError{Err: errors.New("authentication error"), Message: "custom msg", RequestID: "MockID"}
	suite.privateKeyError = &PrivateKeyError{Err: errors.New("private key error"), Message: "custom msg", RequestID: "MockID"}
	suite.publicKeyError = &PublicKeyError{Err: errors.New("public key error"), Message: "custom msg", RequestID: "MockID"}
	suite.signError = &SignError{Err: errors.New("sign error"), Message: "custom msg", RequestID: "MockID"}
	suite.contentTypeError = &ContentTypeError{Err: errors.New("wrong content type"), RequestID: "MockID"}
	suite.mysqlFatalError = &MySQLFatalError{Err: errors.New("fatal mysql error"), RequestID: "MockID"}
	suite.unssuportedAlgorithmError = &UnsupportedAlgorithmError{Err: errors.New("unsupported algorithm"), RequestID: "MockID"}
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}
