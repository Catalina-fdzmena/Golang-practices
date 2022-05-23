package main

import (
	"fmt"
)

//Declarar funciones con diversos tipos de parámetros

//Ejemplo generar producot punto de dos vectores

func dot_product(v1, v2 [3]float64) (float64, bool) {
	if len(v1) != len(v2) {
		return 0.0, false
	}

	if len(v1) == len(v2) {
		return 0.0, true
	}


	var result float64
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result
}


//*******************************************

//Función utilizando slices 

func dot_product2(vector1, vector2 []float64) (float64, bool){
	if len(vector1) != len(vector2) {
		return 0.0, false
	}	

	
	var result float64
	for i := 0; i < len(v1); i++ {
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

	var arreglo [5]int //Declarar memoria para el arreglo
	for i := range arreglo {
		arreglo[i] = i + 1
		fmt.Println("El %d-ésimo elemento es: %d\n", i, arreglo[i])
	}

	for value, i := range arreglo {
		fmt.Println("Soy el elemento %d y valgo %d\n", i, value)
	}

	arreglo2 := [5]int{10, 20, 30, 40, 50}
	arreglo3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} //No tiene un límite, se asigna el tamaño
	fmt.Println(arreglo2)
	fmt.Println(arreglo3)
	//**********************************************

	v1 := [3]float64{1, 2, 3}
	v2 := [3]float64{4, 5, 6}

	result.error := dot_product(v1, v2)
	result := dot_product(v1, "dot product", v2, result, "=", "error?", error)


	fmt.Println(arreglo)
	s1 := arreglo[1:3]
	s2 := arreglo[1:]
	s3 := arreglo[:3]
	s3 := arreglo[:]

	slice_directo := []int{1,2,3,4,5,6,7,8,9}

	fmt.Println(s1, "--", s2 ,"--", s3, "--", s4, slice_directo)

	vector1 := []float64{1,2,3,4}
	vector2 := []float64{1,2}

	result.error := dot_product2(vector1, vector2)
	result := dot_product2(vector1, "dot product", vector2, result, "=", "error?", error)

	s1 = append(s1, 8)
}
