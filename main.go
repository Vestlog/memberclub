package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("p", "8080", "Port to listen on.")
	resp = flag.Bool("r", false, "Enable response logging. Disabled by default.")
)

func main() {
	flag.Parse()

	c := CreateController()

	mux := http.NewServeMux()
	mux.HandleFunc("/", c.GetAll)
	mux.HandleFunc("/save", c.Save)

	handler := RequestLoggerMiddleware(mux)
	if *resp {
		handler = ResponseLoggerMiddleware(handler)
	}

	log.Println("Listening on", *port)
	log.Println(http.ListenAndServe(":"+*port, handler))
}
