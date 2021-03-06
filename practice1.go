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

	//Estrcutura de switch cases
	switch name {
	case "Anubis":
		fmt.Println("Anubis es el perro")
	case "Suri":
		fmt.Println("Suri es el perro")
	default:
		fmt.Println("Max es el perro")
	}

	//

	//Otra versión de hacer el swtich

	switch {
	case name == "Anubis":
		fmt.Println("Anubis es el perro")
	case name == "Suri":
		fmt.Println("Suri es el perro")
	default:
		fmt.Println("Max es el perro")
	}

	//Palabra reservada fallthrough

	//Es una cascada para que se pase al siguiente pre de igual manera se muestra lo que estaba previamente

	switch {
	case name == "Anubis":
		fallthrough
	case name == "Suri":
		fmt.Println("Suri es el perro")
	default:
		fmt.Println("Max es el perro")
	}

	//SAsignación de variables
	switch rand_num := rand.Intn(2); rand_num {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")

	default:
		fmt.Println("Otro número")
	}

	//Trabajar con ciclos

	fmt.Println("++++++++++++++++++++++++++++++++++++++++")

	var count = 0

	for {
		if count == 5 {
			fmt.Println("El count es 5, bye!")
		} else if count == 4 {
			count++
			continue
		}
	}

	//Contador 2 vive unicamente en el for
	for Contador = 0; Contador < 5; Contador++ {
		fmt.Println("El count es: ", Contador)
	}

	//COntador 2 vive unicamente en el for
	for contador2 := 0; contador2 < 5; contador2++ {
		fmt.Println("El contador2 es: ", contador2)
	}

}
