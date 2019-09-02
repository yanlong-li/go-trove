package list

import (
	"fmt"
	"trove/config"
)

func Get() (config.TrovePackage, error) {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println(err)
		return trovePackage, err
	}
	for name, customerRequire := range trovePackage.Custom {
		fmt.Printf("Name:%s\tVersion:%s\tType:%s\tSource:%s\n", name, customerRequire.Version, customerRequire.Type, customerRequire.Source)
	}

	return trovePackage, nil

}
