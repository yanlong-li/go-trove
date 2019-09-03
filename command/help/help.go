package Help

import (
	"fmt"
	"trove/config"
)

var Commands map[string]string
var Options map[string]string

func init() {
	Commands = make(map[string]string)
	Commands["require"] = "require package"
	Commands["install [packageName]"] = "Install packages that do not exist, optionally specify packages"
	Commands["update [packageName]"] = "Update packages, optionally update specified packages"
	Commands["remove packageName"] = "Remove the specified package"

	Options = make(map[string]string)
	Options["-h, --help"] = "display help information"
	Options["--list"] = "Display a list of introduced packages"
	Options["-v, --version"] = "Display Trove version number"
}

func Version() {
	fmt.Printf("Trove version %s %s \n", config.Version, "2019-9-2 17:47:55")
}
func Command() {
	fmt.Println()
	fmt.Println("Usage:\ncommand [options] [arguments]\n ")
	fmt.Println("Available commands:")
	for k, v := range Commands {
		fmt.Println(k, "\t", v)
	}
	fmt.Println()
	fmt.Println("Options:")
	for k, v := range Options {
		fmt.Println(k, "\t", v)
	}
}
func Header() {
	a := `
         _____________________________        ____________                   _______________
        /_____    _______     ______  \      /   ____      \                /  /___________/
             /   /       |   /      \  \    /  /     \   \  \              /  /
            /   /        |  |        |  |  /  /       \   \  \            /  /
           /   /         |   \_____ /   / |  |         |  |\  \          /  /__________
          /   /          |   ________/\  \|  |         |  | \  \        /  ___________/
         /   /           |  |          \  \   \       /  /   \  \      /  /
        /   /            |  |           \  \   \_____/  /     \  \____/  /_________
       /___/             |__|            \_____________/       \________/_________/

        		             Go Treasure Trove Package Manage
`
	fmt.Print(a)
}
