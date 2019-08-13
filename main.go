package main

import (
	"fmt"

	"github.com/brandonyeoxg/go-design-pattern/creational/abstractfactory"
	"github.com/brandonyeoxg/go-design-pattern/creational/builder"
	"github.com/brandonyeoxg/go-design-pattern/creational/factory"
	"github.com/brandonyeoxg/go-design-pattern/creational/prototype"

	"github.com/brandonyeoxg/go-design-pattern/structural/adaptor"
)

func main() {
	prettifyDemo(factory.Init, "Factory")
	prettifyDemo(builder.Init, "Builder")
	prettifyDemo(abstractfactory.Init, "Abstract Factory")
	prettifyDemo(prototype.Init, "Prototype")

	prettifyDemo(adaptor.Init, "Adaptor")
}

func prettifyDemo(fn func(), patternName string) {
	fmt.Printf("\n===\nExecuting %s pattern...\n\n", patternName)

	fn()

	fmt.Printf("\n\n%s pattern completed!\n===\n", patternName)
}
