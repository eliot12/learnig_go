package main

import "fmt"

func main()  {
	/* for
	for i := 0; i <10; i++ {
		fmt.Println("Hello Word", i)
	} */

	/* While
	i:= 0
	for i < 10 {
		fmt.Println(i)
		i++
	} 	 */
	/*
	i:=0
	for  {
		fmt.Println(i)
		i++
		if i >10 {
			break
		}
	} */

	// continue
	i:= 0
	for  {
		if i == 2 {
			i++
			continue
		}
		fmt.Println(i)
		i++

		if i > 10{
			break
		}
	}
}
