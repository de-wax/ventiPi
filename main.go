package main

import (
	"fmt"
	"log"
	github.com/de-wax/ventiPi/pkg/dewpoint
)

func main() {
	DP, err := dewpoint.Calculate(21.5, 56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DP)
}
