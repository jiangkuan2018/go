package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range request.Form {
		fmt.Fprintf(writer, "Form[%q] = [%q]\n", k, v)
	}
}
