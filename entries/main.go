package main

import(
	"fmt"
	at "github.com/aleksandarmilanovic/go-js-mirror/arraytype"
)

type arrayT at.ArrayType

//Entry ...
type Entry struct {
	key int
	value int
}

//Entries ...
type Entries struct {
	entries []Entry
}

//NextResult ...
type NextResult struct {
	value Entry
}

func main() {

	//you can change type in array_type.go and here //
	x := arrayT {
		Elem: []int{11,2,4,6,3}, //[]float32{1.2,35.6,75.4} etc...
	}

	z := x.entries();

	w := Entries {
		entries: z,
	}

	fmt.Println(w.next().value);
	fmt.Println(w.next().value);
	fmt.Println(w.next().value);
}

var current int = 0

func(a arrayT) entries() ([]Entry) { 
	var temp = []Entry{}
	for index, value := range a.Elem {
		t := Entry {
			key: index,
			value: value,
		}
		temp = append(temp, t)
	}
	return temp
}

func(e Entries) next() (v NextResult) {
	currentValue := NextResult {
		value: e.entries[current],
	}
	current++;
	return currentValue;
}
