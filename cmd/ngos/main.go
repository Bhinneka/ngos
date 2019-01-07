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
		fmt.Println(err)
		args.Help()
		os.Exit(0)
	}

	if args.ShowVersion {
		fmt.Printf("%s version %s (runtime: %s)\n", os.Args[0], ngos.Version, runtime.Version())
		os.Exit(0)
	}

	ng := ngos.New(args)

	ng.Run()
}
