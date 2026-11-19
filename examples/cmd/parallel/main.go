package main

import (
	"fmt"
	"strings"
	"sync"
)

func uppercase(titles []string) []string {
	const workers = 2
	result := make([]string, len(titles))
	jobs := make(chan int)
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				result[i] = strings.ToUpper(titles[i])
			}
		}()
	}
	for i := range titles {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return result
}

func main() {
	fmt.Println(uppercase([]string{"Dune", "Solaris", "Kindred"}))
}
