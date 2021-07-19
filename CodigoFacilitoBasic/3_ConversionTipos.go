package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22
    year := "50"
	edad_str := strconv.Itoa(edad) // De Entero a un String
    year_int,_ := strconv.Atoi(year) // De String a Entero = Ojo devuelve varias variables solucionamos con _
	fmt.Println("Mi edad es " + edad_str)
    fmt.Println("La edad el es:" , year_int)
}