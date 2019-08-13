package adaptor

import "fmt"

func Init() {
	roundHole := RoundHole{
		5,
	}

	roundPeg := RoundPeg{
		5,
	}

	if roundHole.isFit(&roundPeg) {
		fmt.Println("Round peg r5 fits round hole r5.")
	}

	smallSqPeg := SquarePeg{
		2,
	}

	largeSqPeg := SquarePeg{
		20,
	}

	smallSqPegAdapter := SquarePegAdaptor{
		smallSqPeg,
	}

	largeSqPegAdapter := SquarePegAdaptor{
		largeSqPeg,
	}

	if roundHole.isFit(&smallSqPegAdapter) {
		fmt.Println("Square peg w2 fits round hole r5.")
	}
	if !roundHole.isFit(&largeSqPegAdapter) {
		fmt.Println("Square peg w20 dos not fit into round hole r5.")
	}

}
