package example

import (
	"fmt"
)

type Person interface {
	Speak(string)
	GetName() string
}

type Teacher struct {
	Name string
}

func (t *Teacher) Speak(name string) {
	fmt.Printf("Teacher:this is %s\n", t.Name)
}

func (t *Teacher) GetName() (name string) {
	name = t.Name
	return
}

type Student struct {
	Name string
}

func (s *Student) Speak(name string) {
	fmt.Printf("Student:this is %s\n", s.Name)
}

func (s *Student) GetName() (name string) {
	name = s.Name
	return
}

func InterfaceExample() {
	t := Teacher{
		Name: "good",
	}
	s := Student{
		Name: "ggg",
	}
	assertInterface(&t)
	assertInterface(&s)
	/*
		tt := &t
		tt.GetName()
		t.Speak("")
		(&t).Speak("aa")

		fmt.Println(t.Name)
		fmt.Println(tt.Name)

		personMethod(&t)
		personMethod(&s)
	*/
}

func personMethod(p Person) {
	name := p.GetName()
	p.Speak(name)
}

func assertInterface(p Person) {
	_, ok := p.(*Teacher)
	if ok {
		fmt.Printf("Teacher:%T\n", p)
	} else {
		fmt.Printf("not teacher:%T\n", p)
	}

}
