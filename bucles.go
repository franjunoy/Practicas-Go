package main

import (
	"fmt"
)

func condicional() {

	defer fmt.Println("Fin del programa")

	//if
	edad := 20

	if edad < 18{
		println("Eres menor de edad")
		return
	}
	fmt.Println("Eres mayor de edad")

	//for
	for i := 0; i<10; i++{
		fmt.Printf("Iteracion %d\n", i)
	}

	//while
	n := 0
	for n < 10 {
		fmt.Printf("Iteracion %d\n", n)
		n++
	}

	//bucle infinito
	/* n = 0
	for{
		n++
		if n > 5 {
			continue
		}

		fmt.Printf("Iteracion %d\n", n)

		if n >= 7 {
			break
		}
	} */

	//range
	slice := []string{"uno", "dos", "tres"}

	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	slice2 := []string{"cuatro", "cinco", "seis"}

	for index := range slice2 {
		fmt.Printf("Index: %d\n", index,)
	}

	//switch
	valor := 2
	switch valor {
	case 1:
		fmt.Println("Valor es 1")
	case 2:
		fmt.Println("Valor es 2")
	default:
		fmt.Println("Valor no es 1 ni 2")
	}

}