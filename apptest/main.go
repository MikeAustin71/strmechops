package main

import (
	"fmt"
	"github.com/MikeAustin71/strmechops/apptest/examples"
)

func main() {
	mt := examples.MainTest{}
	numStrRunes := "123.456-"
	leadingNegativeSignChars := []rune{0}
	trailingNegativeSignChars := []rune{'-'}
	decimalSeparatorChars := []rune{'.'}

	err := mt.ExampleExtractNumRunes03(
		[]rune(numStrRunes),
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		"main()")

	if err != nil {
		fmt.Printf("%s\n\n", err.Error())
	}

}
