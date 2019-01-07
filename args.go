package ngos

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const (
	// Version,  the version of Ngos
	Version = "0.0.0"
)

// Arguments struct will hold flag and arguments from stdin
type Arguments struct {
	OldCSVFile    string
	NewCSVFile    string
	OutputCSVFile string
	ShowVersion   bool
	Help          func()
}

// ParseArgs function, this function will parse flag and arguments from stdin to Argumetns struct
func ParseArgs() (*Arguments, error) {
	var (
		oldCSVFile   string
		newCSVFile   string
		ouputCSVFile string
		showVersion  bool
	)

	flag.StringVar(&oldCSVFile, "old", "old.csv", "old CSV file")

	flag.StringVar(&newCSVFile, "new", "new.csv", "new CSV file")

	flag.StringVar(&ouputCSVFile, "output", "output_diff.csv", "output CSV file")
	flag.StringVar(&ouputCSVFile, "o", "output_diff.csv", "output CSV file")

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")

	fmt.Println(oldCSVFile)
	fmt.Println(newCSVFile)
	fmt.Println(ouputCSVFile)

	flag.Usage = func() {

		fmt.Fprintln(os.Stderr, "	-old | --old", "old CSV file")
		fmt.Fprintln(os.Stderr, "	-new", "new CSV file")
		fmt.Fprintln(os.Stderr, "	-o | --o", "output CSV file")

	}

	flag.Parse()

	if len(oldCSVFile) < 0 {
		return &Arguments{Help: flag.Usage}, errors.New("	(-old) arg required")
	}

	if len(newCSVFile) < 0 {
		return &Arguments{Help: flag.Usage}, errors.New("	(-new) arg required")
	}

	return &Arguments{
		OldCSVFile:    oldCSVFile,
		NewCSVFile:    newCSVFile,
		OutputCSVFile: ouputCSVFile,
		ShowVersion:   showVersion,
		Help:          flag.Usage,
	}, nil
}
