package main

import "fmt"

func main()  {
	/*Escribí un programa que solicite al usuario ingresar la cantidad de kilómetros recorridos
	por una motocicleta y la cantidad de litros de combustible que consumió durante ese recorrido.
	Mostrar el consumo de combustible por kilómetro.*/

	/*Kilómetros recorridos: 260
	Litros de combustible gastados: 12.5
	El consumo por kilómetro es de 20.8 */

	var kilometro, combustible float64
	fmt.Println("Kilómetros Recorridos:")
	fmt.Scanf("%v\n",&kilometro)
	fmt.Println("Litros de combustible gastados:")
	fmt.Scanf("%v\n",&combustible)
	resul := kilometro / combustible
	fmt.Println("El consumo por kilómetro es de:",resul)
}
