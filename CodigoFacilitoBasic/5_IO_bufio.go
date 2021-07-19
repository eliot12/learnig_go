package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	name,err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println(err)
	}else{
	fmt.Println("Hola "+ name)
	}
}
