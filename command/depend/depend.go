package depend

import (
	"fmt"
	"os"
	"trove/command/version"
	"trove/config"
)

// 递归依赖管理

func HandlePackage(customerPackageName string, customerPackage config.CustomerPackage) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		//fmt.Println("加载配置文件失败")
		return
	}

	if trovePackage.Name == customerPackageName {
		// 当前包无需额外处理
		// 将指定包信息写入锁定表
		config.TrovePackagesLock[trovePackage.Name] = trovePackage
		return
	}
	if _, ok := config.TrovePackagesLock[customerPackageName]; ok {
		// 已处理过，无需再次处理
		return
	}
	// 将该包写入锁定区域
	config.TrovePackagesLock[trovePackage.Name] = trovePackage
	// 判断文件夹是否存在
	_, err = os.Stat(config.VendorPath + customerPackageName)
	// 如果不存在则启动 git 克隆
	if err != nil {
		// 克隆并自动切换分支切换版本
		version.GitClone(customerPackage, customerPackageName)
	} else {
		version.GitUpdate(customerPackage, customerPackageName)
		version.GitCheckoutVersion(customerPackage, customerPackageName)
	}
	// 已存在或克隆后输出版本信息
	version.GitVersion(customerPackageName, customerPackage)
	fmt.Println()
	// 递归加载所有的依赖
	dependPackage, err := config.Load(config.VendorPath + customerPackageName)
	if err != nil {
		//fmt.Println("加载配置文件失败")
		return
	} else {
		for k, v := range dependPackage.Custom {
			HandlePackage(k, v)
		}
	}

}
