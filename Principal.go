package main

import "fmt"

func main() {

	fmt.Print("XD")
	fmt.Println("JEJEJEJE")
	n := "Lux"
	e := 19

	// %s tipo string   %d tipo de dato entero  %v no se sabe el tipo de dato que se va a agregar
	// %T para saber que tipo de dato es
	fmt.Printf("Mi nombre es, %s y tengo %d \n", n, e)

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&n)
	print("Otro nombre: ", n)

}
