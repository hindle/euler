package solution

import "strconv"

func GetProblem001() string {

	return "Find the sum of all the multiples of 3 or 5 below 1000."
}

func GetSolution001() string {

  sum := 0
  for i := 0; i < 1000; i++ {
    if i % 3 == 0 {
      sum += i
    } else if i % 5 == 0 {
      sum += i
    }
  }

  return strconv.Itoa(sum)
}
