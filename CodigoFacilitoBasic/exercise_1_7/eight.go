/* Escribí un programa que solicite al usuario el ingreso de dos palabras, las cuales se guardarán en
dos variables distintas. A continuación, almacená en una variable la concatenación de la primera palabra,
más un espacio, más la segunda palabra. Mostrá este resultado en pantalla.*/

/*Primera palabra: derribando
Segunda palabra: muros
derribando muros*/
package main

import "fmt"

func main()  {
	var letra, letra1 string
    fmt.Println("Escriba la primera letra:")
    fmt.Scanf("%v\n",&letra)
	fmt.Println("Escriba la segunda letra")
	fmt.Scanf("%v\n",&letra1)
	fmt.Printf(letra + " " + letra1)
}
