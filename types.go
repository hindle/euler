package euler

type Answer func(imp string) (answer string)
type Imp func() (imp string)

type Solution struct {
	Problem string
	Answer  Answer
	Imp     Imp
}
