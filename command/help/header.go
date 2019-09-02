package Help

import (
	"fmt"
	"time"
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
	fmt.Printf("Trove version %s %s \n", config.Version, time.Now().Format("2006-01-02 15:04:05"))
}
func Command() {
	fmt.Println("Usage:\ncommand [options] [arguments]")
	fmt.Println("Available commands:")
	fmt.Println("Options:")
}
