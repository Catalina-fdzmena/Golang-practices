package main

import (
	"fmt"
)

func DotProduct(v1, v2 [3]float64) (float64, bool) {
	if len(v1) != len(v2) {
		return 0.0, true
	}
	var result float64
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result, false
}

func DotProduct2(v1, v2 []float64) (float64, bool) {
	if len(v1) != len(v2) {
		return 0.0, false
	}
	var result float64
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result, true
}

func main() {
	//var valor = 3.141592654
	//valor2 := 3.141592654
	//var valor3 float64 = 3.141592654

	//var cadena = "hola,mundo"
	//cadena2 := "hola,mundo"
	//var cadena3 string = "hola,mundo"
	//fmt.Println(valor, valor2, valor3)
	//fmt.Println(cadena, cadena2, cadena3)

	var arreglo [6]int

	for i := range arreglo {
		arreglo[i] = i + 1
		//fmt.Printf("El %d-ésimo elemento es: %d\n", i, arreglo[i])
	}

	for value, i := range arreglo {
		fmt.Printf("Soy el elemento %d y valgo %d\n", i, value)
	}

	//arreglo2 := [5]int{10, 20, 30, 40, 50}
	//arreglo3 := [...]int{1, 2, 34, 4, 5, 6, 7, 8, 9}
	//fmt.Println(arreglo2)
	//fmt.Println(arreglo3)
	v1 := [3]float64{1, 2, 3}
	v2 := [3]float64{4, 5, 6}
	result, error := DotProduct(v1, v2)
	fmt.Println(v1, " * ", v2, " = ", result, " error? ", error)

	/*****************************************/
	//fmt.Println(arreglo)
	//s1 := arreglo[1:3]
	//s2 := arreglo[1:]
	//s3 := arreglo[:3]
	//s4 := arreglo[:]
	//slice_directo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//fmt.Println(s1, "--", s2, "--", s3, "--", s4, "--", slice_directo)
	vector1 := []float64{1, 2, 3, 4}
	vector2 := []float64{1, 2}

	result, error = DotProduct2(vector1, vector2)
	fmt.Println(vector1, " * ", vector2, " = ", result, " error? ", error)

	sk := arreglo[1:2:3]
	fmt.Println(arreglo, sk)
	sk = append(sk, 7)
	fmt.Println(arreglo, sk)
	sk = append(sk, 8)
	fmt.Println(arreglo, sk)
	sk = append(sk, 9)
	fmt.Println(arreglo, sk)
	sk = append(sk, 10)
	fmt.Println(arreglo, sk)

	/*
		fmt.Println(arreglo, s1)
		s1 = append(s1, 8)
		fmt.Println(arreglo, s1)
		s1 = append(s1, 9)
		fmt.Println(arreglo,  s1)
		s1 = append(s1, 10)
		fmt.Println(arreglo,  s1)
		s1 = append(s1, 11)
		fmt.Println(arreglo,  s1)
	*/
}
