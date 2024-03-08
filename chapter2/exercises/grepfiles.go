package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func grepFile(word, file string) {
	dat, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	hasWord := strings.Contains(string(dat), word)
	fmt.Printf("%s ", file)
	if hasWord {
		fmt.Println("✅")
	} else {
		fmt.Println("❎")
	}
}

func main() {
	word := os.Args[1]
	files := os.Args[2:]
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func(word, file string) {
			defer wg.Done()
			grepFile(word, file)
		}(word, file)
	}

	wg.Wait()
}
