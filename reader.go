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

	//fmt.Println(lines[0])

	return lines, nil
}
