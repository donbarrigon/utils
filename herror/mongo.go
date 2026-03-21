package herror

import (
	"context"
	"errors"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/donbarrigon/utils/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const mongoErrorName = "db-error"
const hexErrorName = "id-error"

func HexID(e any) *HttpError {

	return &HttpError{
		Status:     http.StatusBadRequest,
		StatusText: StatusTextBadRequest,
		Message:    "The provided identifier is not a valid ObjectID.",
		Name:       hexErrorName,
		Cause:      causeData(e),
		Data:       nil,
		Stack:      stackData(),
		Fatal:      false,
	}
}

// Mongo converts MongoDB driver errors (v2) into standardized HTTP errors.
func Mongo(e error) *HttpError {
	if e == nil {
		return nil
	}

	var (
		stack string
		cause any
	)

	if config.AppDebug {
		stack = string(debug.Stack())
		cause = e.Error()
	}

	// No document found
	if errors.Is(e, mongo.ErrNoDocuments) {
		return &HttpError{
			Status:     http.StatusNotFound,
			StatusText: StatusTextNotFound,
			Message:    "No matching record was found.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// Client disconnected
	if errors.Is(e, mongo.ErrClientDisconnected) {
		return &HttpError{
			Status:     http.StatusServiceUnavailable,
			StatusText: StatusTextServiceUnavailable,
			Message:    "The database connection was interrupted. Please try again.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// Deadline exceeded (timeout)
	if errors.Is(e, context.DeadlineExceeded) {
		return &HttpError{
			Status:     http.StatusRequestTimeout,
			StatusText: StatusTextRequestTimeout,
			Message:    "The database operation timed out. Please try again.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// Context canceled (client aborted request)
	if errors.Is(e, context.Canceled) {
		return &HttpError{
			Status:     http.StatusBadRequest,
			StatusText: StatusTextBadRequest,
			Message:    "The request was canceled before the database operation completed.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// Duplicate key errors
	if mongo.IsDuplicateKeyError(e) {
		return &HttpError{
			Status:     http.StatusConflict,
			StatusText: StatusTextConflict,
			Message:    "A record with the same unique value already exists.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// WriteException (most write-related errors)
	var writeException mongo.WriteException
	if errors.As(e, &writeException) {
		for _, writeError := range writeException.WriteErrors {
			switch writeError.Code {
			case 11000, 11001: // duplicate key
				return &HttpError{
					Status:     http.StatusConflict,
					StatusText: StatusTextConflict,
					Message:    "A record with the same unique value already exists.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			case 2: // BadValue
				return &HttpError{
					Status:     http.StatusBadRequest,
					StatusText: StatusTextBadRequest,
					Message:    "One of the provided values is invalid.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			case 9: // FailedToParse
				return &HttpError{
					Status:     http.StatusBadRequest,
					StatusText: StatusTextBadRequest,
					Message:    "The database could not process the provided data.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			case 14: // TypeMismatch
				return &HttpError{
					Status:     http.StatusBadRequest,
					StatusText: StatusTextBadRequest,
					Message:    "A field has an unexpected data type.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			case 16755: // Location error
				return &HttpError{
					Status:     http.StatusBadRequest,
					StatusText: StatusTextBadRequest,
					Message:    "The geographic location value is not valid.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			case 17280: // KeyTooLong
				return &HttpError{
					Status:     http.StatusBadRequest,
					StatusText: StatusTextBadRequest,
					Message:    "A provided key is too long.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			case 10334: // BSONObjectTooLarge
				return &HttpError{
					Status:     http.StatusBadRequest,
					StatusText: StatusTextBadRequest,
					Message:    "The document is too large to store.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			}
		}

		if writeException.WriteConcernError != nil {
			switch writeException.WriteConcernError.Code {
			case 64: // WriteConcernFailed
				return &HttpError{
					Status:     http.StatusServiceUnavailable,
					StatusText: StatusTextServiceUnavailable,
					Message:    "The database could not confirm the write. Please try again.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			case 79: // UnknownReplWriteConcern
				return &HttpError{
					Status:     http.StatusBadRequest,
					StatusText: StatusTextBadRequest,
					Message:    "The write configuration is not valid.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			}
		}

		return &HttpError{
			Status:     http.StatusInternalServerError,
			StatusText: StatusTextInternalServerError,
			Message:    "The database write operation failed.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// BulkWriteException
	var bulkWriteException mongo.BulkWriteException
	if errors.As(e, &bulkWriteException) {
		for _, writeError := range bulkWriteException.WriteErrors {
			if writeError.Code == 11000 || writeError.Code == 11001 {
				return &HttpError{
					Status:     http.StatusConflict,
					StatusText: StatusTextConflict,
					Message:    "Some records already exist with the same unique value.",
					Name:       mongoErrorName,
					Cause:      cause,
					Data:       nil,
					Stack:      stack,
					Fatal:      false,
				}
			}
		}

		if bulkWriteException.WriteConcernError != nil {
			return &HttpError{
				Status:     http.StatusServiceUnavailable,
				StatusText: StatusTextServiceUnavailable,
				Message:    "The database could not confirm all updates. Please try again.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		}

		return &HttpError{
			Status:     http.StatusBadRequest,
			StatusText: StatusTextBadRequest,
			Message:    "Some records could not be processed.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// CommandError
	var commandError mongo.CommandError
	if errors.As(e, &commandError) {
		switch commandError.Code {
		case 2: // BadValue
			return &HttpError{
				Status:     http.StatusBadRequest,
				StatusText: StatusTextBadRequest,
				Message:    "One of the request parameters is invalid.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 9: // FailedToParse
			return &HttpError{
				Status:     http.StatusBadRequest,
				StatusText: StatusTextBadRequest,
				Message:    "The database could not process the request.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 13: // Unauthorized
			return &HttpError{
				Status:     http.StatusForbidden,
				StatusText: StatusTextForbidden,
				Message:    "You do not have permission to perform this action.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 18: // AuthenticationFailed
			return &HttpError{
				Status:     http.StatusUnauthorized,
				StatusText: StatusTextUnauthorized,
				Message:    "Authentication with the database failed.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 26: // NamespaceNotFound
			return &HttpError{
				Status:     http.StatusNotFound,
				StatusText: StatusTextNotFound,
				Message:    "The requested collection does not exist.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 59: // CommandNotFound
			return &HttpError{
				Status:     http.StatusBadRequest,
				StatusText: StatusTextBadRequest,
				Message:    "The requested operation is not valid.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 61: // ShardKeyNotFound
			return &HttpError{
				Status:     http.StatusBadRequest,
				StatusText: StatusTextBadRequest,
				Message:    "A required shard key is missing.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 72: // InvalidOptions
			return &HttpError{
				Status:     http.StatusBadRequest,
				StatusText: StatusTextBadRequest,
				Message:    "The provided options are not valid.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 96: // OperationFailed
			return &HttpError{
				Status:     http.StatusInternalServerError,
				StatusText: StatusTextInternalServerError,
				Message:    "The database operation failed.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 11600: // InterruptedAtShutdown
			return &HttpError{
				Status:     http.StatusServiceUnavailable,
				StatusText: StatusTextServiceUnavailable,
				Message:    "The database server is shutting down. Please try again.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 11601: // Interrupted
			return &HttpError{
				Status:     http.StatusServiceUnavailable,
				StatusText: StatusTextServiceUnavailable,
				Message:    "The database operation was interrupted. Please try again.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 13435: // ShardKeyTooBig
			return &HttpError{
				Status:     http.StatusBadRequest,
				StatusText: StatusTextBadRequest,
				Message:    "The provided shard key is too large.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 16550: // DocumentValidationFailure
			return &HttpError{
				Status:     http.StatusBadRequest,
				StatusText: StatusTextBadRequest,
				Message:    "The provided document does not meet validation rules.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		case 50: // MaxTimeMSExpired
			return &HttpError{
				Status:     http.StatusRequestTimeout,
				StatusText: StatusTextRequestTimeout,
				Message:    "The database operation timed out.",
				Name:       mongoErrorName,
				Cause:      cause,
				Data:       nil,
				Stack:      stack,
				Fatal:      false,
			}
		}

		return &HttpError{
			Status:     http.StatusInternalServerError,
			StatusText: StatusTextInternalServerError,
			Message:    "The database command failed.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// ServerError fallback
	var serverError mongo.ServerError
	if errors.As(e, &serverError) {
		_ = serverError // reserved for future extra context
		return &HttpError{
			Status:     http.StatusServiceUnavailable,
			StatusText: StatusTextServiceUnavailable,
			Message:    "There is a problem with the database server. Please try again later.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// Network/connection errors fallback by message
	errorMsg := strings.ToLower(e.Error())
	switch {
	case strings.Contains(errorMsg, "connection refused"),
		strings.Contains(errorMsg, "no reachable servers"),
		strings.Contains(errorMsg, "server selection timeout"):
		return &HttpError{
			Status:     http.StatusServiceUnavailable,
			StatusText: StatusTextServiceUnavailable,
			Message:    "We could not connect to the database. Please try again.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	case strings.Contains(errorMsg, "authentication failed"):
		return &HttpError{
			Status:     http.StatusUnauthorized,
			StatusText: StatusTextUnauthorized,
			Message:    "Authentication with the database failed.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	case strings.Contains(errorMsg, "not authorized"):
		return &HttpError{
			Status:     http.StatusForbidden,
			StatusText: StatusTextForbidden,
			Message:    "You do not have permission to access the requested database resource.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	case strings.Contains(errorMsg, "invalid namespace"):
		return &HttpError{
			Status:     http.StatusBadRequest,
			StatusText: StatusTextBadRequest,
			Message:    "The requested collection name is not valid.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	case strings.Contains(errorMsg, "exceeds maximum"):
		return &HttpError{
			Status:     http.StatusBadRequest,
			StatusText: StatusTextBadRequest,
			Message:    "The provided data exceeds the maximum allowed size.",
			Name:       mongoErrorName,
			Cause:      cause,
			Data:       nil,
			Stack:      stack,
			Fatal:      false,
		}
	}

	// Default: Internal Server Error
	return &HttpError{
		Status:     http.StatusInternalServerError,
		StatusText: StatusTextInternalServerError,
		Message:    "The server encountered an unexpected database error.",
		Name:       mongoErrorName,
		Cause:      cause,
		Data:       nil,
		Stack:      stack,
		Fatal:      false,
	}
}
