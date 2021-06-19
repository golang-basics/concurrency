package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"concurrency/patterns/pipeline/digest-tree/digestion"
)

const tmpDir = "tmp"

func TestMain(m *testing.M) {
	createTmpFiles()
	m.Run()
	clearTmpFiles()
}

// to run all the benchmarks cd into "digest-tree" directory and run
// go test -bench=. -benchtime=3s
func BenchmarkMD5AllSimpleDigestion(b *testing.B) {
	b.ReportAllocs()
	_, err := digestion.MD5AllSimple(tmpDir)
	if err != nil {
		b.Fatalf("could not digest tree: %v", err)
	}
}

// to run all the benchmarks cd into "digest-tree" directory and run
// go test -bench=. -benchtime=3s
func BenchmarkMD5AllParallelDigestion(b *testing.B) {
	b.ReportAllocs()
	_, err := digestion.MD5AllParallel(tmpDir)
	if err != nil {
		b.Fatalf("could not digest tree: %v", err)
	}
}

// to run all the benchmarks cd into "digest-tree" directory and run
// go test -bench=. -benchtime=3s
func BenchmarkMD5AllBoundedParallelism(b *testing.B) {
	b.ReportAllocs()
	_, err := digestion.MD5AllBoundedParallelism(tmpDir)
	if err != nil {
		b.Fatalf("could not digest tree: %v", err)
	}
}

func createTmpFiles() {
	err := os.Mkdir(tmpDir, 0755)
	if err != nil {
		log.Fatalf("could not create directory: %v", err)
	}

	createFile := func(fileID string, numberOfRepeats int) {
		file, err := os.Create(tmpDir + "/file" + fileID + ".txt")
		if err != nil {
			log.Fatalf("could not create file: %v", err)
		}

		s := strings.Repeat("file"+fileID, numberOfRepeats)
		_, err = file.Write([]byte(s))
		if err != nil {
			log.Fatalf("could not write to file: %v", err)
		}
	}

	bigFilesCreated := 0
	for i := 0; i < 500; i++ {
		fileID := strconv.Itoa(i + 1)
		if i%2 == 0 && bigFilesCreated < 20 {
			createFile(fileID, 3_000_000)
			bigFilesCreated++
			continue
		}
		createFile(fileID, 10_000)
	}
}

func clearTmpFiles() {
	err := os.RemoveAll(tmpDir)
	if err != nil {
		log.Fatalf("could not remove temporary directory: %v", err)
	}
}
