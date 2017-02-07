package errors

import (
	"strconv"
	"strings"
	"sync"

	"go.bq.com/pdtdev/go-errors.git/dto"

	"github.com/getsentry/raven-go"
)

// ErrorHandler defines an instance for rendering errors
type ErrorHandler struct {
	Values       map[string]dto.ErrorDto
	defaultError string
}

// ErrorHandlerBehavior wraps all errors for easier message access
type ErrorHandlerBehavior interface {
	Error(e error, extraTags map[string]string) dto.ErrorDto
}

var once sync.Once
var ErrorHandlerInstance *ErrorHandler

// NewErrorHandler returns a global ErrorHandler instance
func NewErrorHandler(enviroment, sentryDSN, version string) *ErrorHandler {

	once.Do(func() {

		raven.SetEnvironment(enviroment)
		raven.SetDSN(sentryDSN)
		raven.SetRelease(version)

		ErrorHandlerInstance = &ErrorHandler{
			Values:       values,
			defaultError: "BQ0000000",
		}

	})

	return ErrorHandlerInstance
}

func (eh *ErrorHandler) Error(err error, extraTags map[string]string) dto.ErrorDto {
	rawErrorMsg := err.Error()
	errorDto := eh.Values[eh.defaultError]

	if len(rawErrorMsg) == 0 {
		errorDto.ComponentMsg = "Undefined error code. No error code found"
	} else {
		msg := strings.Split(rawErrorMsg, "|")
		if len(msg) == 0 {
			errorDto.ComponentMsg = "Undefined error code. No error code found"
		} else {
			var id, ok = msg[0], false
			if errorDto, ok = eh.Values[id]; ok {
				errorDto.ComponentMsg = msg[1]
				errorDto.SentryCode = msg[2]
			} else {
				errorDto = eh.Values[eh.defaultError]
				errorDto.ComponentMsg = "this error doesn't exist. Msg " + rawErrorMsg
			}
		}
	}

	tags := eh.getSentryErrorsTags(extraTags)
	tags["id"] = errorDto.ID
	tags["status"] = strconv.Itoa(errorDto.Status)
	tags["sentryCode"] = errorDto.SentryCode

	go func() { raven.CaptureErrorAndWait(err, tags) }()
	return errorDto
}

func (eh *ErrorHandler) getSentryErrorsTags(extraTags map[string]string) map[string]string {
	dst := make(map[string]string, len(extraTags))
	for k, v := range extraTags {
		dst[k] = v
	}

	return dst
}
