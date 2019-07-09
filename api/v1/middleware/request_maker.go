package middleware

import (
	"grpc_unittest/configs"
	"net/http"
)

type RequestMiddleware struct {
}

//Generate Request ID
func (rm RequestMiddleware) RequestIdGenerator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(configs.WithRequestID(r.Context()))
		next.ServeHTTP(w, r)
	})
}
