package test

import (
	"log"
	"trove/config"
)

func Test() {
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		log.Fatal("加载配置文件失败", err)
	}
	config.TrovePackages[trovePackage.Name] = config.TroveLockPackage{TrovePackage: trovePackage, Use: 1}
	Lock()
}
