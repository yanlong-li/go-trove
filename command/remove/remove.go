package remove

import (
	"fmt"
	"os"
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
			err = os.RemoveAll(config.VendorPath + newPackageName)
			if err != nil {
				fmt.Println("Directory Removal Failure", err)
				return
			}
			//todo 移除间引用包
			fmt.Println("Removal success")
		} else {
			fmt.Println("No package introduced:" + newPackageName)
		}
	}

}
