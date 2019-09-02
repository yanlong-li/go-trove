package main

import (
	"os"
	Command "trove/command"
)

func main() {
	var args []string = os.Args
	// 如果没有传递任何参数
	if len(args) == 1 {
		args = append(args, "-h")
	}
	Command.Shunt(args[1:])
}
