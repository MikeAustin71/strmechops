package main

import (
	"fmt"
	"github.com/MikeAustin71/strmechops/apptest/examples"
)

func main() {
	mt := examples.MainTest{}

	err := mt.TextLineSpecStandardLine03()

	if err != nil {
		fmt.Printf("%v", err.Error())
	}

}
