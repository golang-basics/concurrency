package logging

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

// ReaderConfig represents the configuration to start the log reader
type ReaderConfig struct {
	Directory string
	Query     string
	Limit     int
}

// NewReader creates a new instance of log reader
func NewReader(cfg ReaderConfig) (*LogReader, error) {
	filesInfo, err := ioutil.ReadDir(cfg.Directory)
	if err != nil {
		return nil, err
	}

	info := make([]os.FileInfo, 0, len(filesInfo))
	for _, fi := range filesInfo {
		if fi.IsDir() {
			continue
		}

		info = append(info, fi)
	}
	sort.Slice(filesInfo, func(i, j int) bool {
		return filesInfo[i].ModTime().Sub(filesInfo[j].ModTime()) < 0
	})

	lr := &LogReader{
		cfg:       cfg,
		filesInfo: info,
	}
	return lr, nil
}

// LogReader represents the application log reader type
// responsible for reading logs from a given directory
// that were written in the last N minutes
type LogReader struct {
	cfg       ReaderConfig
	filesInfo []os.FileInfo
}

// Read reads the log files using the given LogReader configuration
// and stores it inside a local bytes buffer to be displayed later
func (r *LogReader) Read(ctx context.Context, w io.Writer) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		return r.read(w)
	}
}

func (r *LogReader) read(w io.Writer) error {
	return nil
}

func (r *LogReader) stream(file io.ReadCloser) chan string {
	out := make(chan string)

	go func() {
		defer func() { _ = file.Close() }()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- scanner.Text()
		}

		close(out)
	}()

	return out
}
