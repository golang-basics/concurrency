package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := MD5All(os.Args[1])
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

type MD5Result map[string][md5.Size]byte

func MD5All(root string) (MD5Result, error) {
	done := make(chan struct{})
	defer close(done)

	digestsChan, errChan := sumFiles(done, root)

	m := MD5Result{}
	for r := range digestsChan {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errChan; err != nil {
		return nil, err
	}
	return m, nil
}

// result represents the product of reading and summing a file using MD5.
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// sumFiles returns two channels: one for the results
// and another for the error returned by filepath.Walk.
// The walk function starts a new goroutine
// to process each regular file, then checks done.
// If done is closed, the walk stops immediately
func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	digestsChan := make(chan result)
	errChan := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// check if current path is a directory
			// at least . will be a directory
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				res := result{
					path,
					md5.Sum(data),
					err,
				}
				select {
				case digestsChan <- res:
				case <-done:
				}
				wg.Done()
			}()

			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})

		go func() {
			wg.Wait()
			close(digestsChan)
		}()

		errChan <- err
	}()

	return digestsChan, errChan
}
