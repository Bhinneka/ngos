package ngos

import (
	"fmt"
	"os"
	"strings"
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
	// read old csv file
	linesOld, err := n.Reader.Read(n.Args.OldCSVFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create map from old csv file
	lineOldMapper := make(map[string]bool)
	for _, record := range linesOld {
		lineOldMapper[strings.Join(record, ",")] = true
	}

	// read new csv file
	linesNew, err := n.Reader.Read(n.Args.NewCSVFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(linesNew) <= len(linesOld) {
		fmt.Println("new csv file should larger than old csv file")
		os.Exit(1)
	}

	linesOut := n.compare(linesNew, lineOldMapper)

	err = n.Writer.Write(linesOut, n.Args.OutputCSVFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (n *Ngos) compare(a [][]string, b map[string]bool) [][]string {
	for i := len(a) - 1; i >= 0; i-- {
		if b[strings.Join(a[i], ",")] {
			a = append(a[:i], a[i+1:]...)
		}
	}

	var result [][]string

	for _, v := range a {
		fmt.Println(v)
		result = append(result, v)
	}

	return result
}
