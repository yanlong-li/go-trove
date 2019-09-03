package remove

import (
	"fmt"
	"trove/command/depend"
	"trove/config"
)

func Remove(args []string) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("Configuration file loading failed")
		return
	}
	if len(args) < 1 {
		fmt.Println("Please bring the package name to be removed.")
		fmt.Println("trove remove package/name")
		return
	} else {
		newPackageName := args[0]
		if _, ok := trovePackage.Custom[newPackageName]; ok {
			delete(trovePackage.Custom, newPackageName)
			err = config.Save(trovePackage)
			if err != nil {
				fmt.Println("Packet Removal Failure", err)
				return
			}
			//移除间引用包
			depend.Remove(newPackageName)
			// 更新锁定配置
			err = config.SaveLock()
			if err != nil {
				fmt.Println("Failed to update Lock configuration")
			}
			fmt.Println("Removal success")
		} else {
			fmt.Println("No package introduced:" + newPackageName)
		}
	}

}
