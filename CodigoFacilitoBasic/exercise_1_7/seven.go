/* Escribí un programa que solicite al usuario un número y le reste el 15%,
almacenando. A continuación, mostrar el resultado final en pantalla.*/

/*Ingresá un número: 260
Descontando el 15% queda: 221.0*/
package main

import "fmt"

func main()  {
	var monto, porcent, porcentaje float64
	fmt.Println("Ingrese el Monto:")
	fmt.Scanf("%v\n",&monto)
	fmt.Println("Ingrese el Porcentaje %:")
	fmt.Scanf("%v\n",&porcent)
	porcentaje = (monto * porcent) / 100
	resultado := (monto - porcentaje)
	fmt.Printf("Descuento del %v porciento saldo %v",porcent, resultado)
}