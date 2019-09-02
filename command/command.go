package command

import (
	"fmt"
	Help "trove/command/help"
	"trove/command/install"
	"trove/command/list"
	"trove/command/remove"
	"trove/command/require"
	"trove/command/update"
)

func Shunt(args []string) {
	switch args[0] {
	case "--list":
		fmt.Println("PackageList:")
		_, _ = list.Get()
	case "require":
		require.Require(args[1:])
	case "-h":
		fallthrough
	case "--help":
		Help.Header()
		Help.Version()
		Help.Command()
	case "-v":
		fallthrough
	case "--version":
		Help.Version()
	case "install":
		install.Install(args[1:])
	case "update":
		update.Update(args[1:])
	case "remove":
		remove.Remove(args[1:])
	default:
		fmt.Println("没有发现任何匹配命令", args)
	}
}
