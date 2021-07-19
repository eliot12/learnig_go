package main

import "fmt"

func main()  {
	//if evalua una expresion que tiene que devolver verdadero o falso y partir de eso ejecuta algun proceso

	/*
		== igual a
		!= diferente
		<  menor que
		>  mayor que
		>= mayor igual
		<= menor igual
		&& AND Permite concatenar condiciones y funciona si todas las condiciones son verdaderas
		|| OR Concatena expresiones y devualve verdadero con tan solo una sea verdadero
	*/

	x := 9
	y := 9
	if x > y {
		fmt.Printf("%d es mayor que %d \n",x,y)
	}else if x < y{
		fmt.Printf("%d es menor a y",x )
	} else {
		fmt.Printf("son iguales")
	}
}
