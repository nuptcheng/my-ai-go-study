package main

import "fmt"

func main() {
	fmt.Printf("%T %v\n", vBool, vBool)
	fmt.Printf("%T %v\n", vInt, vInt)
	fmt.Printf("%T %v\n", vUint, vUint)
	fmt.Printf("%T %v\n", vFloat, vFloat)
	fmt.Printf("%T %v %v\n", vComplex, real(vComplex), imag(vComplex))
	fmt.Printf("%T %q\n", vString, vString)
	fmt.Printf("%T %v\n", vArray, vArray)
	fmt.Printf("%T %v\n", vSlice, vSlice)
	fmt.Printf("%T %v\n", vStruct, vStruct)
	fmt.Printf("%T %v\n", vPointer, *vPointer)
	fmt.Printf("%T %v\n", vFunc, vFunc(2, 3))
	fmt.Printf("%T %v\n", vInterface, vInterface.Greet())
	fmt.Printf("%T %v\n", vMap, vMap)
	vSlice = append(vSlice, 3)
	fmt.Printf("slice len=%d cap=%d %v\n", len(vSlice), cap(vSlice), vSlice)
	vMap["c"] = 3
	val, ok := vMap["a"]
	fmt.Printf("map get a: %d %t\n", val, ok)
	vPointer.Age++
	fmt.Printf("person via pointer: %v\n", *vPointer)
	vChan <- 10
	x := <-vChan
	fmt.Printf("%T %v\n", vChan, x)
}
