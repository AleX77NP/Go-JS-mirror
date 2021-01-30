package main

import(
	"fmt"
	at "github.com/aleksandarmilanovic/go-js-mirror/arraytype"
)

type arrayT at.ArrayType

type userFunc func(n int) bool

func isEven(num int) bool {
	if num % 2 == 0 {
		return true
	}
	return false
}

func main() {

	//you can change type in array_type.go and here //
	x := arrayT {
		Elem: []int{12,2,4,6,36}, //[]float32{1.2,35.6,75.4} etc...
	}
	
	fmt.Println(x.every(isEven))

}

func (a arrayT) every(f userFunc) (bool) {
	flag := false
	for _, value := range(a.Elem) {
		if f(value) == true {
			flag = true
		} else {
			flag = false
		}
	}
	return flag
} 