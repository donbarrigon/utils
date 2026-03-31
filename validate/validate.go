package validate

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/donbarrigon/utils/handler"
	"github.com/donbarrigon/utils/herror"
	"github.com/donbarrigon/utils/str"
)

type ValidationFunc func(value reflect.Value, params ...string) (string, str.Placeholder, bool)

type Validation struct {
	Type   string
	Params []string
}

type Rule struct {
	Field       string
	Validations []Validation
}

type Rules []Rule

type Validator interface {
	Rules() Rules
	PrepareForValidation(c *handler.Context) *herror.ValidationError
}

func Body(c *handler.Context, validator Validator) herror.Error {
	if e := c.GetBody(validator); e != nil {
		return e
	}

	e := validator.PrepareForValidation(c)
	// if e == nil {
	// 	e = herror.NewValidationError()
	// }

	rules := validator.Rules()
	return validate(c, validator, rules, e)
}

func Form(c *handler.Context, data any, rules Rules) herror.Error {
	e := herror.NewValidationError()
	return validate(c, data, rules, e)
}

func validate(c *handler.Context, validator any, rules Rules, e *herror.ValidationError) herror.Error {

	val := reflect.ValueOf(validator)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return herror.InternalServerError("Validation failed: expected a pointer to a struct.")
	}

	for _, rule := range rules {
		var v reflect.Value
		var tagName string
		if strings.Contains(rule.Field, ".") {
			fields := strings.Split(rule.Field, ".")
			v = val
			t := val.Type()
			nf := len(fields)

			for i := 0; i < nf; i++ {

				f, ok := t.FieldByName(fields[i])
				if !ok {
					return herror.InternalServerError(fmt.Sprintf("Validation failed: field [%s] not found.", rule.Field))
				}
				tag := strings.Split(f.Tag.Get("json"), ",")[0]
				if tag == "" || tag == "-" {
					tag = fields[i] // fallback al nombre del campo
				}
				if i == 0 {
					tagName = tag
				} else {
					tagName += "." + tag
				}

				v = v.FieldByName(fields[i])
				if v.Kind() == reflect.Pointer {
					if v.IsNil() {
						return herror.InternalServerError(fmt.Sprintf("Validation failed: field [%s] is a nil pointer.", rule.Field))
					}
					v = v.Elem()
				}
				// si no es el ultimo tiene que ser otro struct
				if i < (nf - 1) {
					if !(v.Kind() == reflect.Struct) {
						return herror.InternalServerError(fmt.Sprintf("Validation failed: field [%s] is not a struct.", rule.Field))
					}

					// avanso el tipo al siguiente nivel
					t = f.Type
					if t.Kind() == reflect.Pointer {
						t = t.Elem()
					}
				}
			}
		} else {
			t, ok := val.Type().FieldByName(rule.Field)
			if !ok {
				return herror.InternalServerError(fmt.Sprintf("Validation failed: field [%s] not found.", rule.Field))
			}
			tagName = strings.Split(t.Tag.Get("json"), ",")[0]
			if tagName == "" || tagName == "-" {
				tagName = rule.Field // fallback al nombre del campo
			}

			v = val.FieldByName(rule.Field)
			if v.Kind() == reflect.Pointer {
				if v.IsNil() {
					return herror.InternalServerError(fmt.Sprintf("Validation failed: field [%s] is a nil pointer.", rule.Field))
				}
				v = v.Elem()
			}
		}
		if !v.IsValid() {
			return herror.InternalServerError(fmt.Sprintf("Validation failed: field [%s] not valid.", rule.Field))
		}

		for _, validation := range rule.Validations {
			switch validation.Type {
			case "required":
				if msg, ph, hasError := Required(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "min":
				if msg, ph, hasError := Min(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "max":
				if msg, ph, hasError := Max(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "between":
				if msg, ph, hasError := Between(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "regex":
				if msg, ph, hasError := Regex(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "exists":
				if msg, ph, hasError := Exists(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "not_exists", "notExists":
				if msg, ph, hasError := NotExists(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "unique":
				uparams := append(validation.Params, c.Request.URL.Query().Get("id"))
				if msg, ph, hasError := Unique(v, uparams); hasError {
					e.Append(tagName, msg, ph)
				}
			case "in":
				if msg, ph, hasError := In(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "nin":
				if msg, ph, hasError := Nin(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "before":
				if msg, ph, hasError := Before(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "after":
				if msg, ph, hasError := After(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "before_now", "beforeNow":
				if msg, ph, hasError := BeforeNow(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			case "after_now", "afterNow":
				if msg, ph, hasError := AfterNow(v, validation.Params); hasError {
					e.Append(tagName, msg, ph)
				}
			}
		}
	}
	return e.HasErrors()
}

// type Rules map[string]map[string][]string

// func validate(c *handler.Context, validator any, rules Rules, e *herror.ValidationError) herror.Error {
// 	val := reflect.ValueOf(validator)
// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	if val.Kind() != reflect.Struct {
// 		return herror.InternalServerError("Validation failed: expected a pointer to a struct.")
// 	}
// 	val.FieldByName(rules[0].Field)
// 	typ := val.Type()
// 	numFields := typ.NumField()
// 	for i := range numFields {
// 		field := typ.Field(i)
// 		tag := field.Tag.Get("json")
// 		tagName := strings.Split(tag, ",")[0]
// 		if tagName == "-" || tagName == "id" {
// 			continue
// 		}
// 		if tagName == "" {
// 			tagName = field.Name
// 		}

// 		value := val.Field(i)
// 		if value.Kind() == reflect.Pointer {
// 			if !value.IsNil() {
// 				value = value.Elem()
// 			}
// 		}

// 		if value.Kind() == reflect.Struct {
// 			r := Rules{}
// 			for key, rule := range rules {
// 				if k, found := strings.CutPrefix(key, tagName+"."); found {
// 					r[k] = rule
// 				}
// 			}
// 			e2 := herror.NewValidationError()
// 			if validate(c, value.Interface(), r, e2) != nil {
// 				for e2k, e2v := range e2.Messages {
// 					e.Messages[tagName+"."+e2k] = e2v
// 					e.Placeholders[tagName+"."+e2k] = e2.Placeholders[e2k]
// 				}
// 			}
// 			continue
// 		}

// 		validations := rules[tagName]
// 		if validations == nil {
// 			continue
// 		}
// 		for rule, params := range validations {
// 			switch rule {
// 			case "required":
// 				if msg, ph, hasError := Required(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "min":
// 				if msg, ph, hasError := Min(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "max":
// 				if msg, ph, hasError := Max(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "between":
// 				if msg, ph, hasError := Between(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "regex":
// 				if msg, ph, hasError := Regex(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "exists":
// 				if msg, ph, hasError := Exists(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "not_exists", "notExists":
// 				if msg, ph, hasError := NotExists(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "unique":
// 				params = append(params, c.Request.URL.Query().Get("id"))
// 				if msg, ph, hasError := Unique(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "in":
// 				if msg, ph, hasError := In(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "nin":
// 				if msg, ph, hasError := Nin(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "before":
// 				if msg, ph, hasError := Before(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "after":
// 				if msg, ph, hasError := After(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "before_now", "beforeNow":
// 				if msg, ph, hasError := BeforeNow(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			case "after_now", "afterNow":
// 				if msg, ph, hasError := AfterNow(value, params...); hasError {
// 					e.Append(tagName, msg, ph)
// 				}
// 			}
// 		}
// 	}
// 	return e.HasErrors()
// }

// func (u *UserUpdateProfile) Rules() validation.Rules {
// 	return validation.Rules{
// 		"nickname": {
// 			"between": {"3", "255"},
// 		},
// 		"name": {
// 			"between": {"3", "255"},
// 		},
// 		"phone": {
// 			"regex": {"phone"},
// 		},
// 		"discord": {
// 			"regex":   {"discord"},
// 			"between": {"3", "32"},
// 		},
// 		"cityId": {
// 			"required": {},
// 			"exists":   {"cities", "_id"},
// 		},
// 	}
// }
