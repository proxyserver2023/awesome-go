package practice

import "fmt"

func f() int { fmt.Println("f"); return 1 }
func g() int { fmt.Println("g"); return 2 }
func h() int { fmt.Println("h"); return 3 }

func Run() {
	var (
		a int = f()
		b int = g()
		c int = h()
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
