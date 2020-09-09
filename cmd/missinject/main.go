package main

import (
	"github.com/theoden9014/missinject"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(missinject.Analyzer) }
