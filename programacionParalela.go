package main

import (
	"fmt"
)

func hola_hilo() {
	fmt.Println("Hola, soy Panchito concurrente")
}

func main() {
	go hola_hilo()
}

//Inicia hilo procesamiento pero no lo marca de regreso por el tiempo que toma un cambio de contexto.
