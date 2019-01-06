# Awesome Golang

## Tour Of Golang
* Multiple Results

``` go
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```
* Named Return Values

``` go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

* Stacking Defers

``` go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

* Pointers

``` go
primes := [6]int{1,2,3,4,5,6}
var s []int = primes[1:4]
```

* Slices are like references to arrays

``` go
names := [4]string{
    "John", "Paul", "George", "Ringo",
}
fmt.Println(names)

a := names[0:2]
b := names[1:3]
fmt.Println(a, b)

b[0] = "XXX"
fmt.Println(a, b)
fmt.Println(names)

/*
Outputs
-------
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
*/
```
* Slice Literals

``` go
q := []int{2, 3, 5, 7, 11, 13}
fmt.Println(q)

r := []bool{true, false, true, true, false}
fmt.Println(r)

s := []struct{
        i int
        b bool
    }{
            {2, true},
            {1, false},
            {3, true},
            {1, true},
    }
```
First an array is created then builds a slice that references it.
