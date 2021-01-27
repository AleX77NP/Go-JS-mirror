package main 

import(
	"fmt"
	at "github.com/aleksandarmilanovic/go-js-mirror/arraytype"
)

type arrayT at.ArrayType

func main() {

	//you can change type in array_type.go and here //
	x := arrayT {
		Elem: []int{1,2,3,4,5}, //[]float32{1.2,35.6,75.4} etc...
	}

	fmt.Println(x.copyWithin(0,3,4)); // use d = -1 if you want it to go until the last elem
}

func(a arrayT) copyWithin(b, c, d int) (e arrayT) {
	var temp []int
	var r int
	if d == -1 {
		r = len(a.Elem)
	} else {
		r = d
	}
		if r > c {
		for j :=c; j<r; j++ {
			temp = append(temp, a.Elem[j])
		}
		var i int = 0;
		for k:= b; k<b+(r-c); k++ {
			a.Elem[k] = temp[i]
			i++;
		}
	  }
	return a
}