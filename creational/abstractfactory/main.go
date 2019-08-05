package abstractfactory

func Init() {
	paint(WinFactory{})
	paint(MacFactory{})
}

func paint(factory AbstractFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.paint()
	checkbox.paint()
}
