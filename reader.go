package ngos

import (
	"encoding/csv"
	"io"
)

// Reader interface
type Reader interface {
	Read(reader *csv.Reader) ([][]string, error)
}

// CSVReader struct
type CSVReader struct {
}

// Read function
func (CSVReader) Read(reader *csv.Reader) ([][]string, error) {
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
