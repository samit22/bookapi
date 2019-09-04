package main

import (
	"bookapi/bindings"

	"github.com/graniticio/granitic/v2"
)

func main() {
	granitic.StartGranitic(bindings.Components())
}
