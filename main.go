package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	fmt.Println("hello world")
}
