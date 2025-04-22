package main

import "fmt"

// Funciones
// Plabara reservada func, la llave de apertura siempre va alado

func main() {

	miFuncion()
	miFUncionConParametros(5, 10)
	valorRetorno := miFuncionConRetorno("Israel")
	fmt.Println(valorRetorno)

	nombre, apellido, edad := miFuncionConRetornoMultiple("israel", "bobadilla", 20)

	fmt.Println(nombre, apellido, edad)
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

func miFuncionConRetorno(nombre string) string {
	return "Retornando valor desde la funcion " + nombre
}

func miFuncionConRetornoMultiple(nombre string, apellido string, edad int) (string, string, int) {
	return "Hola " + nombre, "Apellido: " + apellido, edad
}
