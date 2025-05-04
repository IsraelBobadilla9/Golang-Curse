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

// EJECUTAR go run github.com/pilu/fresh
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
	// rutas
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", rutas.Parametros)
	mux.HandleFunc("/parametros-querystring", rutas.ParametrosQueryString)
	// ejecucion de servidor
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
