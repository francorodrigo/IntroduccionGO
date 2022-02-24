package main

import "fmt"

func main() {

	//FOR
	for i := 0; i < 10; i++ {
		fmt.Println("El valor de i es:", i)
	}

	//FOR sin asignacion de variables. En Go no existe el while para iterar
	j := 0
	for j < 10 {
		j++
		fmt.Println("El valor de j es :", j)
	}
	sum := 10
	if sum > 10 {
		fmt.Println("sum es mayor que 10")
	}

	//if con variable declarada en statement

	if v := sum + 1; v > 11 {
		fmt.Println("El valor es mayor que 11")
	} else {
		//Si en los statements de un if hago una declaracion de variable, esta variable es accesible desde los else
		fmt.Println("El valor no es mayor que 11. El valor es :", v)
	}
	//switch
	condicion := "Condicion 2"
	switch condicion {
	case "Condicion 1":
		fmt.Println("Se asigno la condicion 1")
		break
	case "Condicion 2":
		fmt.Println("Se asigno la condicion 2")
		break
	case "Condicion 3":
		fmt.Println("Se asigno la condicion 3")
		break
	default:
	}

	//Switch sin definicion de variable a utilizar
	var t int = 1
	switch {
	case t == 1:
		fmt.Println("t es igual a 1")
	default:
	}

	//Funcion defer : se ejecuta luego de que finalice la funcion circundante(la que la engloba). sin embargo, los parametros se analizan a pesar de que no se ejecuta
	defer fmt.Println("Hola diferido")
	//Las funciones defer se stackean en una pila lifo
	defer fmt.Println("Hola diferido2")
	fmt.Println("Hola normal")

	//Ejemplo contador invertido:
	for i := 0; i < 10; i++ {
		defer fmt.Println("El valor de i es:", i)
	}
}
