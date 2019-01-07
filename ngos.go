package ngos

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// Ngos struct
type Ngos struct {
	Reader Reader
	Writer Writer
	Args   *Arguments
}

// New function return *Ngos
func New(args *Arguments) *Ngos {
	csvReader := CSVReader{}
	csvWriter := CSVWriter{}
	return &Ngos{
		Reader: csvReader,
		Writer: csvWriter,
		Args:   args,
	}
}

// Run function, will Run Ngos
func (n *Ngos) Run() {
	oldFile, err := os.Open(n.Args.OldCSVFile)

	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	readerOldCSVFile := csv.NewReader(bufio.NewReader(oldFile))

	// read new csv file
	newFile, err := os.Open(n.Args.NewCSVFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	readerNewCSVFile := csv.NewReader(bufio.NewReader(newFile))

	linesOld, err := n.Reader.Read(readerOldCSVFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	linesNew, err := n.Reader.Read(readerNewCSVFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(linesNew) <= len(linesOld) {
		fmt.Println("new csv file should larger than old csv file")
		os.Exit(1)
	}

	linesOut := n.compare(linesNew, linesOld)

	fmt.Println(linesOut)

	err = n.Writer.Write(linesOut, n.Args.OutputCSVFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (n *Ngos) compare(a, b []string) [][]string {
	for i := len(a) - 1; i >= 0; i-- {
		for _, vD := range b {
			if a[i] == vD {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}

	var result [][]string

	for _, v := range a {
		result = append(result, []string{v})
	}

	return result
}
