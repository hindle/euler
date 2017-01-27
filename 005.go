package euler

import "strconv"

func Sol005() Solution {
	sol := Solution{
		Problem: "What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?",
		Answer:  answer005,
		Imp:     imp005,
	}

	return sol
}

func imp005() string {
	return "brute"
}

func answer005(imp string) string {
	var ans string

	switch imp {
	case "brute":
		ans = gcd005()
	default:
		ans = gcd005()
	}

	return ans
}

func gcd005() string {
	lcmVal := 2

	for i := 1; i <= 20; i++ {
		lcmVal = lcm(i, lcmVal)
	}

	return strconv.Itoa(lcmVal)
}

// GCD = Greatest Common Divisor
func gcd(a int, b int) int {
	var retVal int

	//if a is less than b, switch the numbers
	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	if a == 0 {
		retVal = b
	} else if b == 0 {
		retVal = a
	} else {
		r := a % b
		retVal = gcd(b, r)
	}

	return retVal
}

//LCM = Lowest Common Multiple
func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}
