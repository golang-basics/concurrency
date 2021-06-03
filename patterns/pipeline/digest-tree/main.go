package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"concurrency/patterns/pipeline/digest-tree/digestion"
)

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	if len(os.Args) < 2 {
		log.Fatal("no root directory provided")
	}
	m, err := digestion.MD5AllBoundedParallelism(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
