package config

type TrovePackage struct {
	Name        string                 // 包名称
	Version     string                 // 版本号
	Description string                 // 描述
	Keywords    []string               //关键词
	Homepage    string                 // 网站页面
	Type        string                 // 项目类型
	License     string                 // 许可协议
	Authors     []Author               // 作者列表
	Support     map[string]interface{} // 技术支持
	Require     PackageRequire         // Trove管理包
	Custom      CustomerRequire        // 自定义包
}
type PackageRequire map[string]string
type CustomerRequire map[string]CustomerPackage
type CustomerPackage struct {
	Version string
	Source  string
	Type    string
}
type Author struct {
	Name        string // 名称
	Email       string // 电子邮箱
	Description string // 描述
}

var TrovePackagePath string = "trove.json"
var Version string = "0.0.1"

var VendorPath string = "vendor"
