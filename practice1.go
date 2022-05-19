package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hola mundo")
	const PI = 3.1416
	var a, b = 1, 2

	// Comentario
	/* Este es un comentario largo */

	var flag1 = true
	var flag2 = false

	fmt.Println("a + b", a+b)
	fmt.Println("a =  ", a)
	a++
	fmt.Println("a++", a)

	// + - / * *= +=/ = ++ --
	//== >= <= <> !=

	fmt.Println("f1 == f2", (flag1 == flag2))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Mi nombre es: ")
	name, _ := reader.ReadString('\n')
	fmt.Print(name)

	if strings.Compare(name, "Anubis") == 0 {
		fmt.Println("Anubis está enferma")

	} else if strings.Compare(name, "Suri\r\n") == 0 {
		fmt.Println("Suri necesita bajar de peso")

	} else {
		fmt.Println("No hay perro")
	}

	if rand_num := rand.Intn(10); rand_num == 5 {
		fmt.Println("El numero es 5")
	} else {
		fmt.Println("No es 5")
	}
}
