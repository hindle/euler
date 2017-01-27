package main

import (
	"fmt"
	"github.com/hindle/euler"
	"github.com/urfave/cli"
	"os"
	"time"
)

var m map[int]euler.Solution

func main() {
	var probNo int
	var imp string
	var list bool

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "problem, p",
			Value:       1,
			Usage:       "Problem to be solved",
			Destination: &probNo,
		},
		cli.StringFlag{
			Name:        "imp, i",
			Value:       "NA",
			Usage:       "Solution to be used",
			Destination: &imp,
		},
		cli.BoolFlag{
			Name:        "list, ls",
			Usage:       "List available implementations",
			Destination: &list,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Running solution for problem %v:\n", probNo)

		prob, ok := getProblem(probNo)

		fmt.Printf("%v\n\n", prob)

		if !ok {
			return nil
		}

		if list {
			listImplementations(probNo)
		} else {
			displayAnswer(probNo, imp)
		}

		return nil
	}

	app.Run(os.Args)
}

func listImplementations(probNo int) {
	imp := "No implementations found!"

	sol, ok := m[probNo]

	if ok {
		imp = sol.Imp()
	}

	fmt.Printf("Available implementations: %v\n", imp)
}

func getProblem(probNo int) (string, bool) {
	prob := "Problem not found!"

	sol, ok := m[probNo]

	if ok {
		prob = sol.Problem
	}

	return prob, ok
}

func displayAnswer(probNo int, imp string) {
	ans := "Answer not defined!"

	sol, ok := m[probNo]

	if ok {
		start := time.Now()
		ans = sol.Answer(imp)
		execTime := time.Since(start)

		fmt.Printf("Answer: %v\n", ans)
		fmt.Printf("Execution time: %v\n", execTime)
	} else {
		fmt.Printf("%v\n", ans)
	}
}

func init() {
	m = make(map[int]euler.Solution)

	m[1] = euler.Sol001()
	m[2] = euler.Sol002()
	m[3] = euler.Sol003()
	m[4] = euler.Sol004()
	m[5] = euler.Sol005()
}
