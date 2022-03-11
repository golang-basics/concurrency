package crypto

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"os"
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

func (r *TransactionsReader) read(w io.Writer) error {
	return nil
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
