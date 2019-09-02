package Help

import (
	"fmt"
	"trove/config"
)

var Commands map[string]string
var Options map[string]string

func init() {
	Commands = make(map[string]string)
	Commands["require"] = "引入 Git 包"
	Commands["install [packageName]"] = "安装不存在的包,可选指定包"
	Commands["update [packageName]"] = "更新包 ,可选更新指定包"
	Commands["remove packageName"] = "移除指定包"

	Options = make(map[string]string)
	Options["-h, --help"] = "显示帮助信息"
	Options["--list"] = "显示已引入的包列表"
	Options["-v, --version"] = "显示 Trove 版本号"
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
