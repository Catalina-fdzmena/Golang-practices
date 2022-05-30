package main

import (
	"fmt"
)

func hola_hilo(id int, c chan int) {
	fmt.Printf("Hola, soy Panchito %d concurrente\n", id)
	c <- id
}

func main() {
	//Canles para comunicación de hilos

	c := make(chan int)

	for i := 0; i < 20; i++ {
		go hola_hilo(i, c)

	}

	for i := 0; i < 20; i++ {
		resultado := <-c
		if resultado == i {
			fmt.Println("Un hilo ha acabado", resultado)
		}
	}
	//time.Sleep(30 * time.Second) //Con este comando se desarrolla comando para desarrolle un intervalo de la reproducción en bajo nivel
}

//Inicia hilo procesamiento pero no lo marca de regreso por el tiempo que toma un cambio de contexto.
