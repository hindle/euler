package euler

import "strconv"

func Sol002() Solution {
	return Solution{
		Problem: "By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.",
		Answer:  answer002,
		Imp:     imp002,
	}
}

func imp002() string {
	return "brute"
}

func answer002(imp string) string {
	var ans string

	switch imp {
	case "brute":
		ans = brute002()
	default:
		ans = brute002()
	}

	return ans
}

func brute002() string {
	prevNum := 1
	num := 2
	sum := 0

	for cont := true; cont; cont = (num <= 4000000) {
		if num%2 == 0 {
			sum += num
		}

		tempNum := num
		num += prevNum
		prevNum = tempNum
	}

	return strconv.Itoa(sum)
}
