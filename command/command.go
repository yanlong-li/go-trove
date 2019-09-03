package command

import (
	"fmt"
	"trove/command/create"
	Help "trove/command/help"
	"trove/command/list"
	"trove/command/remove"
	"trove/command/require"
	"trove/command/test"
	"trove/command/update"
)

func Shunt(args []string) {
	switch args[0] {
	case "init":
		create.Create()
	case "--list":
		fmt.Println("PackageList:")
		_, _ = list.Get()
	case "--list--all":
		fmt.Println("Not yet supported")
	case "require":
		require.Require(args[1:])
	case "-h":
		fallthrough
	case "--help":
		Help.Header()
		Help.Version()
		Help.Command()
	case "-V":
		fallthrough
	case "--version":
		Help.Version()
	case "update":
		update.Update(args[1:])
	case "remove":
		remove.Remove(args[1:])
	case "test":
		test.Test()
	default:
		fmt.Println("No matching commands were found", args)
	}
}
