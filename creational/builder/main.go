/*
Package Builder is a creational design pattern that lets you construct complex objects step by step.
The pattern allows you to produce different types and representations of an object using the same construction code.
*/
package builder

import "fmt"

func Init() {
	sender := Sender{}
	jsonMsg, err := sender.BuildMessage(&JSONMessageBuilder{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonMsg.Body))

	xmlMsg, err := sender.BuildMessage(&XMLMessageBuilder{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xmlMsg.Body))
}
