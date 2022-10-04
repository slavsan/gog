package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/slavsan/gog/internal"
)

func types() *Command {
	command := &Command{
		Name:        "types",
		Description: "Display defined types",
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
			for _, p := range directories {
				internal.ParsePackage(p, module, target)
			}

			fmt.Printf("%s", internal.FormatTypes(directories, module))

			return nil
		},
	}
	return command
}