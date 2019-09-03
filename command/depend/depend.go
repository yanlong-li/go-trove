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
		//config.TrovePackages[trovePackage.Name] = config.TroveLockPackage{TrovePackage: trovePackage}
		return
	}
	if troveLockPackage, ok := config.TrovePackages[customerPackageName]; ok {
		// 已处理过，无需再次处理
		troveLockPackage.Use++
		config.TrovePackages[customerPackageName] = troveLockPackage
		return
	}

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

	// 加载指定包的配置信息，如果没有则新建零值配置信息
	// 递归加载所有的依赖
	dependPackage, err := config.Load(config.VendorPath + customerPackageName + "/" + config.TrovePackagePath)
	if err != nil {
		// 将该包写入锁定区域
		config.TrovePackages[customerPackageName] = config.TroveLockPackage{TrovePackage: config.TrovePackage{Name: customerPackageName, Version: customerPackage.Version, Type: config.TypeModule}, Use: 1}
		return
	} else {
		// 将该包写入锁定区域
		config.TrovePackages[customerPackageName] = config.TroveLockPackage{TrovePackage: dependPackage, Use: 1}
		for k, v := range dependPackage.Custom {
			HandlePackage(k, v)
		}
	}

}
