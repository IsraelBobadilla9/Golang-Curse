package main // modulo

import "fmt" // Estructuras

// es un objeto de datos con diferentes campos
// igual los nombres de los campos tienen que ser publicos por eso se colocan asi
type Persona struct {
	Id     int
	Nombre string
	Correo string
	Edad   int
}

type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}

type Producto struct {
	Id          int
	Nombre      string
	Slug        string
	Precio      int
	CategoriaId Categoria
}

/*
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

	// estructura anidada | herencia
	fmt.Println("================ESTRUCTURA ANIDADA ================")

	categoria := Categoria{Id: 1, Nombre: "Categoria 1 ", Slug: "Categoria-1"}
	producto := Producto{Id: 1, Nombre: "Mesa de computadora", Slug: " mesa-de-computadora", Precio: 1234, CategoriaId: categoria}
	fmt.Printf("%+v \n", producto)

}
*/
// Interfaces
type EjemploInterface interface {
	miFuncion() string
	Calculo(n1 int, n2 int) int
}
type Estructura struct {
}

// con * le indicamos que forma parte del contratao de estructura
func (*Estructura) miFuncion() string {
	return "texto desde mi funcion"
}

func (*Estructura) calculo(n1 int, n2 int) int {
	return n1 * n2
}

func main() {

	e := Estructura{}
	imprimir := e.miFuncion()
	fmt.Println(imprimir)
	resultado := e.calculo(10, 5)
	fmt.Println(resultado)
}
