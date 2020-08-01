package solver

import (
	"fmt"
	"os"

	"github.com/ArthurKnoep/scrabble-solver/pkg/config"
	"github.com/ArthurKnoep/scrabble-solver/pkg/dico"
	"github.com/ArthurKnoep/scrabble-solver/pkg/heap"
)

func Solver(cfg *config.Config) {
	if len(cfg.Letters) > 7 {
		_, _ = fmt.Fprintln(os.Stderr, "too much letters provided, a maximum of 7 is allowed")
		return
	}
	array := []rune(cfg.Letters)
	d, err := dico.New(cfg.DicoPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not load dictionary: %v\n", err)
		return
	}
	h := newHandler(d)
	heap.Heap(len(array), array, h.handle)
}
