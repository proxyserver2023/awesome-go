package main

import (
	"fmt"
	"strings"
)

func Compare() {
	fmt.Println("---Compare--")
	a := "Hello2"
	b := "WORLD1"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))
}

func ContainsAny() {
	fmt.Println("---ContainsAny---")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

func Contains() {
	fmt.Println("---Contains---")
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}

func ContainsRune() {
	fmt.Println("--ContainsRune--")
	fmt.Println(strings.ContainsRune("afkasdflkasjdkfa", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))

}

func SplitHavelengthOneCheck() {
	fmt.Println("---Split HaveLenght One For Empty String--")
	test := strings.Split("", ",")
	fmt.Printf("%#v\n", test)
	fmt.Println(len(test))

}

func EqualFold() {
	fmt.Println("-- Unicode Code Folding --")
	fmt.Println(strings.EqualFold("Go", "go"))
}

func HasPrefix() {
	fmt.Println("-- Has Prefix --")
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
}

func HasSuffix() {
	fmt.Println("-- has suffix --")
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
}

func Map() {
	a := "A quick brown fox jumps over the lazy dog."
	f := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	fmt.Println(
		strings.Map(
			f,
			a,
		),
	)
}

func Run() {
	Compare()
	Contains()
	ContainsAny()
	ContainsRune()
	SplitHavelengthOneCheck()
	EqualFold()
	HasPrefix()
	HasSuffix()
	Map()
}
