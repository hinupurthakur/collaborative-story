package middlewares

import "net/http"

func EnableCORS(next http.Handler) http.Handler {
	// TODO: logic for enabling CORS, if time permits
	return next
}
