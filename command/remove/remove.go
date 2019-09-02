package remove

import (
	"fmt"
	"os"
	"trove/config"
)

func Remove(args []string) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("配置文件加载失败")
		return
	}
	if len(args) < 1 {
		fmt.Println("请携带要移除的包名")
		fmt.Println("trove remove package/name")
		return
	} else {
		newPackageName := args[0]
		if _, ok := trovePackage.Custom[newPackageName]; ok {
			delete(trovePackage.Custom, newPackageName)
			err = config.Save(trovePackage)
			if err != nil {
				fmt.Println("包移除失败", err)
				return
			}
			err = os.RemoveAll("vendor/" + newPackageName)
			if err != nil {
				fmt.Println("目录移除失败", err)
				return
			}
			fmt.Println("移除成功")
		} else {
			fmt.Println("未引入包:" + newPackageName)
		}
	}

}
