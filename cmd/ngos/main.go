package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
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

	// read ole csv file
	oldFile, err := os.Open(oldCSVFile)

	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	readerOldCSVFile := csv.NewReader(bufio.NewReader(oldFile))

	// read new csv file
	newFile, err := os.Open(newCSVFile)

	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	readerNewCSVFile := csv.NewReader(bufio.NewReader(newFile))

	var (
		linesOld []string
		linesNew []string
	)

	for {
		line, err := readerOldCSVFile.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		linesOld = append(linesOld, line...)

	}

	for {
		line, err := readerNewCSVFile.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		linesNew = append(linesNew, line...)

	}

	if len(linesNew) <= len(linesOld) {
		fmt.Println("new csv file should larger than old csv file")
		os.Exit(1)
	}

	linesOut := compare(linesNew, linesOld)

	fmt.Println(linesOut)

	err = write(linesOut, ouputCSVFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func compare(a, b []string) [][]string {
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

func write(datas [][]string, output string) error {
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
