package rutas

import (
	"fmt"
	"net/http"
)

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola mundo desde golang")
}
