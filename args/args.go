// Exibir argumentos da linha de comando

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d: %v\n", i, arg)
	}
}
