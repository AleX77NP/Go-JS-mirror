package main

import(
	"fmt"
	at "github.com/aleksandarmilanovic/go-js-mirror/arraytype"
)

type arrayT at.ArrayType

func main() {

	//you can change type in array_type.go and here //
	x := arrayT {
		Elem: []int{11,2,4,6,3}, //[]float32{1.2,35.6,75.4} etc...
	}
	y := arrayT {
		Elem: []int{8,3,2,22},
	}

	fmt.Println(x.concat(y));

}

func (a arrayT) concat(b arrayT) (c arrayT) {
	for _,value := range b.Elem {
		a.Elem = append(a.Elem,value)
	}
	return a;
} 
