package herror

import (
	"net/http"
	"reflect"
	"runtime/debug"
)

func BadRequest(e any) *HttpError {
	return &HttpError{
		Status:     http.StatusBadRequest,
		StatusText: StatusTextBadRequest,
		Message:    StatusMessageBadRequest,
		Name:       "Error",
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
	if e == nil {
		return "<nil>"
	}

	if !config.AppDebug {
		return nil
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
