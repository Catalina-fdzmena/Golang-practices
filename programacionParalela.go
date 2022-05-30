package main

import (
	"fmt"
	"time"
)

func hola_hilo() {
	fmt.Println("Hola, soy Panchito concurrente")
}

func main() {
	go hola_hilo()
	go hola_hilo()
	go hola_hilo()
	go hola_hilo()
	go hola_hilo()

	time.Sleep(30 * time.Second)
}

//Inicia hilo procesamiento pero no lo marca de regreso por el tiempo que toma un cambio de contexto.
