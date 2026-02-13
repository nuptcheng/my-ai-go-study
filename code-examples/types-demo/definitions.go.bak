package main

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hello, " + p.Name
}

type Greeter interface {
	Greet() string
}

var (
	vBool      = true
	vInt       = 42
	vUint      uint = 7
	vFloat     = 3.14
	vComplex   = complex(1, 2)
	vString    = "go"
	vArray     = [3]int{1, 2, 3}
	vSlice     = []int{1, 2}
	vStruct    = Person{Name: "Alice", Age: 30}
	vPointer   = &vStruct
	vFunc      = func(x, y int) int { return x + y }
	vInterface Greeter = vStruct
	vMap       = map[string]int{"a": 1, "b": 2}
	vChan      = make(chan int, 2)
)
