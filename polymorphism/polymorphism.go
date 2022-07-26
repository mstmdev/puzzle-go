package polymorphism

import "fmt"

type Action interface {
	Say(s string)
}

type Human struct {
	Action
}

func (h *Human) Hello() {
	h.Say("hello world")
}

type Student struct {
	Human
}

func (h *Student) Say(s string) {
	fmt.Printf("student:%s\n", s)
}

type Teacher struct {
	Human
}

func (h *Teacher) Say(s string) {
	fmt.Printf("teacher:%s\n", s)
}

type UnimplementedHuman struct {
	Human
}
