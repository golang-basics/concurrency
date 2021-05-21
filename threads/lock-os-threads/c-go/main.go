package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "work.h"
import "C"
import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	thread := C.struct_Thread{
		id: 1,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("go routine")
	}()
	go func() {
		defer wg.Done()
		C.work(&thread)
	}()

	wg.Wait()
}
