// prime_factors project prime_factors.go
package prime_factors

type PrimeFactors struct {
}

// initialize struct variables
func (p *PrimeFactors) Init() {
}

// generate prime factors
func (p *PrimeFactors) Generate(num int) []int {
	result := []int{}

	// outer loop: cycle through prime factors while n > 1
	// inner loop: cycle through all factors of pf while n is divisible by pf
	n := num
	for pf := 2; n > 1; pf++ {
		// note if pf is not a prime factor,
		// all of its factors would have already been extracted;
		// execution will skip inner loop
		for n%pf == 0 {
			result = append(result, pf)
			n /= pf
		}
	}

	return result
}
