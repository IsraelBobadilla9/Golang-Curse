package main

import (
	"fmt"
	"time"
)

func main() {
	// Obtener fecha actual del servidor
	fmt.Println("La hora actual es: ", time.Now())

	fecha := time.Now()

	fmt.Println("El anio es: ", fecha.Year())
	fmt.Println("Mes: ", int(fecha.Month()))
	fmt.Printf("El tipo es %T \n", fecha)
	fmt.Printf("%v/%v/%v \n", fecha.Day(), fecha.Month(), fecha.Year())

	fmt.Println("Fecha menos 20 dias ")
	fecha2 := fecha.Add(time.Hour * 24 * -20)
	fmt.Println(fecha2)

}
