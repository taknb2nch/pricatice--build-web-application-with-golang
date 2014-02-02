package main

import (
	"fmt"
	"net/http"
)

type MuMux struct {
}

func (m *MuMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello muroute!")
}

func main() {
	mux := &MuMux{}

	http.ListenAndServe(":9090", mux)
}
