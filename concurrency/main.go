package main

import (
	"fmt"
	"sync"
)

func main() {

	slice := []string{"a", "b", "c", "d", "e", "f"}
	sliceLen := len(slice)

	// Without parallelism
	fmt.Println("Execution without parallelism")
	for i := 0; i < sliceLen; i++ {
		fmt.Printf("i = %d, val = %s\n", i, slice[i])
	}

	// With parallelism
	fmt.Println("Execution with parallelism")

	var wg sync.WaitGroup
	wg.Add(sliceLen)

	for i := 0; i < sliceLen; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("i = %d, val = %s\n", i, slice[i])
		}(i)
	}

	wg.Wait()
}
