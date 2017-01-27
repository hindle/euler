package euler

import (
	"math"
	"strconv"
)

func Sol004() Solution {
	sol := Solution{
		Problem: "Find the largest palindrome made from the product of two 3-digit numbers.",
		Answer:  answer004,
		Imp:     imp004,
	}

	return sol
}

func imp004() string {
	return "brute"
}

func answer004(imp string) string {
	var ans string

	switch imp {
	case "brute":
		ans = brute004()
	default:
		ans = brute004()
	}

	return ans
}

// Brute force approach - find the largest number possible and decrement until a palindrome is found.
func brute004() string {

	var numOfDigits float64 = 3
	largestNumber := int(math.Pow(math.Pow(10, numOfDigits)-1, 2))
	largestPalindrome := 0

	for num := largestNumber; num > 0; num-- {
		if isPalindrome(num) && hasValidFactors(num, numOfDigits) {
			largestPalindrome = num
			break
		}
	}

	return strconv.Itoa(largestPalindrome)
}

func isPalindrome(num int) bool {

	numStr := strconv.Itoa(num)
	numStrRev := reverseString(numStr)

	isPalindrome := false

	if numStr == numStrRev {
		isPalindrome = true
	}

	return isPalindrome
}

func reverseString(original string) string {

	runes := []rune(original)

	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func hasValidFactors(num int, numOfDigits float64) bool {

	hasValidFactors := false

	largestFactor := int(math.Pow(10, numOfDigits) - 1)
	smallestFactor := int(math.Pow(10, numOfDigits-1))

	for i := largestFactor; i > smallestFactor; i-- {
		if num%i == 0 && len(strconv.Itoa(num/i)) == int(numOfDigits) {
			hasValidFactors = true
		}
	}

	return hasValidFactors
}
