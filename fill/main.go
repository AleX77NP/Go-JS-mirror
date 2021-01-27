package main

import(
	"fmt"
	at "github.com/aleksandarmilanovic/go-js-mirror/arraytype"
)

type arrayT at.ArrayType

func main() {

	//you can change type in array_type.go and here //
	x := arrayT {
		Elem: []int{15,64,13,6,42,2}, //[]float32{1.2,35.6,75.4} etc...
	}

	x.fill(10)

	fmt.Println(x);
}

func(a arrayT) fill(b int) { // b float32
	for index := range a.Elem {
		a.Elem[index] = b;
	}
}