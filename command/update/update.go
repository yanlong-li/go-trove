package update

import (
	"fmt"
	"os"
	"trove/command/require"
	"trove/config"
)

func Update(args []string) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("配置文件加载失败")
		return
	}
	if len(args) > 0 {
		newPackageName := args[0]
		if customerPackage, ok := trovePackage.Custom[newPackageName]; ok {
			_, err := os.Stat("vendor/" + newPackageName)
			if err != nil {
				require.GitClone(customerPackage, newPackageName)
			} else {
				require.GitUpdate(customerPackage, newPackageName)
			}
			require.GitVersion(newPackageName, customerPackage)
			fmt.Println()
		} else {
			fmt.Println("未引入包:" + newPackageName)
		}
	} else {
		for k, v := range trovePackage.Custom {

			_, err := os.Stat("vendor/" + k)
			if err != nil {
				require.GitClone(v, k)
			} else {
				require.GitUpdate(v, k)
			}
			require.GitVersion(k, v)
			fmt.Println()
		}
	}
}
