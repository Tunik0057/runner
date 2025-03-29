package main

import (
<<<<<<< HEAD
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
=======
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(5)

    for i := 0; i < 5; i++ {
        go func() {
            fmt.Println(i)
            wg.Done()
        }()
    }

    wg.Wait()
>>>>>>> 2e39cc0fe088f7874e06279c463ee8e1936b479a
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

<<<<<<< HEAD
	return b
=======
    return b
>>>>>>> 2e39cc0fe088f7874e06279c463ee8e1936b479a
}
