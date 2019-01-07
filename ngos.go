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

	lineOldMapper := make(map[string]bool)
	for _, record := range linesOld {
		lineOldMapper[record] = true
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

	linesOut := n.compare(linesNew, lineOldMapper)

	fmt.Println(linesOut)

	err = n.Writer.Write(linesOut, n.Args.OutputCSVFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (n *Ngos) compare(a []string, b map[string]bool) [][]string {
	for i := len(a) - 1; i >= 0; i-- {
		if b[a[i]] {
			a = append(a[:i], a[i+1:]...)
		}
	}

	var result [][]string

	for _, v := range a {
		result = append(result, []string{v})
	}

	return result
}
