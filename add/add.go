package add

import "fmt"

func Add(lhs int, rhs int) int {
	fmt.Println("hei d")
	return lhs + rhs<<2
}

type Sayer interface {
	Say() string
}

type Cat struct{}

func (cat *Cat) Say() string {
	// implement me.
	return ""
}
