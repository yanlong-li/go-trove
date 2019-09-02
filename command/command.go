package command

import (
	"fmt"
	Help "trove/command/help"
	"trove/command/list"
	"trove/command/require"
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
	default:
		fmt.Println("没有发现任何匹配命令", args)
	}
}
