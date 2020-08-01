run:
	go run cmd/main.go

build:
	go build -o scrabble-solver cmd/main.go

clean:
	rm scrabble-solver
