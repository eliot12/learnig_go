package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese su nombre:")
	name,err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println(err)
	}else {
		fmt.Println("Tu nombre: "+ name)
		fmt.Println("Ahora estas en matrix, "+ name)
	}
}
