package main

import "fmt"

func main() {
	fmt.Println("Operadores lógicos")

	edad := 18


	fmt.Println(edad > 18) // true
	fmt.Println(edad < 18) // false
	fmt.Println(edad >= 18) // true
	fmt.Println(edad <= 18) // true
	fmt.Println(edad == 18) // true
	fmt.Println(edad != 18) // false

	fmt.Println(edad > 18 && edad < 21) // true
	fmt.Println(edad > 18 || edad < 21) // true
	fmt.Println(!(edad > 18 && edad < 21)) // false
	fmt.Println(!(edad > 18 || edad < 21)) // false

	/* 	
		Esta es una condición compleja, que se puede entender mejor si la ponemos en	 este caso: 
		Una persona puede acceder a un servicio si es mayor de 18 años y menor de 25 años o
		mayor de 60 años y siempre sea menor de 120 años
	 */
	fmt.Println(edad >= 18 && edad <= 25 || edad >= 60 && edad < 120) // true

	/*
		Tabla de verdad || (or)
		true || true = true
		true || false = true
		false || true = true
		false || false = false

		Tabla de verdad && (and)
		true && true = true
		true && false = false
		false && true = false
		false && false = false

		Tabla de verdad ! (not)
		!true = false
		!false = true
	*/

	fmt.Println("Estructura de control if")

	// Permite ejecutar un bloque de código u otro si una condición es true
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// Eres apto para recibir un programa social

	edadUsuario := 17

	if edadUsuario >= 12 && edadUsuario <= 18 {
		fmt.Println("Eres apto para recibir el programa de estudiantes")
	} else if edadUsuario >= 18 && edadUsuario <= 30 || edadUsuario >= 18 && edadUsuario <= 35  {
		fmt.Println("Eres apto para recibir el programa de educación universitaria o de trabajo")
	} else if edadUsuario >= 60 && edadUsuario <= 120 {
		fmt.Println("Eres apto para recibir el programa de tercera edad")
	}else {
		fmt.Println("No eres apto para recibir ningún programa")
	}

	if value := true; value {
		fmt.Println("El valor es true")
	}

	// Conversión de la estructura if-else a switch
	switch {
	case edadUsuario >= 12 && edadUsuario <= 18:
		fmt.Println("Eres apto para recibir el programa de estudiantes")
	case (edadUsuario >= 18 && edadUsuario <= 30) || (edadUsuario >= 18 && edadUsuario <= 35):
		fmt.Println("Eres apto para recibir el programa de educación universitaria o de trabajo")
	case edadUsuario >= 60 && edadUsuario <= 120:
		fmt.Println("Eres apto para recibir el programa de tercera edad")
	default:
		fmt.Println("No eres apto para recibir ningún programa")
	}

	// Hay programas sociales por estado

	estado := "Hidalgo"
	estado = "Chihuahua"

	switch estado {
	case "Hidalgo":
		fmt.Println("Hay programas activos en tu estado")
	case "México":
		fmt.Println("Hay programas activos en tu estado")
	default:
		fmt.Println("No hay programas activos en tu estado")
	}

	switch edad := 25; edad {
	case 12:
		fmt.Println("El valor es el número 12")
	case 18:
		fmt.Println("El valor es el número 18")
	case 21:
		fmt.Println("El valor es el número 21")
	case 30:
		fmt.Println("El valor es el número 30")
	case 35:
		fmt.Println("El valor es el número 35")
	case 60:
		fmt.Println("El valor es el número 60")
	case 120:
		fmt.Println("El valor es el número 120")
	default:
		fmt.Println("El valor no es el número 12, 18, 21, 30, 35, 60 o 120")
	}

	// Homework

	{
		nombre:= "Juanito"
		edad := 10
		licencia := true

		if (edad < 15 || licencia != true) {
			fmt.Printf("%s No puede seguir avanzado \n", nombre)

			if licencia {
				fmt.Println("Tu licencia es falsa!")
			}
		} else {
			fmt.Printf("%s Puede seguir avanzado \n", nombre)
		}
	}

}