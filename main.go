package main

import (
	# Standard Packages
	"fmt"
	"log"
	
	# Self-Written Packages
	"github.com/de-wax/go-pkg/dewpoint"
	
	# Third-Party Packages
)

func main() {
	DP, err := dewpoint.Calculate(21.5, 56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DP)
}
