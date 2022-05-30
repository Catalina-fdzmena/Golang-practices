//4 hilos
//Inputs: valor a buscar, arreglo con longitud par

//Outputs: indice del arreglo donde está el valor deseado

package main

import (
	"fmt"
)

func buscador(arr []int, num int, low int, high int, c chan int) {
	for i := low; i < high; i++ {
		if arr[i] == num {
			c <- i
		}
	}
	c <- -1
}

func main() {
	c := make(chan int)
	arreglo := [8]int{1, 3, 9, 4, 8, 6, 2, 5}

	go buscador(arreglo[:], 7, 0, 2, c)
	go buscador(arreglo[:], 5, 2, 4, c)
	go buscador(arreglo[:], 10, 4, 6, c)
	go buscador(arreglo[:], 2, 6, 8, c)

	// time.Sleep(100 * time.Second)

	for i := 0; i < 4; i++ {
		index := <-c
		if index != -1 {
			fmt.Println("El valor esta en el índice ", index)
		}
	}

}
