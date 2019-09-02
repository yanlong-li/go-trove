package Help

import (
	"fmt"
	"trove/config"
)

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
func Version() {
	fmt.Printf("Trove version %s %s \n", config.Version, "2019-9-2 17:47:55")
}
func Command() {
	fmt.Println("Usage:\ncommand [options] [arguments]")
	fmt.Println("Available commands:")
	for k, v := range Commands {
		fmt.Println(k, "\t", v)
	}
	fmt.Println("Options:")
	for k, v := range Options {
		fmt.Println(k, "\t", v)
	}
}
