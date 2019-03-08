package main

func main() {

}

type Cond struct {
	L Locker
}

func NewCond(l Locker) *Cond {}
func (c *Cond) Broadcast()   {}
func (c *Cond) Signal()      {}

type Locker interface {
	Lock()
	Unlock()
}
