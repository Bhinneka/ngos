package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Bhinneka/ngos"
)

func main() {
	args, err := ngos.ParseArgs()

	if err != nil {
		fmt.Printf("\033[31m%s\033[0m%s", err.Error(), "\n")
		args.Help()
		os.Exit(0)
	}

	if args.ShowVersion {
		fmt.Printf("\033[35m%s version %s (runtime: %s)\033[0m%s", os.Args[0], ngos.Version, runtime.Version(), "\n")
		os.Exit(0)
	}

	ng := ngos.New(args)

	ng.Run()
}
