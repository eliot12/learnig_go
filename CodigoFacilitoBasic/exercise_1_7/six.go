/*Escribí un programa que solicite al usuario ingresar tres números
	para luego mostrarle el promedio delos tres.*/

/*Primer número: 8.5
Segundo número: 10
Tercer número: 5.5
El promedio de los tres es 8.0*/
package main

import "fmt"

func main()  {
	var num1, num2, num3, aux float64
	fmt.Println("Ingrese el primer número:")
	fmt.Scanf("%v\n",&num1)
	aux = aux + 1
	fmt.Println("Ingrese el segundo número:")
	fmt.Scanf("%v\n",&num2)
	aux = aux + 1
	fmt.Println("Ingrese el tercer número:")
	fmt.Scanf("%v\n",&num3)
	aux = aux + 1
	sumatoria := num1 + num2 + num3
	resultado := sumatoria / aux
	fmt.Println("El promedio es:", resultado)
}
