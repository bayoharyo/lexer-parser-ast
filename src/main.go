package main

import (
	"os"

	"github.com/bayoharyo/lexer-parser-ast/src/lexer"
	"github.com/bayoharyo/lexer-parser-ast/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/04.lang")
	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
