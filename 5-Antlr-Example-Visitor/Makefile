.PHONY: all

all: build

genParser: parser/Expr.g4
	antlr4-go -no-listener -visitor parser/Expr.g4

build: genParser
	go build
