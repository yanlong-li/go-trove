package install

import (
	"fmt"
	"os"
	"trove/command/require"
	"trove/config"
)

func Install(args []string) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("配置文件加载失败")
		return
	}
	if len(args) > 0 {
		newPackageName := args[0]
		fmt.Println(newPackageName)
		if customerPackage, ok := trovePackage.Custom[newPackageName]; ok {
			_, err := os.Stat("vendor/" + newPackageName)
			if err != nil {
				require.GitClone(customerPackage, newPackageName)
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
				fmt.Println(k)
				require.GitClone(v, k)
			} else {
				fmt.Println(k)
			}
			require.GitVersion(k, v)
			fmt.Println()
		}
	}
}
