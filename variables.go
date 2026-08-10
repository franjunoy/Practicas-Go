package main

import (
	"fmt"
	"strings"
)

func variables() {
	//numeros
	var entero int = 10
	entero2 := 10
	var floatante float64 = 10.5
	floatente2 := 10.5
	suma := entero + int(floatante)

	//strings
	var nombre string = "Juan"
	apellido := "Perez"

	concatenado := nombre + " " + apellido + "" +"10"
	enMayusculas := strings.ToUpper(concatenado)

	//arrays
	var arrayFijo = [6]int{1, 2, 3, 4, 5, 6}
	arrayFijo2 := [5]int{1, 2, 3, 4, 5}

	var sliceDinamico = []int{1, 2, 3, 4, 5}
	sliceDinamico2 := []int{1, 2, 3, 4, 5}

	sliceDinamico = append(sliceDinamico, 6)

	//mapas
	diccionario := map[string]int{
		"Juan": 25,
		"Maria": 30,
	}

	//struct
	type Persona struct {
		Nombre string
		Edad int
	}

	persona := Persona{
		Nombre: "Juan",
		Edad: 25,
	}

	fmt.Println("Hello, World!")
	fmt.Println(entero)
	fmt.Println(entero2)
	fmt.Println(floatante)
	fmt.Println(floatente2)
	fmt.Println(suma)
	fmt.Println(concatenado)
	fmt.Println(enMayusculas)
	fmt.Println(arrayFijo)
	fmt.Println(arrayFijo2)
	fmt.Println(sliceDinamico)
	fmt.Println(sliceDinamico2)
	fmt.Println(diccionario)
	fmt.Println("Esta persona es", persona.Nombre, "y su edad es", persona.Edad)
}