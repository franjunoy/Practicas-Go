package main

import (
	"errors"
	"fmt"
)

func suma(a, b int) int {
	fmt.Printf("La suma de %d y %d es %d", a, b, a+b)
	return a + b
}

func dividir(a, b float64) (float64, error){
	if b == 0 {
		return 0, errors.New("No se puede dividr entre 0")
	}
	cociente := a / b
	return cociente, nil
}