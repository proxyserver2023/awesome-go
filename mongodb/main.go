package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

/*
1. type declaraction
2. new id
3. id in string
4. string -> id
5. is valid id
6. marshal json
7. unmarshal json
8. GetBSON
9. SetBSON
*/

// 1. type declaration
type ID bson.ObjectId

// 2. NewID() create a new ID
func NewID() ID {
	return ID(bson.NewObjectId())
}

// 3. ObjectID to String
func (i ID) String() string {
	return bson.ObjectId(i).Hex()
}

// 4. string -> id
func StringToID(s string) ID {
	return ID(bson.ObjectIdHex(s))
}

// 5. is valid ID
func IsValidID(s string) bool {
	return bson.IsObjectIdHex(s)
}

func main() {
	a := NewID()
	fmt.Printf("[%T] -> %v\n", a, a)

	s := fmt.Sprint(a)
	fmt.Printf("[%T] -> %v\n", s, s)

	if valid := IsValidID(s); valid {
		fmt.Println("Valid")
	}

	b := StringToID(s)
	fmt.Printf("[%T] -> %v\n", b, b)
}
