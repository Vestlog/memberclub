package main

import (
	"log"
	"net/http"
)

type CustomResponseWriter struct {
	w http.ResponseWriter
}

func (cw *CustomResponseWriter) Header() http.Header {
	return cw.w.Header()

}

func (cw *CustomResponseWriter) Write(data []byte) (int, error) {
	log.Println(string(data))
	return cw.w.Write(data)
}

func (cw *CustomResponseWriter) WriteHeader(statusCode int) {
	log.Println(statusCode)
	cw.w.WriteHeader(statusCode)
}

func ResponseLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cw := &CustomResponseWriter{w}
		next.ServeHTTP(cw, r)
	})
}

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
