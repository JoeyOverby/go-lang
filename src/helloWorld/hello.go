package main

import "fmt"
import "stringutil"

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("Now Backwards\n")
	fmt.Printf(stringutil.Reverse("hello, world\n"))
}
