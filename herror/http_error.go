package herror

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/donbarrigon/utils/lang"
)

// HttpError representa un error HTTP con la misma forma que el objeto
// retornado por Nuxt `createError` (status/statusText/message/etc).
type HttpError struct {
	Status     int    `json:"status,omitempty"`     // Código HTTP asociado al error (ej: 404, 500).
	StatusText string `json:"statusText,omitempty"` // Texto corto del estado para el cliente (ej: "Not Found").
	Message    string `json:"message,omitempty"`    // Mensaje legible del error (detalle para logs/UX).
	Name       string `json:"name,omitempty"`       // Nombre/tipo del error (clasificación).
	Cause      any    `json:"cause,omitempty"`      // Causa subyacente (error original) para diagnóstico.
	Data       any    `json:"data,omitempty"`       // Datos adicionales que quieres devolver junto al error.
	Stack      string `json:"stack,omitempty"`      // Stack trace como string para depuración.
	Fatal      bool   `json:"fatal,omitempty"`      // Indica si el error debe tratarse como fatal en el frontend.
}

func New(status int, statusText string, message string, name string, cause any, data any, stack string, fatal bool) Error {
	return &HttpError{
		Status:     status,
		StatusText: statusText,
		Message:    message,
		Name:       name,
		Cause:      cause,
		Data:       data,
		Stack:      stack,
		Fatal:      fatal,
	}
}

func (e *HttpError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Message != "" {
		return e.Message
	}
	if e.StatusText != "" {
		return e.StatusText
	}
	if e.Name != "" {
		return e.Name
	}
	return "Something went wrong"
}

func (e *HttpError) Translate(l string) {
	e.Message = lang.T(l, e.Message, nil)
	e.StatusText = lang.T(l, e.StatusText, nil)
}

var httpErrorBufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func (e *HttpError) WriteJSON(w http.ResponseWriter) Error {
	if w == nil {
		return InternalServerError(errors.New("http.ResponseWriter is nil"))
	}

	buf := httpErrorBufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer httpErrorBufPool.Put(buf)

	if err := json.NewEncoder(buf).Encode(e); err != nil {
		return InternalServerError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	_, _ = w.Write(buf.Bytes())
	return nil
}

func (e *HttpError) WriteProto(w http.ResponseWriter) Error {
	return nil
}
