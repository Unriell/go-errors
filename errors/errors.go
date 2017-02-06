package errors

// Below are predefined keys to errors
const (
	ErrBind                     = "BQ0000001"
	ErrResourceNotFound         = "BQ0000002"
	ErrUndefined                = "BQ0000003"
	ErrLDAPServerFatal          = "BQ0000004"
	ErrAuthentication           = "BQ0000005"
	ErrPrivateKey               = "BQ0000006"
	ErrPublicKey                = "BQ0000007"
	ErrSign                     = "BQ0000008"
	ErrUnexpectedContentType    = "BQ0000009"
	ErrMySQLFatal               = "BQ0000010"
	ErrUnsupportedAuthAlgorithm = "BQ0000011"
)

// CustomError wraps any error message
type CustomError interface {
	Error() string
}

// ResourceNotFoundError wraps 404 errors
type ResourceNotFoundError struct {
	Err       error
	Message   string
	RequestID string
}

func (resourceNotFoundError *ResourceNotFoundError) Error() string {
	return ErrResourceNotFound + "|" + resourceNotFoundError.Message + "|" + resourceNotFoundError.RequestID
}

// UndefinedError wraps all unknown errors
type UndefinedError struct {
	Err error
}

func (undefinedError *UndefinedError) Error() string {
	return ErrUndefined + "|" + undefinedError.Err.Error()
}

// BindError wraps JSON struct binding errors
type BindError struct {
	Err       error
	Message   string
	RequestID string
}

func (bindError *BindError) Error() string {
	return ErrBind + "|" + bindError.Err.Error() + ". " + bindError.Message + "|" + bindError.RequestID
}

// LDAPServerFatalError wraps common ldap/ad server errors
type LDAPServerFatalError struct {
	Err       error
	Message   string
	RequestID string
}

func (ldapServerFatalError *LDAPServerFatalError) Error() string {
	return ErrLDAPServerFatal + "|" + ldapServerFatalError.Err.Error() + ". " + ldapServerFatalError.Message + "|" + ldapServerFatalError.RequestID
}

// AuthenticationError wraps all failed auth attempts (both client and server)
type AuthenticationError struct {
	Err       error
	Message   string
	RequestID string
}

func (authenticationError *AuthenticationError) Error() string {
	return ErrAuthentication + "|" + authenticationError.Err.Error() + ". " + authenticationError.Message + "|" + authenticationError.RequestID
}

// PrivateKeyError wraps all errors regarding the private key used for signin
// the JWT tokens
type PrivateKeyError struct {
	Err       error
	Message   string
	RequestID string
}

func (privateKeyError *PrivateKeyError) Error() string {
	return ErrPrivateKey + "|" + privateKeyError.Err.Error() + ". " + privateKeyError.Message + "|" + privateKeyError.RequestID
}

// PublicKeyError wraps all errors regarding the private key used for signin
// the JWT tokens
type PublicKeyError struct {
	Err       error
	Message   string
	RequestID string
}

func (publicKeyError *PublicKeyError) Error() string {
	return ErrPublicKey + "|" + publicKeyError.Err.Error() + ". " + publicKeyError.Message + "|" + publicKeyError.RequestID
}

// SignError wraps all failed token sign attempts
type SignError struct {
	Err       error
	Message   string
	RequestID string
}

func (signError *SignError) Error() string {
	return ErrSign + "|" + signError.Err.Error() + ". " + signError.Message + "|" + signError.RequestID
}

// ContentTypeError wraps all missing or wrong content type headers in requests
type ContentTypeError struct {
	Err       error
	RequestID string
}

func (contentTypeError *ContentTypeError) Error() string {
	return ErrUnexpectedContentType + "|" + contentTypeError.Err.Error() + "|" + contentTypeError.RequestID
}

// MySQLFatalError wraps all mysql common errors
type MySQLFatalError struct {
	Err       error
	RequestID string
}

func (mysqlFatalError *MySQLFatalError) Error() string {
	return ErrMySQLFatal + "|" + mysqlFatalError.Err.Error() + "|" + mysqlFatalError.RequestID
}

// UnsupportedAlgorithmError wraps unknown algorithm requests
type UnsupportedAlgorithmError struct {
	Err       error
	RequestID string
}

func (unssuportedAlgorithmError *UnsupportedAlgorithmError) Error() string {
	return ErrUnsupportedAuthAlgorithm + "|" + unssuportedAlgorithmError.Err.Error() + "|" + unssuportedAlgorithmError.RequestID
}
