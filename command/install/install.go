package install

import (
	"fmt"
	"os"
	"trove/command/version"
	"trove/config"
)

func Install(args []string) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("Configuration file loading failed")
		return
	}
	if len(args) > 0 {
		newPackageName := args[0]
		fmt.Println(newPackageName)
		if customerPackage, ok := trovePackage.Custom[newPackageName]; ok {
			_, err := os.Stat("vendor/" + newPackageName)
			if err != nil {
				version.GitClone(customerPackage, newPackageName)
			}
			version.GitVersion(newPackageName, customerPackage)
			fmt.Println()
		} else {
			fmt.Println("No package introduced:" + newPackageName)
		}
	} else {
		for k, v := range trovePackage.Custom {

			_, err := os.Stat("vendor/" + k)
			if err != nil {
				fmt.Println(k)
				version.GitClone(v, k)
			} else {
				fmt.Println(k)
			}
			version.GitVersion(k, v)
			fmt.Println()
		}
	}
}
