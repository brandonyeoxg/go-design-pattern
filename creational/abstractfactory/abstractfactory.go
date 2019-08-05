package abstractfactory

type AbstractFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WinFactory struct{}

func (WinFactory) CreateButton() Button {
	return new(WinButton)
}

func (WinFactory) CreateCheckbox() Checkbox {
	return new(WinCheckbox)
}

type MacFactory struct{}

func (MacFactory) CreateButton() Button {
	return new(MacButton)
}

func (MacFactory) CreateCheckbox() Checkbox {
	return new(MacButton)
}
