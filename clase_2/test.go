package main

import "fmt"

// Funciones
// Plabara reservada func, la llave de apertura siempre va alado

func main() {

	miFuncion()
	miFUncionConParametros(5, 10)

}

// Declaracion de funcion sin modularizar

func miFuncion() {
	fmt.Println("created function")
}

func miFUncionConParametros(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("Parametro uno ", n1, "Segundo paramtro ", n2)
	fmt.Println("La suma es ", resultado)
}
