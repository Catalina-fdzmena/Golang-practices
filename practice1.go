package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola mundo")
	const PI = 3.1416
	distancia2 := 3
	distancia = 4
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
}
