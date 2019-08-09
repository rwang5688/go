// prime_factors_app project main.go
package main

import (
	"app/prime_factors"
	"fmt"
)

func main() {
	fmt.Println("prime_factors_app.exe")

	// replicate logic of the final test
	p := &prime_factors.PrimeFactors{}
	num := 140
	fmt.Print(num, ": ")
	list := p.Generate(num)
	fmt.Println(list)
}
