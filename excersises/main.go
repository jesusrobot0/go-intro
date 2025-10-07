package main

import "fmt"

func main() {
	// Primer ejercicio
	// array := [5]int{4, 2, 5, 6, 7}

	// // realizar la funcionalidad

	// for i, v := range array {
	// 	array[i] = v + 20
	// }

	// fmt.Println("Los valores del array son: ", array)

	// Segundo ejercicio
	// var datosIngresados []int
	// var datoIngresado int

	// for {
	// 	fmt.Print("Ingresa un número (0 para salir): ")
	// 	_, err := fmt.Scanln(&datoIngresado)
	// 	if err != nil {
	// 		fmt.Println("Error al leer el número:", err)
	// 		continue
	// 	}

	// 	if datoIngresado == 0 {
	// 		break
	// 	}

	// 	datosIngresados = append(datosIngresados, datoIngresado)
	// }

	// fmt.Println("Datos ingresados:", datosIngresados)

	// Ejercicio 3

	productos := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	productosSelecciados := []string{}
	productoActual := ""

	for {
		fmt.Print("Ingresa el código del producto (0 para salir): ")

		_, err := fmt.Scanln(&productoActual)
		if err != nil {
			fmt.Println("Error al leer el código del producto:", err)
			continue
		}

		if productoActual == "0" {
			break
		}

		if _, exists := productos[productoActual]; !exists {
			productosSelecciados = append(productosSelecciados, "Producto no encontrado")
			continue
		}

		productosSelecciados = append(productosSelecciados, productos[productoActual])
	}

	fmt.Println("Productos seleccionados:", productosSelecciados)

}
