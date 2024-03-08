package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func grepDir(word, dir string, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			go grepDir(word, filepath.Join(dir, entry.Name()), wg)
		} else if filepath.Ext(entry.Name()) == ".txt" {
			wg.Add(1)
			go grepFile(word, filepath.Join(dir, entry.Name()), wg)
		}
	}
}

func grepFile(word, file string, wg *sync.WaitGroup) {
	defer wg.Done()
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", file, err)
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
	if len(os.Args) < 3 {
		fmt.Println("Usage: <word> <directory>")
		os.Exit(1)
	}

	word := os.Args[1]
	dir := os.Args[2]

	wg := sync.WaitGroup{}
	wg.Add(1)
	go grepDir(word, dir, &wg)

	wg.Wait()
}
