package main

import "fmt"

func main()  {
	//Leyendo con enteros
	var edad int
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n",&edad)
	fmt.Printf("Mi edad es %d\n",edad)
	// Leyendo con string
	var name string
	fmt.Println("Coloque su noombre: ")
	fmt.Scanf("%v\n",&name)
	fmt.Print("Tú nombre es: ", name)
	/* Println tiene un salto de linea \n
	fmt.Println("Hello Word")
	// Imprimiendo con diferentes verbos Printf sin salto de linea %v %t %s
	edad := 22
	fmt.Printf("Mi edad es %d",edad)*/
}