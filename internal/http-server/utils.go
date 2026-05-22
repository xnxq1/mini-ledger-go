package http_server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse[T any] struct {
	Data T `json:"data"`
}

func DecodeJson(body io.Reader, res any) error {
	d := json.NewDecoder(body)
	d.DisallowUnknownFields()
	if err := d.Decode(res); err != nil {
		log.Println(err)
		return CannotDecodeRequest
	}
	return nil
}

func WriteJson(w http.ResponseWriter, res any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

type errCtxKey struct{}

type ErrHolder struct {
	err error
}

func SetError(r *http.Request, err error) {
	if h, ok := r.Context().Value(errCtxKey{}).(*ErrHolder); ok {
		h.err = err
	}
}
