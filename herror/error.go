package herror

import (
	"io"
)

type Error interface {
	Error() string
	Translate(lang string)
	WriteJSON(w io.Writer) Error
	WriteProto(w io.Writer) Error
}
