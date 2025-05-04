package rutas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola mundo desde golang")
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "sobre nosotros nuevo")
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ID = "+vars["id"]+" slug = "+vars["slug"])
	fmt.Fprintln(response, "FINAL DE DATOS")
}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Fprintln(response, request.URL)
	fmt.Fprintln(response, request.URL.RawQuery)
	fmt.Fprintln(response, request.URL.Query())
	fmt.Fprintln(response, request.URL.Query().Get("id"))
	id := request.URL.Query().Get("id")
	fmt.Fprintln(response, "id = %s", id)
}
