package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human     // embedded
	specialty string
}

func main() {
	// cmd.Execute()
	mark := Student{
		Human{
			"Mark",
			25,
			120,
		},
		"Computer Science",
	}

	// access fields
	mark.Human = Human{"Marcus", 55, 220}
	mark.Human.age--

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)

	// modify mark's specialty
	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)

	fmt.Println("Mark become old. He is not an athlete anymore")
	mark.age = 46
	mark.weight += 60
	fmt.Println("His age is", mark.age)
	fmt.Println("His weight is", mark.weight)

	alamin := Student{
		Human: Human{
			name:   "alamin",
			age:    24,
			weight: 75,
		},
		specialty: "Software Development",
	}

	fmt.Println(alamin.name)
	fmt.Println(alamin.age)
	fmt.Println(alamin.weight)
	fmt.Println(alamin.specialty)

}
