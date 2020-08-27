package main

import (
	"missdeps"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(missdeps.Analyzer) }