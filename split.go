package filesplit

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

const (
	fileChunk   = 524288
	fileMaxSize = 52428800
)

// Chunk contains the file chunk data.
type Chunk struct {
	Name    string
	Content []byte
}

// Split splits a file in smaller pieces.
func Split(file string) ([]Chunk, error) {
	var chunks []Chunk

	// open file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// get file size
	fileInfo, _ := f.Stat()
	fileSize := fileInfo.Size()

	// files larger than fileMaxSize should be rejected
	if fileSize > fileMaxSize {
		return nil, fmt.Errorf("Sorry, we don't process files larger than %v bytes", fileMaxSize)
	}

	// calculate total number of parts the file will be chunked into
	chunksNumber := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	// split file
	for i := uint64(0); i < chunksNumber; i++ {
		chunkSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		chunkBuffer := make([]byte, chunkSize)

		f.Read(chunkBuffer)

		fileName := filepath.Base(file) + strconv.FormatUint(i, 10)

		if err != nil {
			return nil, err
		}

		// store chunk in struct
		chunks = append(chunks, Chunk{Name: fileName, Content: chunkBuffer})
	}

	return chunks, nil
}

// SplitFromBytes is the same than Split but accepts a []byte
func SplitFromBytes(file []byte) ([]Chunk, error) {
	var chunks []Chunk

	fileSize := int64(len(file))

	// files larger than fileMaxSize should be rejected
	if fileSize > fileMaxSize {
		return nil, fmt.Errorf("Sorry, we don't process files larger than %v bytes", fileMaxSize)
	}

	// calculate total number of parts the file will be chunked into
	chunksNumber := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	// split file
	for i := uint64(0); i < chunksNumber; i++ {
		chunkSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		chunkBuffer := make([]byte, chunkSize)

		fileIndex := strconv.FormatUint(i, 10)

		// store chunk in struct
		chunks = append(chunks, Chunk{Name: fileIndex, Content: chunkBuffer})
	}

	return chunks, nil
}

// Save saves chunks into files.
func Save(chunks []Chunk, path string) error {
	for _, chunk := range chunks {
		if err := ioutil.WriteFile(path+chunk.Name, chunk.Content, 0644); err != nil {
			return err
		}
	}
	return nil
}

// CreateTestFile creates a file for testing.
func CreateTestFile(file string, size int64) error {
	fd, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("Failed to create %s", file)
	}
	_, err = fd.Seek(int64(size)-1, 0)
	if err != nil {
		return fmt.Errorf("Failed to seek")
	}
	_, err = fd.Write([]byte{0})
	if err != nil {
		return fmt.Errorf("Write failed")
	}
	err = fd.Close()
	if err != nil {
		return fmt.Errorf("Failed to close file")
	}

	return nil
}
