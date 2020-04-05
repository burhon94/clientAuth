package logger

import (
	"log"
	"net/http"
)

func Logger(loggerText string) func(
	next http.HandlerFunc,
	) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			log.Printf("Method: %s, Path: %s, Logger: %s",
				request.Method,
				request.URL.Path,
				loggerText,
			)
			next(writer, request)
		}
	}
}