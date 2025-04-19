package main // Este es el nombre del modulo, se le dice que es un handler, ya que nosotros modularizamos las funciones y las dejamos para que la aplicacion las ejecute bajo demanda
// fmt permite generar salidas de texto
import (
	"fmt"
)

// Como ejecutar?
// 1) go mod init clase_1
// 2 ) go run .
// En caso de no generar el archivo mod: go run . main.go

// Todo tiene que estar dentro de una funcion, la funcion main es especial

// Variables y constantes
// Si se crea una constante y la primer letra es mayuscula es de tipo publico y se utiliza en cualquier modulo
/*
const MICONSTANTE = "Bobadilla constante Ñandúes" // No es necesario poner el tipo de dato

func main() { // las llaves siempre tienen que ir arriba junto con la definicion de la funcion
	fmt.Println("Hola mundo desde GO")

	// Declaracion por inferencia
	var nombre string = "Dan"
	fmt.Println(nombre)
	//Declaracion rapida o corta
	nombre2 := "Israel"
	fmt.Println(nombre2)
	fmt.Printf("El valor de mi constante %s \n", MICONSTANTE)

}
*/
/*
// Tipos de datos
func main() {
	var string1 string = "Texto"
	fmt.Println(string1)
	// Se recomienda utilizar este tipo de `` para texto grande
	textoGrande := `y. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum`
	fmt.Println(textoGrande)
	var estado bool = true
	fmt.Println(estado)

	var flotante32 float32 = 32.33
	fmt.Println(flotante32)

	var flotante64 float64 = 64.33
	fmt.Println(flotante64)

	var entero int = 123
	fmt.Println(entero)

	var entero_int8 int8 = 123
	fmt.Println(entero_int8)

	var entero_int16 int16 = 123
	fmt.Println(entero_int16)

	var entero_uint8 uint8 = 123
	fmt.Println(entero_uint8)

}
*/
// Refleject y TypeOf
// Permite detectar el tipo de dato que usa una variable

/*

func main() {
	var string1 string = "texto"
	floatX := 32.4
	fmt.Println(reflect.TypeOf(string1))
	fmt.Println(reflect.TypeOf(floatX))

}
*/

// Punteros
// Permite acceder al valor binario del objeto???, es la referencia en memoria de donde se aloja el valor

/*
func main() {
	color := "rojo"
	fmt.Println(color, &color)
	var estado bool = true
	fmt.Println(&estado)
}
*/

// Condicionales

func main() {
	edad := 20

	if edad >= 18 {
		fmt.Println("Mayor de edad")
	} else {
		fmt.Println("Menor de edad")
	}

	color := "blanco"

	if color == "rojo" {
		fmt.Println(" Rojo -")
	} else if color == "blanco" {
		fmt.Println("Blanco -")
	}
	// declaracion en tiempop de ejecucion de una variable
	if variable := 2; variable == 2 {
		fmt.Println("La variable es igual a 2")
	}
	fmt.Println("CASE ")
	switch color {
	case "rojo":
		fmt.Println("rojo")

	case "blanco":
		fmt.Println("blanco")

	default:
		fmt.Println("DEFAULT ")
	}

}
