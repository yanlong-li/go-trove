package update

import (
	"fmt"
	"os"
	"trove/command/version"
	"trove/config"
)

func Update(args []string) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("Configuration file loading failed")
		return
	}
	if len(args) > 0 {
		newPackageName := args[0]
		if customerPackage, ok := trovePackage.Custom[newPackageName]; ok {
			_, err := os.Stat("vendor/" + newPackageName)
			if err != nil {
				version.GitClone(customerPackage, newPackageName)
			} else {
				version.GitUpdate(customerPackage, newPackageName)
				version.GitCheckoutVersion(customerPackage, newPackageName)
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
				version.GitClone(v, k)
			} else {
				version.GitUpdate(v, k)
				version.GitCheckoutVersion(v, k)
			}
			version.GitVersion(k, v)
			fmt.Println()
		}
	}
}
