package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/slavsan/gog/internal"
)

func imports() *Command {
	return &Command{
		Name:        "imports",
		Description: "Display imports",
		Run: func(args []string) error {
			var target string
			var module string
			var err error
			var directories map[string]*internal.Directory

			target, err = filepath.Abs(args[0])
			if err != nil {
				panic(err)
			}

			module, err = getModule(target)
			if err != nil {
				panic(err)
			}

			directories, err = internal.LoadPackages(target, module, target)
			if err != nil {
				panic(err)
			}

			for _, directory := range directories {
				internal.ParsePackage(directory, module, target)
			}

			fmt.Printf("%s\n", internal.FormatImports(directories))

			return nil
		},
	}
}