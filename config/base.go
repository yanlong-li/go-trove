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
type TroveLockPackage struct {
	TrovePackage     // 原包信息
	Use          int // 包引用计数
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

type TrovePackagesLock struct {
	Packages       map[string]TroveLockPackage // 包列表
	GenerationTime string                      // 生成时间
}

var TrovePackages = map[string]TroveLockPackage{}
var TrovePackagePath = "trove.json"      // 包配置文件
var TrovePackagesLockPath = "trove.lock" // 包配置锁定文件
var Version = "0.0.1.13"                 // 默认版本
var VendorPath = "vendor/"               // 包储存目录
var (
	TypeProject = "project"
	TypeModule  = "module"
)
