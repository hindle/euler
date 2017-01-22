package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"github.com/hindle/euler/solution"
	"time"
)

func main() {
	var problem int

	app := cli.NewApp()

	app.Flags =[]cli.Flag {
		cli.IntFlag {
			Name:        "problem, p",
			Value:       1,
			Usage:       "problem solution to be run",
			Destination: &problem,
		},
	}

	app.Action = func(c *cli.Context) error {

		fmt.Printf("Running solution for problem %v:\n", problem)
		fmt.Printf("%v\n\n", getProblem(problem))

		start := time.Now()
		solution := getSolution(problem)
		execTime := time.Since(start)

		fmt.Printf("Solution: %v\n", solution)
		fmt.Printf("Execution time: %v\n", execTime)
		return nil
	}

	app.Run(os.Args)
}

func getProblem(problemNo int) string {

	switch problemNo {
	case 1:
		return solution.GetProblem001()
	case 2:
		return solution.GetProblem002()
	default:
		return "Problem not found"
	}
}

func getSolution(problemNo int) string {

	switch problemNo {
	case 1:
		return solution.GetSolution001()
	case 2:
		return solution.GetSolution002()
	default:
		return "Solution not found"
	}
}
