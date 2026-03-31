package validate

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/donbarrigon/utils/str"
)

// In valida que el valor esté presente en la lista de valores permitidos
func In(value reflect.Value, params []string) (string, str.Placeholder, bool) {
	if len(params) < 1 {
		return "At least one allowed value is required.", str.Placeholder{}, true
	}

	ph := str.Placeholder{{Key: "values", Value: fmt.Sprintf("%v", params)}}

	switch value.Kind() {
	case reflect.String:
		valueStr := value.String()
		for _, param := range params {
			if valueStr == param {
				return "", nil, false
			}
		}
		return "The selected :field is invalid.", ph, true

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valueInt := value.Int()
		for _, param := range params {
			if paramInt, e := strconv.ParseInt(param, 10, 64); e == nil {
				if valueInt == paramInt {
					return "", nil, false
				}
			}
		}
		return "The selected :field is invalid.", ph, true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valueUint := value.Uint()
		for _, param := range params {
			if paramUint, e := strconv.ParseUint(param, 10, 64); e == nil {
				if valueUint == paramUint {
					return "", nil, false
				}
			}
		}
		return "The selected :field is invalid.", ph, true

	case reflect.Float32, reflect.Float64:
		valueFloat := value.Float()
		for _, param := range params {
			if paramFloat, e := strconv.ParseFloat(param, 64); e == nil {
				if valueFloat == paramFloat {
					return "", nil, false
				}
			}
		}
		return "The selected :field is invalid.", ph, true

	case reflect.Slice, reflect.Array:
		// Verificar que todos los elementos del array/slice estén en la lista permitida
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			elementFound := false

			switch element.Kind() {
			case reflect.String:
				elementStr := element.String()
				for _, param := range params {
					if elementStr == param {
						elementFound = true
						break
					}
				}

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				elementInt := element.Int()
				for _, param := range params {
					if paramInt, e := strconv.ParseInt(param, 10, 64); e == nil {
						if elementInt == paramInt {
							elementFound = true
							break
						}
					}
				}

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				elementUint := element.Uint()
				for _, param := range params {
					if paramUint, e := strconv.ParseUint(param, 10, 64); e == nil {
						if elementUint == paramUint {
							elementFound = true
							break
						}
					}
				}

			case reflect.Float32, reflect.Float64:
				elementFloat := element.Float()
				for _, param := range params {
					if paramFloat, e := strconv.ParseFloat(param, 64); e == nil {
						if elementFloat == paramFloat {
							elementFound = true
							break
						}
					}
				}

			default:
				return "The :field contains an item type that is not supported by the in rule.", ph, true
			}

			if !elementFound {
				return "Each value in :field must be one of: :values.", ph, true
			}
		}
		return "", nil, false

	default:
		return "The :field type is not supported by the in rule.", ph, true
	}
}

// Nin valida que el valor NO esté presente en la lista de valores prohibidos
func Nin(value reflect.Value, params []string) (string, str.Placeholder, bool) {
	if len(params) < 1 {
		return "At least one forbidden value is required.", str.Placeholder{}, true
	}

	ph := str.Placeholder{{Key: "values", Value: fmt.Sprintf("%v", params)}}

	switch value.Kind() {
	case reflect.String:
		valueStr := value.String()
		for _, param := range params {
			if valueStr == param {
				return "The :field must not be one of: :values.", ph, true
			}
		}
		return "", nil, false

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valueInt := value.Int()
		for _, param := range params {
			if paramInt, e := strconv.ParseInt(param, 10, 64); e == nil {
				if valueInt == paramInt {
					return "The :field must not be one of: :values.", ph, true
				}
			}
		}
		return "", nil, false

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valueUint := value.Uint()
		for _, param := range params {
			if paramUint, e := strconv.ParseUint(param, 10, 64); e == nil {
				if valueUint == paramUint {
					return "The :field must not be one of: :values.", ph, true
				}
			}
		}
		return "", nil, false

	case reflect.Float32, reflect.Float64:
		valueFloat := value.Float()
		for _, param := range params {
			if paramFloat, e := strconv.ParseFloat(param, 64); e == nil {
				if valueFloat == paramFloat {
					return "The :field must not be one of: :values.", ph, true
				}
			}
		}
		return "", nil, false

	case reflect.Slice, reflect.Array:
		// Verificar que ningún elemento del array/slice esté en la lista prohibida
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)

			switch element.Kind() {
			case reflect.String:
				elementStr := element.String()
				for _, param := range params {
					if elementStr == param {
						return "The :field must not contain any of: :values.", ph, true
					}
				}

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				elementInt := element.Int()
				for _, param := range params {
					if paramInt, e := strconv.ParseInt(param, 10, 64); e == nil {
						if elementInt == paramInt {
							return "The :field must not contain any of: :values.", ph, true
						}
					}
				}

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				elementUint := element.Uint()
				for _, param := range params {
					if paramUint, e := strconv.ParseUint(param, 10, 64); e == nil {
						if elementUint == paramUint {
							return "The :field must not contain any of: :values.", ph, true
						}
					}
				}

			case reflect.Float32, reflect.Float64:
				elementFloat := element.Float()
				for _, param := range params {
					if paramFloat, e := strconv.ParseFloat(param, 64); e == nil {
						if elementFloat == paramFloat {
							return "The :field must not contain any of: :values.", ph, true
						}
					}
				}

			default:
				return "The :field contains an item type that is not supported by the nin rule.", ph, true
			}
		}
		return "", nil, false

	default:
		return "The :field type is not supported by the nin rule.", ph, true
	}
}
