package main

import (
	"fmt"
)

	//Declarar funciones con diversos tipos de parámetros 

	//Ejemplo generar producot punto de dos vectores

	func dot_product(v1, v2 [3]float64) (float64, bool) {
		if len(v1) != len(v2){
			return 0.0, false 
		}

		var result float64
		for i := 0 ; i < len(v1); i++{
			result += v1[i] * v2[i]
		}
		return result
	}

func main() {
	var valor = 3.1416
	valor2 := 3.1416789543
	var valor3 float64 = 3.14267895

	var cadena = "Hola, Mundo"
	cadena2 := "Hola, Mundo"
	var cadena3 string = "Hola, Mundo"

	fmt.Println(valor, valor2, valor3)
	fmt.Println(cadena, cadena2, cadena3)

	//**********************************************

	//Slice vistas de una arreglo que permiten manera el tamaño de los parámetros

	var arreglo [5]int  //Declarar memoria para el arreglo 
	for i:= range arreglo{
		arreglo[i] = i+1
		fmt.Println("El %d-ésimo elemento es: %d\n", i, arreglo[i])
	}

	forn value, i:=range arreglo{
		fmt.Println( "Soy el elemento %d y valgo %d\n", i, value)
	}

	arreglo2:= [5]int(10,20,30,40,50)
	arrelo3 := [...]int(1,2,3,4,5,6,7,8,9)  //No tiene un límite, se asigna el tamaño
	fmt.Println(arreglo2)
	fmt.Println(arreglo3)
	//**********************************************

	v1 := [3]float64(1,2,3)
	v2 := [3]float64(4,5,6)

	result := dot_product(v1,v2)




}
