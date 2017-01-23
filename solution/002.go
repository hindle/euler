package solution

import "strconv"

func Sol002() Solution {
	return Solution{
		Problem: "By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.",
		Answer:  answer002,
	}
}

func answer002() string {
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
