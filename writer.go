package ngos

import (
	"encoding/csv"
	"os"
)

// Writer interface
type Writer interface {
	Write([][]string, string) error
}

// CSVWriter struct
type CSVWriter struct {
}

// Write function
func (CSVWriter) Write(datas [][]string, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range datas {
		err = writer.Write(value)
		if err != nil {
			return err
		}
	}

	return nil
}
