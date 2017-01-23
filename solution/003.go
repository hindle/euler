package solution

import "strconv"

func Sol003() Solution {
	sol := Solution{
		Problem: "What is the largest prime factor of the number 600851475143?",
		Answer:  answer003,
	}

	return sol
}

//Brute force approach
func answer003() string {

	num := 600851475143
	primeFactors := make([]int, 1)
	divisor := 2

	for cont := true; cont; cont = (num > 1) {
		if num%divisor == 0 {
			primeFactors = append(primeFactors, divisor)
			num = num / divisor
		}
		divisor += 1
	}

	largestPrimeFactor := primeFactors[len(primeFactors)-1]

	return strconv.Itoa(largestPrimeFactor)
}
