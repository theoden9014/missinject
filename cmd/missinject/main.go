package main

import (
	"github.com/theoden9014/miss-inject"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(missinject.Analyzer) }
