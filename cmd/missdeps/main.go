package main

import (
	"github.com/theoden9014/missdeps"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(missdeps.Analyzer) }
