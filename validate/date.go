package validate

import (
	"reflect"
	"time"

	"github.com/donbarrigon/utils/str"
)

// Before valida que la fecha sea anterior a la fecha especificada
func Before(value reflect.Value, params []string) (string, str.Placeholder, bool) {
	if len(params) < 1 {
		return "A date parameter is required.", str.Placeholder{}, true
	}

	beforeStr := params[0]
	ph := str.Placeholder{{Key: "date", Value: beforeStr}}

	if value.Type() != reflect.TypeFor[time.Time]() {
		return "The :field field must be a valid date.", ph, true
	}

	before, err := time.Parse("2006-01-02", beforeStr)
	if err != nil {
		before, err = time.Parse("2006-01-02T15:04:05Z07:00", beforeStr)
		if err != nil {
			return "The date does not match a valid format (expected YYYY-MM-DD or ISO 8601).", ph, true
		}
	}

	fieldTime := value.Interface().(time.Time)
	if !fieldTime.Before(before) {
		return "The :field must be a date before :date.", ph, true
	}

	return "", nil, false
}

// After valida que la fecha sea posterior a la fecha especificada
func After(value reflect.Value, params []string) (string, str.Placeholder, bool) {
	if len(params) < 1 {
		return "A date parameter is required.", str.Placeholder{}, true
	}

	afterStr := params[0]
	ph := str.Placeholder{{Key: "date", Value: afterStr}}

	if value.Type() != reflect.TypeFor[time.Time]() {
		return "The :field field must be a valid date.", ph, true
	}

	after, err := time.Parse("2006-01-02", afterStr)
	if err != nil {
		after, err = time.Parse("2006-01-02T15:04:05Z07:00", afterStr)
		if err != nil {
			return "The date does not match a valid format (expected YYYY-MM-DD or ISO 8601).", ph, true
		}
	}

	fieldTime := value.Interface().(time.Time)
	if !fieldTime.After(after) {
		return "The :field must be a date after :date.", ph, true
	}

	return "", nil, false
}

// BeforeNow valida que la fecha sea anterior a la fecha actual
func BeforeNow(value reflect.Value, params []string) (string, str.Placeholder, bool) {
	ph := str.Placeholder{{Key: "date", Value: "now"}}

	if value.Type() != reflect.TypeFor[time.Time]() {
		return "The :field field must be a valid date.", ph, true
	}

	fieldTime := value.Interface().(time.Time)
	now := time.Now()

	if !fieldTime.Before(now) {
		return "The :field must be a date before :date.", ph, true
	}

	return "", nil, false
}

// AfterNow valida que la fecha sea posterior a la fecha actual
func AfterNow(value reflect.Value, params []string) (string, str.Placeholder, bool) {
	ph := str.Placeholder{{Key: "date", Value: "now"}}

	if value.Type() != reflect.TypeFor[time.Time]() {
		return "The :field field must be a valid date.", ph, true
	}

	fieldTime := value.Interface().(time.Time)
	now := time.Now()

	if !fieldTime.After(now) {
		return "The :field must be a date after :date.", ph, true
	}

	return "", nil, false
}
