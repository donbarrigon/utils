package herror

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/donbarrigon/utils/lang"
	"github.com/donbarrigon/utils/str"
)

type ValidationError struct {
	Messages     map[string][]string
	Placeholders map[string][]str.Placeholder
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		Messages:     map[string][]string{},
		Placeholders: map[string][]str.Placeholder{},
	}
}

func (e *ValidationError) Append(field string, message string, ph str.Placeholder) {
	ph.Append("field", field)
	e.Messages[field] = append(e.Messages[field], message)
	e.Placeholders[field] = append(e.Placeholders[field], ph)
}

func (e *ValidationError) AppendM(field string, message string) {
	e.Messages[field] = append(e.Messages[field], message)
	e.Placeholders[field] = append(e.Placeholders[field], str.Placeholder{{Key: "field", Value: field}})
}

func (e *ValidationError) HasErrors() error {
	if len(e.Messages) > 0 {
		return e
	}
	return nil
}

// ================================================================
// Funciones para la interfaz de Errror
// ================================================================

func (ve *ValidationError) Error() string {
	b, err := json.Marshal(ve.Messages)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func (ve *ValidationError) Translate(l string) {
	for key, messages := range ve.Messages {
		for i, message := range messages {
			messages[i] = lang.T(l, message, ve.Placeholders[key][i])
		}
	}
}

func (ve *ValidationError) WriteJSON(w http.ResponseWriter) Error {
	if w == nil {
		return InternalServerError(errors.New("http.ResponseWriter is nil"))
	}

	data := ve.Messages
	if data == nil {
		data = map[string][]string{}
	}

	dataBytes, e := json.Marshal(data)
	if e != nil {
		return InternalServerError(e)
	}

	var buf bytes.Buffer
	buf.WriteString(`{"status":`)
	buf.WriteString(strconv.Itoa(http.StatusUnprocessableEntity))
	buf.WriteString(`,"statusText":`)
	buf.WriteString(strconv.Quote(StatusTextUnprocessableEntity))
	buf.WriteString(`,"message":`)
	buf.WriteString(strconv.Quote(StatusMessageUnprocessableEntity))
	buf.WriteString(`,"name":"validation-error","data":`)
	buf.Write(dataBytes)
	buf.WriteByte('}')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, _ = w.Write(buf.Bytes())
	return nil
}

func (ve *ValidationError) WriteProto(w http.ResponseWriter) Error {
	return nil
}
