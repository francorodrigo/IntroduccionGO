package main

import (
	"fmt"
)

//Declaración de funciones
func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

//Funciones con variables con nombre. Se lo denomina "naked" return. Util en funciones cortas
func split(number int) (x, y int) {
	x = number * 2
	y = number * 3
	return
}

var variable bool

func main() {

	//Inicilizacion de variables
	var boolean, string, entero = true, "string", 1
	fmt.Println("Las variables de pruebas son:", boolean, string, entero)

	//Inicializacion de variables corta
	shortBoolean, shortString, shortEntero := true, "shortString", 2

	fmt.Println("Las variables inicializadas con la forma corta son:", shortBoolean, shortString, shortEntero)
	fmt.Printf("Tipo de dato: %T Value: %v\n", shortBoolean, shortBoolean)
	//pruebas de funciones

	// fmt.Println(add(10, 10))
	// fmt.Println(swap("hola", "pepe"))
	// fmt.Println(split(3))

	//Parseo de tipos de datos
	fmt.Printf("El valor en formato entero es: %d, y el valor en formato decimal es %.2f\n", entero, float64(entero))

	fmt.Println("Prueba de print")

}
