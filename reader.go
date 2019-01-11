package ngos

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

// Reader interface
type Reader interface {
	Read(string) ([][]string, error)
}

// CSVReader struct
type CSVReader struct {
}

// Read function
func (CSVReader) Read(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(file))

	// If LazyQuotes is true, a quote may appear in an unquoted field and a
	// non-doubled quote may appear in a quoted field.
	// see: https://golang.org/pkg/encoding/csv/#Reader
	reader.LazyQuotes = true

	var lines [][]string

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		lines = append(lines, line)

	}

	return lines, nil
}
