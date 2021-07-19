package main

import "fmt"

func main()  {
	/*Escribí un programa que solicite al usuario que ingrese su nombre.
	El nombre se debe almacenar en una variable llamada nombre. A continuación
	se debe mostrar en pantalla el texto “Ahora estás en la matrix, [usuario]”,
	donde “[usuario]” se reemplazará por el nombre que el usuario haya ingresado.*/
	var name string
	fmt.Println("Introduce tú nombre")
	fmt.Scanf("%v\n",&name)
	fmt.Println("Tú nombre: ",name)
	fmt.Println("Ahora estás en la matrix,"+name)
}
