package digestion

import (
	"crypto/md5"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// MD5AllBoundedParallelism reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.  In that case,
// MD5All does not wait for inflight read operations to complete.
func MD5AllBoundedParallelism(root string) (MD5Result, error) {
	done := make(chan struct{})
	defer close(done)

	pathsChan, errChan := walkFiles(done, root)

	// start a fixed number of goroutines
	// to read and digest files
	digestsChan := make(chan result)
	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, pathsChan, digestsChan)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(digestsChan)
	}()

	m := MD5Result{}
	for r := range digestsChan {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	// check if Walk failed
	if err := <-errChan; err != nil {
		return nil, err
	}
	return m, nil
}

// walkFiles starts a goroutine to walk the directory tree at root and sends the
// path of each regular file on the string channel. It sends the result of the
// walk on the error channel. If done is closed, walkFiles abandons its work.
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	pathsChan := make(chan string)
	errChan := make(chan error, 1)
	go func() {
		defer close(pathsChan)
		errChan <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// check if current path is a directory
			// at least . will be a directory
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case pathsChan <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return pathsChan, errChan
}

// digester reads path names from paths and sends digests of the corresponding
// files on resultChan until either paths or done is closed.
func digester(done <-chan struct{}, paths <-chan string, digestsChan chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		res := result{
			path,
			md5.Sum(data),
			err,
		}
		select {
		case digestsChan <- res:
		case <-done:
			return
		}
	}
}
