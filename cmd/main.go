package main

import (
	"github.com/chikulla/literalcheck/literalcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(literalcheck.Analyzer)
}
