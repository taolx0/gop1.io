package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//goland:noinspection ALL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "header[%q]=[%q]\n", k, v)
	}
	fmt.Fprintf(w, "host=%s\n", r.Host)
	fmt.Fprintf(w, "remoteAddr=%s\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, val := range r.Form {
		fmt.Fprintf(w, "Form[%q]=[%q]\n", k, val)
	}
}
