package depend

import (
	"fmt"
	"os"
	"trove/config"
)

func Remove(customerPackageName string) {
	if customerPackage, ok := config.TrovePackages[customerPackageName]; ok {
		customerPackage.Use--
		// 如果引用计数小于等于0 则移除
		if customerPackage.Use <= 0 {
			// 从锁定列表移除
			delete(config.TrovePackages, customerPackageName)
			// 从磁盘中移除文件
			err := os.RemoveAll(config.VendorPath + customerPackageName)
			if err != nil {
				fmt.Println("Directory Removal Failure", err)
			}
		} else {
			// 如果还有引用 则写入更新
			config.TrovePackages[customerPackageName] = customerPackage
		}
		// 递归处理剩余依赖
		for k, _ := range customerPackage.Custom {
			Remove(k)
		}
	}
}
