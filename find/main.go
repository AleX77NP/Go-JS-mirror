package main

import(
	"fmt"
	at "github.com/aleksandarmilanovic/go-js-mirror/arraytype"
)

type arrayT at.ArrayType

type userFunc func(n int) bool

func isBiggerTen(num int) bool {
	if num  > 10 {
		return true
	}
	return false
}

func main() {

	//you can change type in array_type.go and here //
	x := arrayT {
		Elem: []int{2,5,14,6,36}, //[]float32{1.2,35.6,75.4} etc...
	}
	
	fmt.Println(x.find(isBiggerTen))
}

func (a arrayT) find(f userFunc) (int, string) {
	found := 10
	for _, value := range(a.Elem) {
		if f(value) == true {
			found = value
			break
		} 
	}
	if found > 10 {
		return found, ""
	}
	return 0, "not found"
} 