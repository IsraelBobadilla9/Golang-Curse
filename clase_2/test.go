package main

import (
	"fmt"
	"time"
)

/*
// Funciones
// Plabara reservada func, la llave de apertura siempre va alado

func main() {

	miFuncion()
	miFUncionConParametros(5, 10)
	valorRetorno := miFuncionConRetorno("Israel")
	fmt.Println(valorRetorno)

	nombre, apellido, edad := miFuncionConRetornoMultiple("israel", "bobadilla", 20)

	fmt.Println(nombre, apellido, edad)
	//Funciones anonimas
	fmt.Println("Uso de funciones anonimas ------------------------------")
	fmt.Println(suma(10, 5))
	fmt.Println("Uso de funciones clousure ------------------------------")

	Tabla := tabla(2)
	for i := 1; i <= 10; i++ {
		fmt.Println(" 2x ", i, " = ", Tabla())
	}
}




// funcion anonima, se declara la funcion utilizando una variable
// mas usado para las conexiones a la base de datos
var suma = func(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// funciones close
// son funciones que retornan una funcio
func tabla(numero1 int) func() int {
	numero := numero1
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}

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
*/
//Goroutines
// son rutinas que se ejecutan en hilos diferentes y se pueden pausar o colocarse en un canal
// Usan el paquete time
func main() {
	fmt.Println(miFuncion("israel"))
	fmt.Println("DETENIENDO ...")
	time.Sleep(time.Second * 5)
	fmt.Println("INICIANDO")
	fmt.Println(miFuncion("DAN"))

	// ejemplo2
	// los canales permite apoderarnos de los hilos del procesador
	miCanal := make(chan string)
	go func() {
		miCanal <- miFuncion("pedro")
	}()
	fmt.Println(<-miCanal)
	fmt.Println("Continuar con la ejecucion")

}

func miFuncion(parametro string) string {
	return "hola" + parametro
}
