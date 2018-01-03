package main

import (
	"os"
	"github.com/Nygard-R1/go-gen-graphql-schema/generate"
)

func main() {
	args := os.Args[1:]
	generate.Generate(args)
}
