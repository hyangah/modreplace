package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Printf("version = %q\n", debug.ModInfo())
}
