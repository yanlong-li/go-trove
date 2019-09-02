package install

import (
	"fmt"
	"os"
	"trove/command/require"
	"trove/config"
)

func Install() {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("配置文件加载失败")
		return
	}

	for k, v := range trovePackage.Custom {

		_, err := os.Stat("vendor/" + k)
		if err != nil {
			fmt.Println(err)
			require.GitClone(v, k)
		}
		require.GitVersion(k, v)
	}
}
