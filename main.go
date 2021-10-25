package main

import (
	"fmt"
	testlibrary "github.com/Brainsoft-Raxat/onelab-hw2-lib"
)

func main() {
	fmt.Println(testlibrary.FlipString("AbAbAbA"))
	fmt.Println(testlibrary.FlipString("aZ123zA"))

	fmt.Println(testlibrary.FindRoots(1, 5, 6))
	fmt.Println(testlibrary.FindRoots(1, 1, 1))
}