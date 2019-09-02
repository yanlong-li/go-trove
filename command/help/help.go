package Help

var Commands map[string]string
var Options map[string]string

func init() {
	Commands = make(map[string]string)
	Commands["require"] = "引入 Git 包"
	Commands["install [packageName]"] = "安装不存在的包,可选指定包"
	Commands["update [packageName]"] = "更新包 ,可选更新指定包"

	Options = make(map[string]string)
	Options["-h, --help"] = "显示帮助信息"
	Options["--list"] = "显示已引入的包列表"
	Options["-v, --version"] = "显示 Trove 版本号"
}
