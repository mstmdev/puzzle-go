package polymorphism

import (
	"testing"
)

func ExampleTeacher_Say() {
	var teacher Action = &Teacher{}
	teacher.Say("hello")
	// Output: teacher:hello
}

func ExampleStudent_Say() {
	var student Action = &Student{}
	student.Say("hello")
	// Output: student:hello
}

func ExampleTeacher_Hello() {
	teacher := Teacher{}
	teacher.Action = &teacher
	teacher.Hello()
	// Output: teacher:hello world
}

func ExampleStudent_Hello() {
	student := Student{}
	student.Action = &student
	student.Hello()
	// Output: student:hello world
}

func TestUnimplementedHuman_Hello(t *testing.T) {
	human := &UnimplementedHuman{}
	human.Action = human

	// TODO how to check the Say function is available before calling it?

	// panic
	//human.Hello()
}
