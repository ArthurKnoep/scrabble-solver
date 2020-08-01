package dico

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func New(filePath string) (*Dico, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	wordMap := make(map[string]bool)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		normalizedString, _, _ := transform.String(t, scanner.Text())
		if len(normalizedString) <= 7 {
			wordMap[strings.ToUpper(normalizedString)] = true
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Dico{
		words: wordMap,
	}, nil
}
