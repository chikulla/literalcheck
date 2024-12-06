package main

import (
	"github.com/chikulla/literalcheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(literalcheck.Analyzer)
}
