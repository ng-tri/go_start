package intermediate

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Gâu gâu!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meo meo!"
}

func saySomething(a Animal) {
	fmt.Println(a.Speak())
}

func RunInterface() {
	var a Animal = Dog{}
	saySomething(a)

	a = Cat{}
	saySomething(a)
}
