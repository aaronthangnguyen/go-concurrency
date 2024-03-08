package main

import (
	"fmt"
	"os"
	"sync"
)

func catFile(name string) {
	dat, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
}

func main() {
	files := os.Args[1:]
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			catFile(file)
		}(file)
	}

	wg.Wait()
}
