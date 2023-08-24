.PHONY: all example

all: example

genParser:
	antlr4-go -listener -visitor parser/LabeledExpr.g4

example: genParser
	go build -o calc main.go
