package config

type TrovePackage struct {
	Name        string // 包名称
	name        string
	Version     string          // 版本号
	Description string          // 描述
	Homepage    string          // 网站页面
	Type        string          // 项目类型
	License     string          // 许可协议
	Require     PackageRequire  // Trove管理包
	Custom      CustomerRequire // 自定义包
}
type PackageRequire map[string]string
type CustomerRequire map[string]CustomerPackage
type CustomerPackage struct {
	Version string
	Source  string
	Type    string
}

var TrovePackagePath string = "trove.json"
var Version string = "0.0.1"
