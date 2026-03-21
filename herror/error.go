package herror

import "net/http"

type Error interface {
	Error() string
	Translate(lang string)
	WriteJSON(w http.ResponseWriter) Error
	WriteProto(w http.ResponseWriter) Error
}
