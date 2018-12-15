package v7

import "fmt"

type author struct {
	firstName, lastName string
	age                 int
}

func (a *author) getFullNameWithAge() string {
	return fmt.Sprintf("[%d] - %s %s", a.age, a.firstName, a.lastName)
}

type book struct {
	title string
	author
}

func (b *book) getFullNameWithAge() string {
	return fmt.Sprintf("[%s] - %s %s - [%d]", b.title, b.firstName, b.lastName, b.age)
}

func (b *book) bookInfo() {
	fmt.Println(b.title)
	fmt.Println(b.getFullNameWithAge())
}

func Run() {
	author1 := &author{
		"safdsafhao",
		"asaahasfdafaontasfdafay",
		10,
	}
	fmt.Println(author1.getFullNameWithAge())
	book1 := book{
		"Gafafo is awesome",
		*author1,
	}
	book1.bookInfo()
}
