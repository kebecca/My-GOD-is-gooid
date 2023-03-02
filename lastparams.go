package main

import (
	"fmt"
	"os"
)

/*
	func main() {
		tab := os.Args
		  index := len(os.Args) - 1
		  fmt.Println(os.Args[index])

}/*
/*func main() {

		//tab := os.Args
		index := len(tab) - 1
		fmt.Println(os.Args[index])
	}
*/
func main() {

	tab := os.Args
	index := len(tab) - 1
	fmt.Println(tab[index])

}
