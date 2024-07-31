package middleware

import "net/http"

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middlware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	middlware.Handler.ServeHTTP(writer, request)

}
