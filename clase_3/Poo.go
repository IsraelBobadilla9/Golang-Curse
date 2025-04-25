package main // modulo

import (
	"fmt"
	"reflect"
)

// Estructuras

// es un objeto de datos con diferentes campos
// igual los nombres de los campos tienen que ser publicos por eso se colocan asi
type Persona struct {
	Id     int
	Nombre string
	Correo string
	Edad   int
}

func main() {

	fmt.Println("Inicio del main")
	// Generamos una instancia de nuestra estructura
	estructura := Persona{
		Id:     1,
		Nombre: "Israel",
		Correo: "is@hotmail.com",
		Edad:   42,
	}
	fmt.Println(estructura)
	fmt.Println(estructura.Id)
	fmt.Println(reflect.TypeOf(estructura))
	fmt.Printf("%+v \n", estructura)
	fmt.Println("================================")
	// Existe otra forma de instanciar una estructura
	p := new(Persona)
	p.Nombre = "Israel"
	fmt.Printf("%+v \n", p)
	fmt.Println(reflect.TypeOf(p))

}
