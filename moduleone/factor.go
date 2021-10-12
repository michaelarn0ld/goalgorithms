package moduleone

func Factor(primes []int, number int) []int {
	var res []int
	for _, prime := range primes {
		for number%prime == 0 {
			res = append(res, prime)
			number /= prime
		}
	}

	if number > 1 {
		res = append(res, number)
	}

	return res
}
