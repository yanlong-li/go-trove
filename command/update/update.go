package update

import (
	"fmt"
	"trove/command/depend"
	"trove/config"
)

func Update(args []string) {
	// 清空锁定文件
	config.TrovePackagesLock = config.TrovePackages{}.Packages
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("Configuration file loading failed")
		return
	}
	if len(args) > 0 {
		newPackageName := args[0]
		if customerPackage, ok := trovePackage.Custom[newPackageName]; ok {
			depend.HandlePackage(newPackageName, customerPackage)
		} else {
			fmt.Println("No package introduced:" + newPackageName)
		}
	} else {
		for newPackageName, customerPackage := range trovePackage.Custom {
			depend.HandlePackage(newPackageName, customerPackage)
		}
	}
	err = config.SaveLock()
	if err != nil {
		fmt.Println("Write Lock Configuration Failed")
	}
}
