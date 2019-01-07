package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		oldCSVFile   string
		newCSVFile   string
		ouputCSVFile string
	)

	flag.StringVar(&oldCSVFile, "old", "old.csv", "old CSV file")

	flag.StringVar(&newCSVFile, "new", "new.csv", "new CSV file")

	flag.StringVar(&ouputCSVFile, "o", "output_diff.csv", "output CSV file")

	flag.Parse()

	fmt.Println(oldCSVFile)
	fmt.Println(newCSVFile)
	fmt.Println(ouputCSVFile)

	flag.Usage = func() {

		fmt.Fprintln(os.Stderr, "	-old | --old", "old CSV file")
		fmt.Fprintln(os.Stderr, "	-new", "new CSV file")
		fmt.Fprintln(os.Stderr, "	-o | --o", "output CSV file")

	}

}
