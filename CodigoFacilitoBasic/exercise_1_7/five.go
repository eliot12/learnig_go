package main

import (
	"fmt"
)

func main()  {
	/* Escribí un programa que solicite al usuario el ingreso de una temperatura en escala
	Fahrenheit (debe permitir decimales) y le muestre el equivalente en grados Celsius.
	La fórmula de conversión que se usa para este cálculo es: _Celsius = (5/9) * (Fahrenheit-32)_ */

	/*Ingresá una temperatura expresada en Farenheit: 75
	23.88888888888889*/
	var fare, celsius, a,b float64
	a = 5
	b = 9
	fmt.Println("Ingresá la temperatura expresada en Farenheit")
	fmt.Scanf("%v\n",&fare)
	celsius = (a/b)*(fare-32)
	fmt.Println("La conversión en Celsius es:", celsius)
}