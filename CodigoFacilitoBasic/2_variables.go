package main

import "fmt"
func main() {
	/* Declaration variables var + name variable + type date | value initial int =0 string = "" bool = false*/
	var x, y, z int
	var cadena string
	var bandera bool
	/* Declare variable and not used is error*/

	/* Initiation variables */
	x = 23
	y = 10
	z = 20
	cadena = "Hello Word"
	bandera = false

	/* Operator the := no se necesita colocar el tipo de variable es otra manera de
	cuando se usa este operador tiene que tener una nueva variable
	ademas este operador tiene un siclo de vida en la funcion*/
	name := "coco"
    fmt.Println(x,y,z,cadena,bandera,name)
}
