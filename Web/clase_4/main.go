package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// se llama mux porque se utiliza gorilamux
	//mux := http.NewServeMux()
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Hola mundo")
	})
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
