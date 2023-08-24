.PHONY:all listen visitor

all: listen visitor

genParser:
	antlr4 -Dlanguage=Go -listener -visitor parser/Calc.g4

listen: genParser
	go build -o calc_listen listen/main.go

visitor: genParser
	go build -o calc_visit visitor/main.go
