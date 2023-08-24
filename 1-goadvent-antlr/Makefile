.PHONY: all examples

all: examples

genParser:
	antlr4-go -listener -visitor parser/Calc.g4

examples: genParser
	go build -o calc1 example1/example1.go
	go build -o calc2 example2/example2.go
	go build -o calc3 example3/example3.go

test:
	go test ./...
