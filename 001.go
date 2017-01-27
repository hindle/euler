package euler

import "strconv"

func Sol001() Solution {
	sol := Solution{
		Problem: "Find the sum of all the multiples of 3 or 5 below 1000.",
		Answer:  answer001,
		Imp:     imp001,
	}

	return sol
}

func imp001() string {
	return "brute"
}

func answer001(imp string) string {
	var ans string

	switch imp {
	case "brute":
		ans = brute001()
	default:
		ans = brute001()
	}

	return ans
}

func brute001() string {
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
