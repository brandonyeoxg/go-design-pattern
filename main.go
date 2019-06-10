package main

import (
	"fmt"

	"github.com/brandonyeoxg/go-design-pattern/creational/builder"
	"github.com/brandonyeoxg/go-design-pattern/creational/factory"
)

func main() {
	prettifyDemo(factory.Init, "Factory")
	prettifyDemo(builder.Init, "Builder")
}

func prettifyDemo(fn func(), patternName string) {
	fmt.Printf("\nExecuting %s pattern...\n\n", patternName)

	fn()

	fmt.Printf("\n\n%s pattern completed!\n", patternName)
}
