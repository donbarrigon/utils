package herror

import (
	"net/http"
	"reflect"
	"runtime/debug"

	"github.com/donbarrigon/utils/config"
)

const errorName = "Error"

func BadRequest(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusBadRequest,
		StatusText: StatusTextBadRequest,
		Message:    StatusMessageBadRequest,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func BadRequestMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusBadRequest,
		StatusText: StatusTextBadRequest,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func Unauthorized(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusUnauthorized,
		StatusText: StatusTextUnauthorized,
		Message:    StatusMessageUnauthorized,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UnauthorizedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusUnauthorized,
		StatusText: StatusTextUnauthorized,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func PaymentRequired(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusPaymentRequired,
		StatusText: StatusTextPaymentRequired,
		Message:    StatusMessagePaymentRequired,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func PaymentRequiredMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusPaymentRequired,
		StatusText: StatusTextPaymentRequired,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func Forbidden(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusForbidden,
		StatusText: StatusTextForbidden,
		Message:    StatusMessageForbidden,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ForbiddenMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusForbidden,
		StatusText: StatusTextForbidden,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotFound(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusNotFound,
		StatusText: StatusTextNotFound,
		Message:    StatusMessageNotFound,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotFoundMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusNotFound,
		StatusText: StatusTextNotFound,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func MethodNotAllowed(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusMethodNotAllowed,
		StatusText: StatusTextMethodNotAllowed,
		Message:    StatusMessageMethodNotAllowed,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func MethodNotAllowedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusMethodNotAllowed,
		StatusText: StatusTextMethodNotAllowed,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotAcceptable(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusNotAcceptable,
		StatusText: StatusTextNotAcceptable,
		Message:    StatusMessageNotAcceptable,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotAcceptableMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusNotAcceptable,
		StatusText: StatusTextNotAcceptable,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ProxyAuthRequired(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusProxyAuthRequired,
		StatusText: StatusTextProxyAuthRequired,
		Message:    StatusMessageProxyAuthRequired,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ProxyAuthRequiredMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusProxyAuthRequired,
		StatusText: StatusTextProxyAuthRequired,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestTimeout(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestTimeout,
		StatusText: StatusTextRequestTimeout,
		Message:    StatusMessageRequestTimeout,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestTimeoutMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestTimeout,
		StatusText: StatusTextRequestTimeout,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func Conflict(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusConflict,
		StatusText: StatusTextConflict,
		Message:    StatusMessageConflict,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ConflictMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusConflict,
		StatusText: StatusTextConflict,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func Gone(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusGone,
		StatusText: StatusTextGone,
		Message:    StatusMessageGone,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func GoneMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusGone,
		StatusText: StatusTextGone,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func LengthRequired(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusLengthRequired,
		StatusText: StatusTextLengthRequired,
		Message:    StatusMessageLengthRequired,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func LengthRequiredMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusLengthRequired,
		StatusText: StatusTextLengthRequired,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func PreconditionFailed(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusPreconditionFailed,
		StatusText: StatusTextPreconditionFailed,
		Message:    StatusMessagePreconditionFailed,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func PreconditionFailedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusPreconditionFailed,
		StatusText: StatusTextPreconditionFailed,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestEntityTooLarge(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestEntityTooLarge,
		StatusText: StatusTextRequestEntityTooLarge,
		Message:    StatusMessageRequestEntityTooLarge,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestEntityTooLargeMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestEntityTooLarge,
		StatusText: StatusTextRequestEntityTooLarge,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestURITooLong(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestURITooLong,
		StatusText: StatusTextRequestURITooLong,
		Message:    StatusMessageRequestURITooLong,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestURITooLongMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestURITooLong,
		StatusText: StatusTextRequestURITooLong,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UnsupportedMediaType(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusUnsupportedMediaType,
		StatusText: StatusTextUnsupportedMediaType,
		Message:    StatusMessageUnsupportedMediaType,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UnsupportedMediaTypeMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusUnsupportedMediaType,
		StatusText: StatusTextUnsupportedMediaType,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestedRangeNotSatisfiable(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestedRangeNotSatisfiable,
		StatusText: StatusTextRequestedRangeNotSatisfiable,
		Message:    StatusMessageRequestedRangeNotSatisfiable,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestedRangeNotSatisfiableMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestedRangeNotSatisfiable,
		StatusText: StatusTextRequestedRangeNotSatisfiable,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ExpectationFailed(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusExpectationFailed,
		StatusText: StatusTextExpectationFailed,
		Message:    StatusMessageExpectationFailed,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ExpectationFailedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusExpectationFailed,
		StatusText: StatusTextExpectationFailed,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func Teapot(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusTeapot,
		StatusText: StatusTextTeapot,
		Message:    StatusMessageTeapot,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func TeapotMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusTeapot,
		StatusText: StatusTextTeapot,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func MisdirectedRequest(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusMisdirectedRequest,
		StatusText: StatusTextMisdirectedRequest,
		Message:    StatusMessageMisdirectedRequest,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func MisdirectedRequestMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusMisdirectedRequest,
		StatusText: StatusTextMisdirectedRequest,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UnprocessableEntity(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusUnprocessableEntity,
		StatusText: StatusTextUnprocessableEntity,
		Message:    StatusMessageUnprocessableEntity,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UnprocessableEntityMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusUnprocessableEntity,
		StatusText: StatusTextUnprocessableEntity,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func Locked(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusLocked,
		StatusText: StatusTextLocked,
		Message:    StatusMessageLocked,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func LockedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusLocked,
		StatusText: StatusTextLocked,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func FailedDependency(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusFailedDependency,
		StatusText: StatusTextFailedDependency,
		Message:    StatusMessageFailedDependency,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func FailedDependencyMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusFailedDependency,
		StatusText: StatusTextFailedDependency,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func TooEarly(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusTooEarly,
		StatusText: StatusTextTooEarly,
		Message:    StatusMessageTooEarly,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func TooEarlyMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusTooEarly,
		StatusText: StatusTextTooEarly,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UpgradeRequired(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusUpgradeRequired,
		StatusText: StatusTextUpgradeRequired,
		Message:    StatusMessageUpgradeRequired,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UpgradeRequiredMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusUpgradeRequired,
		StatusText: StatusTextUpgradeRequired,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func PreconditionRequired(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusPreconditionRequired,
		StatusText: StatusTextPreconditionRequired,
		Message:    StatusMessagePreconditionRequired,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func PreconditionRequiredMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusPreconditionRequired,
		StatusText: StatusTextPreconditionRequired,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func TooManyRequests(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusTooManyRequests,
		StatusText: StatusTextTooManyRequests,
		Message:    StatusMessageTooManyRequests,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func TooManyRequestsMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusTooManyRequests,
		StatusText: StatusTextTooManyRequests,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestHeaderFieldsTooLarge(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestHeaderFieldsTooLarge,
		StatusText: StatusTextRequestHeaderFieldsTooLarge,
		Message:    StatusMessageRequestHeaderFieldsTooLarge,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func RequestHeaderFieldsTooLargeMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusRequestHeaderFieldsTooLarge,
		StatusText: StatusTextRequestHeaderFieldsTooLarge,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UnavailableForLegalReasons(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusUnavailableForLegalReasons,
		StatusText: StatusTextUnavailableForLegalReasons,
		Message:    StatusMessageUnavailableForLegalReasons,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func UnavailableForLegalReasonsMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusUnavailableForLegalReasons,
		StatusText: StatusTextUnavailableForLegalReasons,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func InternalServerError(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusInternalServerError,
		StatusText: StatusTextInternalServerError,
		Message:    StatusMessageInternalServerError,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func InternalServerErrorMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusInternalServerError,
		StatusText: StatusTextInternalServerError,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotImplemented(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusNotImplemented,
		StatusText: StatusTextNotImplemented,
		Message:    StatusMessageNotImplemented,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotImplementedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusNotImplemented,
		StatusText: StatusTextNotImplemented,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func BadGateway(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusBadGateway,
		StatusText: StatusTextBadGateway,
		Message:    StatusMessageBadGateway,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func BadGatewayMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusBadGateway,
		StatusText: StatusTextBadGateway,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ServiceUnavailable(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusServiceUnavailable,
		StatusText: StatusTextServiceUnavailable,
		Message:    StatusMessageServiceUnavailable,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func ServiceUnavailableMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusServiceUnavailable,
		StatusText: StatusTextServiceUnavailable,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func GatewayTimeout(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusGatewayTimeout,
		StatusText: StatusTextGatewayTimeout,
		Message:    StatusMessageGatewayTimeout,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func GatewayTimeoutMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusGatewayTimeout,
		StatusText: StatusTextGatewayTimeout,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func HTTPVersionNotSupported(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusHTTPVersionNotSupported,
		StatusText: StatusTextHTTPVersionNotSupported,
		Message:    StatusMessageHTTPVersionNotSupported,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func HTTPVersionNotSupportedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusHTTPVersionNotSupported,
		StatusText: StatusTextHTTPVersionNotSupported,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func VariantAlsoNegotiates(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusVariantAlsoNegotiates,
		StatusText: StatusTextVariantAlsoNegotiates,
		Message:    StatusMessageVariantAlsoNegotiates,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func VariantAlsoNegotiatesMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusVariantAlsoNegotiates,
		StatusText: StatusTextVariantAlsoNegotiates,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func InsufficientStorage(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusInsufficientStorage,
		StatusText: StatusTextInsufficientStorage,
		Message:    StatusMessageInsufficientStorage,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func InsufficientStorageMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusInsufficientStorage,
		StatusText: StatusTextInsufficientStorage,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func LoopDetected(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusLoopDetected,
		StatusText: StatusTextLoopDetected,
		Message:    StatusMessageLoopDetected,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func LoopDetectedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusLoopDetected,
		StatusText: StatusTextLoopDetected,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotExtended(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusNotExtended,
		StatusText: StatusTextNotExtended,
		Message:    StatusMessageNotExtended,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NotExtendedMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusNotExtended,
		StatusText: StatusTextNotExtended,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NetworkAuthenticationRequired(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusNetworkAuthenticationRequired,
		StatusText: StatusTextNetworkAuthenticationRequired,
		Message:    StatusMessageNetworkAuthenticationRequired,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func NetworkAuthenticationRequiredMsg(e any, message string) *HttpError {
	return &HttpError{
		Status:     http.StatusNetworkAuthenticationRequired,
		StatusText: StatusTextNetworkAuthenticationRequired,
		Message:    message,
		Name:       errorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

func stackData() string {
	if !config.AppDebug {
		return ""
	}
	return string(debug.Stack())
}

func causeData(e any) any {

	if !config.AppDebug {
		return nil
	}

	if e == nil {
		return "<nil>"
	}

	v := reflect.ValueOf(e)
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		result := make(map[string]any)
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)

			if field.IsExported() {
				fieldValue := v.Field(i)
				result[field.Name] = fieldValue.Interface()
			}
		}

		if len(result) > 0 {
			return result
		}
	}

	if er, ok := e.(error); ok {
		return er.Error()
	}

	return e
}
