package http_server

import (
	"context"
	"net/http"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	holder := ErrHolder{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), errCtxKey{}, &holder))
		next.ServeHTTP(w, r)
		if holder.err != nil {
			apiErr := MappingErrorsToStatus(holder.err)
			WriteJson(w, ErrorResponse{Error: apiErr.Message}, apiErr.Status)
		}
	})
}
