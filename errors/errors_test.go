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
	assert.Equal(suite.T(), ErrResourceNotFound+"|"+suite.resourceNotFoundError.message+"|"+suite.resourceNotFoundError.requestID, result)
}

func (suite *ErrorTestSuite) TestUndefinedError() {

	// invoke
	var result = suite.undefinedError.Error()

	// assert
	assert.Equal(suite.T(), ErrUndefined+"|"+suite.undefinedError.err.Error(), result)
}

func (suite *ErrorTestSuite) TestBindError() {

	// invoke
	var result = suite.bindError.Error()

	// assert
	assert.Equal(suite.T(), ErrBind+"|"+suite.bindError.err.Error()+". "+suite.bindError.message+"|"+suite.bindError.requestID, result)
}

func (suite *ErrorTestSuite) TestLdapServerFatalError() {

	// invoke
	var result = suite.ldapServerFatalError.Error()

	// assert
	assert.Equal(suite.T(), ErrLDAPServerFatal+"|"+suite.ldapServerFatalError.err.Error()+". "+suite.ldapServerFatalError.message+"|"+suite.ldapServerFatalError.requestID, result)
}

func (suite *ErrorTestSuite) TestAuthenticationError() {

	// invoke
	var result = suite.authenticationError.Error()

	// assert
	assert.Equal(suite.T(), ErrAuthentication+"|"+suite.authenticationError.err.Error()+". "+suite.authenticationError.message+"|"+suite.authenticationError.requestID, result)
}

func (suite *ErrorTestSuite) TestPrivateKeyError() {

	// invoke
	var result = suite.privateKeyError.Error()

	// assert
	assert.Equal(suite.T(), ErrPrivateKey+"|"+suite.privateKeyError.err.Error()+". "+suite.privateKeyError.message+"|"+suite.privateKeyError.requestID, result)
}

func (suite *ErrorTestSuite) TestPublicKeyError() {

	// invoke
	var result = suite.publicKeyError.Error()

	// assert
	assert.Equal(suite.T(), ErrPublicKey+"|"+suite.publicKeyError.err.Error()+". "+suite.publicKeyError.message+"|"+suite.publicKeyError.requestID, result)
}

func (suite *ErrorTestSuite) TestSignError() {

	// invoke
	var result = suite.signError.Error()

	// assert
	assert.Equal(suite.T(), ErrSign+"|"+suite.signError.err.Error()+". "+suite.signError.message+"|"+suite.signError.requestID, result)
}

func (suite *ErrorTestSuite) TestContentTypeError() {

	// invoke
	var result = suite.contentTypeError.Error()

	// assert
	assert.Equal(suite.T(), ErrUnexpectedContentType+"|"+suite.contentTypeError.err.Error()+"|"+suite.contentTypeError.requestID, result)
}

func (suite *ErrorTestSuite) TestMysqlFatalError() {

	// invoke
	var result = suite.mysqlFatalError.Error()

	// assert
	assert.Equal(suite.T(), ErrMySQLFatal+"|"+suite.mysqlFatalError.err.Error()+"|"+suite.mysqlFatalError.requestID, result)
}

func (suite *ErrorTestSuite) TestUnssuportedAlgorithmError() {

	// invoke
	var result = suite.unssuportedAlgorithmError.Error()

	// assert
	assert.Equal(suite.T(), ErrUnsupportedAuthAlgorithm+"|"+suite.unssuportedAlgorithmError.err.Error()+"|"+suite.unssuportedAlgorithmError.requestID, result)
}

type ErrorTestSuite struct {
	suite.Suite
	resourceNotFoundError     *ResourceNotFoundError
	undefinedError            *UndefinedError
	bindError                 *BindError
	ldapServerFatalError      *LdapServerFatalError
	authenticationError       *AuthenticationError
	privateKeyError           *PrivateKeyError
	publicKeyError            *PublicKeyError
	signError                 *SignError
	contentTypeError          *ContentTypeError
	mysqlFatalError           *MySQLFatalError
	unssuportedAlgorithmError *UnsupportedAlgorithmError
}

func (suite *ErrorTestSuite) SetupTest() {
	suite.resourceNotFoundError = &ResourceNotFoundError{err: errors.New("db is down"), message: "datastore is down", requestID: "MockID"}
	suite.undefinedError = &UndefinedError{errors.New("undefined")}
	suite.bindError = &BindError{err: errors.New("bind error"), message: "my custom bind msg", requestID: "MockID"}
	suite.ldapServerFatalError = &LdapServerFatalError{err: errors.New("ldap server error"), message: "ldap is down", requestID: "MockID"}
	suite.authenticationError = &AuthenticationError{err: errors.New("authentication error"), message: "custom msg", requestID: "MockID"}
	suite.privateKeyError = &PrivateKeyError{err: errors.New("private key error"), message: "custom msg", requestID: "MockID"}
	suite.publicKeyError = &PublicKeyError{err: errors.New("public key error"), message: "custom msg", requestID: "MockID"}
	suite.signError = &SignError{err: errors.New("sign error"), message: "custom msg", requestID: "MockID"}
	suite.contentTypeError = &ContentTypeError{err: errors.New("wrong content type"), requestID: "MockID"}
	suite.mysqlFatalError = &MySQLFatalError{err: errors.New("fatal mysql error"), requestID: "MockID"}
	suite.unssuportedAlgorithmError = &UnsupportedAlgorithmError{err: errors.New("unsupported algorithm"), requestID: "MockID"}
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}
