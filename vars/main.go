package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Variables

	// Podemos declarar variables de diferentes formas
	var name string = "Jesús"

	// Otra forma de declarar variables es usando el operador :=
	apellido := "Velasco"
	apellido = "Murillo"
	
	fmt.Println(name + " " + apellido)
	
	// Constantes
	const pi = 3.1416
	//p = 3 // error: no se puede reasignar el valor de una constante

	fmt.Printf("\nEl valor del número PI: %f", pi)
	
	// Tipos de datos básicos
	var entero int = 100
	var decimal float64 = 10.5
	var booleano bool = true
	var cadena string = "Hola"
	
	fmt.Printf(`
		El valor del entero es: %d
		El valor del decimal es: %f
		El valor del booleano es: %t
		El valor de la cadena es: %s
	`, entero, decimal, booleano, cadena)

	/*
		Hay otros valores más específicos 
		- int8, int16, int32, int64
		- uint8, uint16, uint32, uint64
		- float32, float64
		- complex64, complex128
	
		Alias (semánticos)
		Son nombres más descriptivos para los tipos de datos
		Usar uint8 transmite “número de 0 a 255”.
		Usar byte transmite “unidad de datos binarios o texto crudo”.
		Usar int32 transmite “entero”.
		Usar rune transmite “carácter Unicode”.
		- byte -> alias para uint8
		- rune -> alias para int32
	
		- string
		- bool (true o false)
	*/

	// rune y byte

	var letra rune = 'A'
	var bytee = byte(letra)
	fmt.Printf("\nEl valor de la letra es: %c y el valor del byte es: %d", letra, bytee)

	nombre := "Jesus 🤖"
	nombreSliceBytes := []byte(nombre)
	nombreSliceRunes := []rune(nombre)
	fmt.Printf("\nEl valor del nombre en bytes es: %v y el valor del nombre en runes es: %v", nombreSliceBytes, nombreSliceRunes)

	// Valores por defecto
	var a int
	var b float64
	var c bool
	var d string

	fmt.Printf(`
		El valor del entero es: %d
		El valor del decimal es: %f
		El valor del booleano es: %t
		El valor de la cadena es: %s
	`, a, b, c, d)
	
	/*
		Los valores por defecto son:
		- int: 0
		- float64: 0.0
		- bool: false
		- string: ""
		- complex64: 0+0i
		- complex128: 0+0i
		- byte: 0
		- rune: 0
		- byte: 0
		- rune: 0
		
		En general si es un tipo numérico, el valor por defecto es 0

		nil para los punteros, slices, maps, channels, funciones y interfaces
	*/


	// Conversiones de tipos
	var e int = 10
	var f float64 = 3.14
	var g int = int(f)
	fmt.Printf("El valor de e es: %d y el valor de g es: %d", e, g)

	suma := e + g
	suma2 := float64(e) + f
	fmt.Printf("\nEl valor de la suma es: %d", suma)
	fmt.Printf("\nEl valor de la segunda suma es: %.2f", suma2)

	añoScrapeado := "2025"
	numero, _ := strconv.Atoi(añoScrapeado)
	proximoAnio := numero + 1
	fmt.Printf("\nEl próximo año es: %d", proximoAnio)

	fmt.Printf("\n %T %d", numero, numero)
	

}
