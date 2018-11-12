package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	college string
	subject string
}

type Employee struct {
	Human
	company string
}

func (h *Human) sayHi() {
	fmt.Printf("%T\n", h)
	fmt.Println(h.name, h.age)
}

func (h *Student) sayHi() {
	fmt.Printf("%T\n", h)
	fmt.Println(h.name, h.age)
}

func (h *Employee) sayHi() {
	fmt.Printf("%T\n", h)
	fmt.Println(h.name, h.age)
}

func main() {
	h := Human{
		name: "alamin-human",
		age:  24,
	}

	h.sayHi()

	s := Student{
		Human: Human{
			name: "Alamin-Student",
			age:  24,
		},
		college: "Dhaka College",
		subject: "Science",
	}
	s.sayHi()

	e := Employee{
		Human: Human{
			name: "Alamin-Employee",
			age:  24,
		},
		company: "Golang Inc",
	}
	e.sayHi()
}
