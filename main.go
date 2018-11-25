package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	persons := []Person{
		{"Alamin", 24},
		{"Jahangir", 34},
		{"Shoeb", 39},
	}

	fmt.Println(persons)
	fmt.Println(Person{"Alamin", 24})
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age > persons[j].Age
	})
	fmt.Println(persons)
}
