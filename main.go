package main

import (
	gy "github.com/graniticio/granitic-yaml/v2"
)
import "bookapi/bindings"

func main() {
	gy.StartGraniticWithYaml(bindings.Components())
}
