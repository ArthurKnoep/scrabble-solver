package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ArthurKnoep/scrabble-solver/pkg/config"
	"github.com/ArthurKnoep/scrabble-solver/pkg/solver"
)

func main() {
	app := kingpin.New("scrabble-solver", "A quick algorithm to solve scrabble")
	cfg, err := config.ParseFlag(app)
	if err != nil {
		app.Usage(os.Args[1:])
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}
	solver.Solver(cfg)
}
