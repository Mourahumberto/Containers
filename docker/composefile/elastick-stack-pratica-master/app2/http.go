package main

import (
	"fmt"
	"log"
	"net/http"

	"go.elastic.co/apm/module/apmhttp"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world!\n")
}

func helloHandler2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world 2!\n")
}
func main() {

	r := http.NewServeMux()
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/teste", helloHandler2)
	log.Fatal(http.ListenAndServe(":5000", apmhttp.Wrap(r)))

}
