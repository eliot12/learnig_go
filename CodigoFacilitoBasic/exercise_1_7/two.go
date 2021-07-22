package main

import (
	"fmt"
	"strconv"
)

/*Escribí un programa que solicite al usuario ingresar un número con decimales y almacenalo
en una variable. A continuación, el programa debe solicitar al usuario que ingrese un número
entero y guardarlo en otra variable. En una tercera variable se deberá guardar el resultado de
la suma de los dos números ingresados por el usuario. Por último, se debe mostrar en pantalla
el texto “El resultado de la suma es [suma]”, donde “[suma]” se reemplazará por el resultado
de la operación.*/

/*Primer número: 14.2
Segundo número: 19
El resultado de la suma es 33.2*/

func main()  {
	var vardec, varentero string
	var decimal, resultado, entero float64


	fmt.Println("Ingrese el número decimal")
	fmt.Scanf("%v\n",&vardec)
	decimal,_ = strconv.ParseFloat(vardec,8)

	fmt.Println("Ingrese el número entero")
	fmt.Scanf("%v\n",&varentero)
	entero,_ = strconv.ParseFloat(varentero,8)

	resultado =  decimal + entero
	fmt.Println("la suma es:", resultado)
}
