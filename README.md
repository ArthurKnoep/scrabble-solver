# Scrabble Solver

A small algorithm to solver scrabble, built with go.

## Usage

```bash
make build
./scrabble-solver -l ABCDEF

# if you want to use a custom dictionary
./scrabble-solver -l ABCDEF -d ./path/to/dico.txt
```

## Dictionaries

This project is shipped with a small french dictionary (it does not contains every words from the French language). If 
you want to use your own dictionnary you simply have to create txt file (UTF-8) and place one word per line.  
