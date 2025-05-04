package main

import (
	"clase4_web/rutas"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// func main() {
// 	// se llama mux porque se utiliza gorilamux
// 	//mux := http.NewServeMux()
// 	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
// 		fmt.Fprintln(response, "Hola mundo")
// 	})
// 	log.Fatal(http.ListenAndServe("localhost:8081", nil))
// }

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", rutas.Home)

	errorVariables := godotenv.Load()

	if errorVariables != nil {
		panic(errorVariables)
		return
	}

	server := &http.Server{
		Addr:         "localhost:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Corriendo servidor desde: ")
	log.Fatal(server.ListenAndServe())

}
