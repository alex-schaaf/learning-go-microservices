package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// greedy matching, anything /... that's not registered to another HandeFunc
	// will execute
	http.HandleFunc("/", handleRoot)
	// bind server to every ip address at given port
	http.ListenAndServe(":9090", nil)
}

func handleRoot(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello World!")
	// read in http request body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello, %s", d)
}
