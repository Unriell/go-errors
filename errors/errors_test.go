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

func (suite *ErrorTestSuite) TestResourceConflictError() {

	// invoke
	var result = suite.resourceConflictError.Error()

	// assert
	assert.Equal(suite.T(), ErrResourceConflict+"|"+suite.resourceConflictError.err.Error()+"|"+suite.resourceConflictError.requestID, result)
}

func (suite *ErrorTestSuite) TestRedisFatalError() {

	// invoke
	var result = suite.errRedisFatalError.Error()

	// assert
	assert.Equal(suite.T(), ErrRedisFatal+"|"+suite.errRedisFatalError.err.Error()+"|"+suite.errRedisFatalError.requestID, result)
}

func (suite *ErrorTestSuite) TestRedisUnmarshalError() {

	// invoke
	var result = suite.errRedisUnmarshalError.Error()

	// assert
	assert.Equal(suite.T(), ErrRedisUnmarshal+"|"+suite.errRedisUnmarshalError.err.Error()+"|"+suite.errRedisUnmarshalError.requestID, result)
}

func (suite *ErrorTestSuite) TestRedisNotFoundError() {

	// invoke
	var result = suite.errRedisNotfoundError.Error()

	// assert
	assert.Equal(suite.T(), ErrRedisNotFound+"|"+suite.errRedisNotfoundError.err.Error()+"|"+suite.errRedisNotfoundError.requestID, result)
}

func (suite *ErrorTestSuite) TestAuthorizationError() {

	// invoke
	var result = suite.errAuthorizationError.Error()

	// assert
	assert.Equal(suite.T(), ErrAuthorization+"|"+suite.errAuthorizationError.err.Error()+". "+suite.errAuthorizationError.message+"|"+suite.errAuthorizationError.requestID, result)
}

func (suite *ErrorTestSuite) TestInvalidAuthorizationTokenError() {

	// invoke
	var result = suite.errInvalidAuthorizationTokenError.Error()

	// assert
	assert.Equal(suite.T(), ErrInvalidAuthorizationToken+"|"+suite.errInvalidAuthorizationTokenError.err.Error()+". "+suite.errInvalidAuthorizationTokenError.message+"|"+suite.errInvalidAuthorizationTokenError.requestID, result)
}

func (suite *ErrorTestSuite) TestTokenExpiresError() {

	// invoke
	var result = suite.errTokenExpiresError.Error()

	// assert
	assert.Equal(suite.T(), ErrTokenExpires+"|"+suite.errTokenExpiresError.err.Error()+". "+suite.errTokenExpiresError.message+"|"+suite.errTokenExpiresError.requestID, result)
}

func (suite *ErrorTestSuite) TestNotificationConflictError() {

	// invoke
	var result = suite.notificationConflictError.Error()

	// assert
	assert.Equal(suite.T(), ErrNotificationConflict+"|"+suite.notificationConflictError.err.Error()+"|"+suite.notificationConflictError.requestID, result)
}

type ErrorTestSuite struct {
	suite.Suite
	resourceNotFoundError             *ResourceNotFoundError
	undefinedError                    *UndefinedError
	bindError                         *BindError
	ldapServerFatalError              *LdapServerFatalError
	authenticationError               *AuthenticationError
	privateKeyError                   *PrivateKeyError
	publicKeyError                    *PublicKeyError
	signError                         *SignError
	contentTypeError                  *ContentTypeError
	mysqlFatalError                   *MySQLFatalError
	unssuportedAlgorithmError         *UnsupportedAlgorithmError
	resourceConflictError             *ResourceConflictError
	errRedisFatalError                *RedisFatalError
	errRedisUnmarshalError            *RedisUnmarshalError
	errRedisNotfoundError             *RedisNotFoundError
	errAuthorizationError             *AuthorizationError
	errInvalidAuthorizationTokenError *InvalidAuthorizationTokenError
	errTokenExpiresError              *TokenExpiresError
	notificationConflictError         *NotificationConflictError
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
	suite.resourceConflictError = &ResourceConflictError{err: errors.New("resource conflict error"), requestID: "MockID"}
	suite.errRedisFatalError = &RedisFatalError{err: errors.New("redis fatal error"), requestID: "MockID"}
	suite.errRedisUnmarshalError = &RedisUnmarshalError{err: errors.New("redis Unmarshal error"), requestID: "MockID"}
	suite.errRedisNotfoundError = &RedisNotFoundError{err: errors.New("redis notfound error"), requestID: "MockID"}
	suite.errAuthorizationError = &AuthorizationError{err: errors.New("authorization error"), requestID: "MockID"}
	suite.errInvalidAuthorizationTokenError = &InvalidAuthorizationTokenError{err: errors.New("invalid authorization token"), requestID: "MockID"}
	suite.errTokenExpiresError = &TokenExpiresError{err: errors.New("token expires"), requestID: "MockID"}
	suite.notificationConflictError = &NotificationConflictError{err: errors.New("resource conflict error"), requestID: "MockID"}
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}
