package solution

type Answer func() (answer string)

type Solution struct {
	Problem string
	Answer  Answer
}
