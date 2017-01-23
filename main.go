package main

import (
	"fmt"
	"github.com/hindle/euler/solution"
	"github.com/urfave/cli"
	"os"
	"time"
)

var m map[int]solution.Solution

func main() {
	var problem int

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "problem, p",
			Value:       1,
			Usage:       "problem solution to be run",
			Destination: &problem,
		},
	}

	app.Action = func(c *cli.Context) error {

		fmt.Printf("Running solution for problem %v:\n", problem)

		prob, ok := getProblem(problem)

		fmt.Printf("%v\n\n", prob)

		if !ok {
			return nil
		}

		start := time.Now()
		ans, ok := getAnswer(problem)
		execTime := time.Since(start)

		if ok {
			fmt.Printf("Answer: %v\n", ans)
			fmt.Printf("Execution time: %v\n", execTime)
		} else {
			fmt.Printf("%v\n", ans)
		}

		return nil
	}

	app.Run(os.Args)
}

func getProblem(probNo int) (string, bool) {
	prob := "Problem not found!"

	sol, ok := m[probNo]

	if ok {
		prob = sol.Problem
	}

	return prob, ok
}

func getAnswer(probNo int) (string, bool) {
	ans := "Answer not defined!"

	sol, ok := m[probNo]

	if ok {
		ans = sol.Answer()
	}

	return ans, ok
}

func init() {
	m = make(map[int]solution.Solution)

	m[1] = solution.Sol001()
	m[2] = solution.Sol002()
	m[3] = solution.Sol003()
}
