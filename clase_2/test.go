package main

import (
	"fmt"
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

/*
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
*/
/*
//Recursividad

func main() {
	miFuncion(8)

}

func miFuncion(valor int) {
	dato := valor + 1
	fmt.Println(dato)
	if dato != 0 && dato < 10 {
		miFuncion(dato - 2)
	}
}
*/

// Defer y panic
// no tiene manejo de excepciones
func main() {
	miFuncion()
}

// Panic permite mostrar un mensaje de error para deter el script, a nivel de terminal
func miFuncion() {

	// usado para cerrar la conexion de la base de datos
	defer fmt.Println("Dentro de un defer") // esto se ejecutara cuando se termine la ejecucion del script

	fmt.Println("Este es un mensaje ")

	a := 1
	if a == 1 {
		panic("Se encontro un error en mifuncion linea 136") // pausa la ejecucion
	}

}
