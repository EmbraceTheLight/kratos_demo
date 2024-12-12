package encoder

import (
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdhttp "net/http"
)

type Errors struct {
	Field []string `json:"body"`
}

type HTTPError struct {
	code   int     `json:"-"`
	Errors *Errors `json:"errors"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError: %d,detail: %v", e.code, e.Errors.Field)
}

func NewHTTPError(code int, field string) *HTTPError {
	return &HTTPError{
		code: code,
		Errors: &Errors{
			Field: []string{field},
		},
	}
}
func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	return &HTTPError{
		Errors: &Errors{
			Field: []string{err.Error()},
		},
	}
}

func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	//fmt.Println(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.code)
	_, _ = w.Write(body)
}
