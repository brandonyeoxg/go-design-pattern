package abstractfactory

import "fmt"

type Button interface {
	paint()
}

type MacButton struct{}

func (MacButton) paint() {
	fmt.Println("You created MacOSButton")
}

type WinButton struct{}

func (WinButton) paint() {
	fmt.Println("You created WinButton")
}

type Checkbox interface {
	paint()
}

type MacCheckbox struct{}

func (MacCheckbox) paint() {
	fmt.Println("You created MacOSCheckbox")
}

type WinCheckbox struct{}

func (WinCheckbox) paint() {
	fmt.Println("You created WinCheckbox")
}
