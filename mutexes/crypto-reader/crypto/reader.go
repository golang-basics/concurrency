package crypto

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"time"
)

// TransactionsReaderConfig represents the configuration to start the crypto transactions reader
type TransactionsReaderConfig struct {
	Address   string
	Interval  time.Duration
	Directory string
	Limit     int
}

// NewTransactionsReader creates a new instance of log reader
func NewTransactionsReader(cfg TransactionsReaderConfig) (*TransactionsReader, error) {
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

	reader := &TransactionsReader{
		cfg:       cfg,
		filesInfo: info,
	}
	return reader, nil
}

// TransactionsReader represents the crypto transactions reader type
// responsible for reading crypto transactions within a given interval for a given address
type TransactionsReader struct {
	cfg       TransactionsReaderConfig
	filesInfo []os.FileInfo
	nowFunc   func() time.Time
}

// Read reads and streams the crypto transactions to an io.Writer using the given config
func (r *TransactionsReader) Read(ctx context.Context, w io.Writer) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		return r.read(w)
	}
}

// if there are an infinite number of log files,
// knowing the exact log rotation period may help
// skip iterations up to the very close of the log file
func (r *TransactionsReader) read(w io.Writer) error {
	logFileIndex := -1
	for i, fi := range r.filesInfo {
		nowMinusT := r.nowFunc().Add(-r.cfg.Interval * time.Minute)
		if nowMinusT.Sub(fi.ModTime()) <= 0 {
			logFileIndex = i
			break
		}
	}
	if logFileIndex == -1 {
		return nil
	}

	filePath := path.Join(r.cfg.Directory, r.filesInfo[logFileIndex].Name())
	f, err := os.Open(filePath)
	defer func() { _ = f.Close() }()
	if err != nil {
		return err
	}

	nowMinusT := r.nowFunc().Add(-r.cfg.Interval * time.Minute)
	file := NewFile(f)
	offset, err := file.IndexTime(nowMinusT)
	if err != nil {
		return err
	}

	others := r.filesInfo[logFileIndex+1 : len(r.filesInfo)]
	readTheRest := func() error {
		for _, fi := range others {
			file, err := os.Open(path.Join(r.cfg.Directory, fi.Name()))
			if err != nil {
				return err
			}

			for line := range r.stream(file) {

				_, err := fmt.Fprintln(w, line)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	if offset < 0 {
		if logFileIndex+1 >= len(r.filesInfo) {
			return nil
		}

		nowMinusT := r.nowFunc().Add(-r.cfg.Interval * time.Minute)
		fi := r.filesInfo[logFileIndex+1]
		if nowMinusT.Sub(fi.ModTime()) > 0 {
			return nil
		}
		return readTheRest()
	}

	_, err = f.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(w)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}

	return readTheRest()
}

func (r *TransactionsReader) stream(file io.ReadCloser) chan string {
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
