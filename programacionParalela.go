package main

import (
	"fmt"
	"time"
)

func hola_hilo(id int) {
	fmt.Printf("Hola, soy Panchito %d concurrente\n", id)
}

func main() {
	for i := 0; i < 20; i++ {
		go hola_hilo(i)
	}

	time.Sleep(30 * time.Second) //Con este comando se desarrolla comando para desarrolle un intervalo de la reproducción
}

//Inicia hilo procesamiento pero no lo marca de regreso por el tiempo que toma un cambio de contexto.
