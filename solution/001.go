package solution

import "strconv"

func Sol001() Solution {
	sol := Solution{
		Problem: "Find the sum of all the multiples of 3 or 5 below 1000.",
		Answer:  answer001,
	}

	return sol
}

func answer001() string {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	return strconv.Itoa(sum)
}
