package validate

import (
	"reflect"
	"strconv"
	"time"

	"github.com/donbarrigon/utils/str"
)

// Assuming fm.fm.Placeholder is defined somewhere like this:
// type fm.Placeholder map[string]string

// Min valida que el valor sea mayor o igual al mínimo especificado
func Min(value reflect.Value, params ...string) (string, str.Placeholder, bool) {
	if len(params) < 1 {
		return "The minimum parameter is required.", str.Placeholder{}, true
	}

	minStr := params[0]
	ph := str.Placeholder{{Key: "min", Value: minStr}}

	switch value.Kind() {
	case reflect.String:
		minLen, err := strconv.Atoi(minStr)
		if err != nil {
			return "The minimum parameter is invalid.", ph, true
		}
		if len(value.String()) < minLen {
			return "The :field must be at least :min characters.", ph, true
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		min, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return "The minimum parameter is invalid.", ph, true
		}
		if value.Int() < min {
			return "The :field must be at least :min.", ph, true
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		min, err := strconv.ParseUint(minStr, 10, 64)
		if err != nil {
			return "The minimum parameter is invalid.", ph, true
		}
		if value.Uint() < min {
			return "The :field must be at least :min.", ph, true
		}

	case reflect.Float32, reflect.Float64:
		min, err := strconv.ParseFloat(minStr, 64)
		if err != nil {
			return "The minimum parameter is invalid.", ph, true
		}
		if value.Float() < min {
			return "The :field must be at least :min.", ph, true
		}

	case reflect.Slice, reflect.Array:
		minLen, err := strconv.Atoi(minStr)
		if err != nil {
			return "The minimum parameter is invalid.", ph, true
		}
		if value.Len() < minLen {
			return "The :field must have at least :min items.", ph, true
		}

	default:
		// Para fechas (time.Time)
		if value.Type() == reflect.TypeFor[time.Time]() {
			min, err := time.Parse("2006-01-02", minStr)
			if err != nil {
				min, err = time.Parse("2006-01-02T15:04:05Z07:00", minStr)
				if err != nil {
					return "The minimum date parameter is not a valid date format.", ph, true
				}
			}
			fieldTime := value.Interface().(time.Time)
			if fieldTime.Before(min) {
				return "The :field must be a date on or after :min.", ph, true
			}
		} else {
			return "The :field type is not supported by the min rule.", ph, true
		}
	}

	return "", nil, false
}

// Max valida que el valor sea menor o igual al máximo especificado
func Max(value reflect.Value, params ...string) (string, str.Placeholder, bool) {
	if len(params) < 1 {
		return "The maximum parameter is required.", str.Placeholder{}, true
	}

	maxStr := params[0]
	ph := str.Placeholder{{Key: "max", Value: maxStr}}

	switch value.Kind() {
	case reflect.String:
		maxLen, err := strconv.Atoi(maxStr)
		if err != nil {
			return "The maximum parameter is invalid.", ph, true
		}
		if len(value.String()) > maxLen {
			return "The :field may not be greater than :max characters.", ph, true
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		max, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return "The maximum parameter is invalid.", ph, true
		}
		if value.Int() > max {
			return "The :field must not be greater than :max.", ph, true
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		max, err := strconv.ParseUint(maxStr, 10, 64)
		if err != nil {
			return "The maximum parameter is invalid.", ph, true
		}
		if value.Uint() > max {
			return "The :field must not be greater than :max.", ph, true
		}

	case reflect.Float32, reflect.Float64:
		max, err := strconv.ParseFloat(maxStr, 64)
		if err != nil {
			return "The maximum parameter is invalid.", ph, true
		}
		if value.Float() > max {
			return "The :field must not be greater than :max.", ph, true
		}

	case reflect.Slice, reflect.Array:
		maxLen, err := strconv.Atoi(maxStr)
		if err != nil {
			return "The maximum parameter is invalid.", ph, true
		}
		if value.Len() > maxLen {
			return "The :field must not have more than :max items.", ph, true
		}

	default:
		// Para fechas (time.Time)
		if value.Type() == reflect.TypeFor[time.Time]() {
			max, err := time.Parse("2006-01-02", maxStr)
			if err != nil {
				max, err = time.Parse("2006-01-02T15:04:05Z07:00", maxStr)
				if err != nil {
					return "The maximum date parameter is not a valid date format.", ph, true
				}
			}
			fieldTime := value.Interface().(time.Time)
			if fieldTime.After(max) {
				return "The :field must be a date on or before :max.", ph, true
			}
		} else {
			return "The :field type is not supported by the max rule.", ph, true
		}
	}

	return "", nil, false
}

// Between valida que el valor esté dentro del rango especificado
func Between(value reflect.Value, params ...string) (string, str.Placeholder, bool) {
	if len(params) < 2 {
		return "Both minimum and maximum parameters are required.", str.Placeholder{}, true
	}

	minStr := params[0]
	maxStr := params[1]
	ph := str.Placeholder{{Key: "min", Value: minStr}, {Key: "max", Value: maxStr}}

	switch value.Kind() {
	case reflect.String:
		min, err1 := strconv.Atoi(minStr)
		max, err2 := strconv.Atoi(maxStr)
		if err1 != nil || err2 != nil {
			return "The length range parameters are invalid.", ph, true
		}
		length := len(value.String())
		if length < min || length > max {
			return "The :field must be between :min and :max characters.", ph, true
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		min, err1 := strconv.ParseInt(minStr, 10, 64)
		max, err2 := strconv.ParseInt(maxStr, 10, 64)
		if err1 != nil || err2 != nil {
			return "The range parameters are invalid.", ph, true
		}
		val := value.Int()
		if val < min || val > max {
			return "The :field must be between :min and :max.", ph, true
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		min, err1 := strconv.ParseUint(minStr, 10, 64)
		max, err2 := strconv.ParseUint(maxStr, 10, 64)
		if err1 != nil || err2 != nil {
			return "The range parameters are invalid.", ph, true
		}
		val := value.Uint()
		if val < min || val > max {
			return "The :field must be between :min and :max.", ph, true
		}

	case reflect.Float32, reflect.Float64:
		min, err1 := strconv.ParseFloat(minStr, 64)
		max, err2 := strconv.ParseFloat(maxStr, 64)
		if err1 != nil || err2 != nil {
			return "The range parameters are invalid.", ph, true
		}
		val := value.Float()
		if val < min || val > max {
			return "The :field must be between :min and :max.", ph, true
		}

	default:
		// Para fechas (time.Time)
		if value.Type() == reflect.TypeFor[time.Time]() {
			min, err1 := time.Parse("2006-01-02", minStr)
			if err1 != nil {
				min, err1 = time.Parse("2006-01-02T15:04:05Z07:00", minStr)
			}
			max, err2 := time.Parse("2006-01-02", maxStr)
			if err2 != nil {
				max, err2 = time.Parse("2006-01-02T15:04:05Z07:00", maxStr)
			}
			if err1 != nil || err2 != nil {
				return "The date range parameters are not valid date formats.", ph, true
			}
			fieldTime := value.Interface().(time.Time)
			if fieldTime.Before(min) || fieldTime.After(max) {
				return "The :field must be between :min and :max.", ph, true
			}
		} else {
			return "The :field type is not supported by the between rule.", ph, true
		}
	}

	return "", nil, false
}
