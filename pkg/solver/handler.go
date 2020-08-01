package solver

import (
	"fmt"

	"github.com/ArthurKnoep/scrabble-solver/pkg/dico"
)

type Handler struct {
	dico *dico.Dico
}

func newHandler(dico *dico.Dico) *Handler {
	return &Handler{
		dico: dico,
	}
}

func (h *Handler) handle(permutation []rune) {
	tmpWord := string(permutation)
	if h.dico.Find(tmpWord) {
		fmt.Println(tmpWord)
	}
}
