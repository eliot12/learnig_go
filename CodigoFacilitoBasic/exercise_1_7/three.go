package main

import (
	"fmt"
)
 	/*
 		Escribí un programa que solicite al usuario dos números y los almacene en dos variables.
 		En otra variable, almacená el resultado de la suma de esos dos números y luego mostrá ese
 		resultado en pantalla.
 		A continuación, el programa debe solicitar al usuario que ingrese un tercer número, el cual
 		se debe almacenar en una nueva variable. Por último, mostrá en pantalla el resultado de la
 		multiplicación de este nuevo número por el resultado de la suma anterior.
 	*/
 	/*Ingresá un número: 1
	Ingresá otro número: 2
	Suman: 3
	Ingresá un nuevo número: 3
	Multiplicación de la suma por el último número: 9 */

func main()  {
	var num1, num2, num3, res1 int
	fmt.Println("Ingrese el primer número:")
	fmt.Scanf("%v\n", &num1)
	fmt.Println("Ingrese el segundo número:")
	fmt.Scanf("%v\n",&num2)
	res1 = num1 + num2
	fmt.Println("Suman:", res1)
	fmt.Println("Ingresá un nuevo número:")
	fmt.Scanf("%v\n",&num3)
	resultado :=  res1 * num3
	fmt.Println("Resultado final:", resultado)
}
