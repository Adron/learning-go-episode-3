package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var someInt int
	var anotherValue int64
	thisValue := 45
	thisVAlue, thatValue, someValue := 20, 23, 15
	someInt = thisVAlue
	anotherValue = 64

	const knownSet = 7.12904
	const anotherKnown = 9.10347

	for x := 0; x < 20; x++ {
		fmt.Printf("x = %d  e = %8.3f\n", x, math.Exp(float64(x)))
	}

	fmt.Printf("Known Set and Another Known: %2.7f, %2.7f", knownSet, anotherKnown)

	fmt.Printf("The values: %6d, %6d, %6d, and %s.\n", thisValue, thisVAlue, thatValue, strconv.Itoa(someValue))
	fmt.Printf("Other values: %4d, %4d.\n", someInt, anotherValue)

	fmt.Printf("Function doThings: %6d\n", doThings(thisVAlue, thisValue))
	fmt.Printf("Function doMoreThings: %s", doMoreThings(someInt, int(anotherValue), "Adron"))
}

func doThings (valueOne int, valueTwo int) int64 {
	result := valueOne * valueTwo
	return int64(result)
}

func doMoreThings (valueOne int, valueTwo int, name string) string {
	result := valueTwo + valueOne * 10
	return "The result added and exponentially multiplied for you " + name + " this: " + strconv.Itoa(result)
}