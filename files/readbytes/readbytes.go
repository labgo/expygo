package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("notsobig.dat")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(content), "bytes read.")
}
