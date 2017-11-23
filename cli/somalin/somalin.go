// Somar números na linha de comando

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0.0
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println("*** número inválido:", arg)
			continue
		}
		total += x
	}
	fmt.Println("Total:", total)
}
