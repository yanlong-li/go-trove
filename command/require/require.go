package require

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"trove/config"
)

func Require(args []string) {
	if len(args) < 1 {
		fmt.Println("command fail:")
		fmt.Println("trove require [source]")
		return
	}
	// 首先加载配置文件
	trovePackage, err := config.Load(config.TrovePackagePath)
	if err != nil {
		fmt.Println("加载错误", err)
	}

	source, versionType := Control(args[0])
	//fmt.Println(source, versionControl)
	//return
	sourceUrl, err := url.Parse(source)
	if err != nil {
		fmt.Println(err)
	}
	if sourceUrl == nil {
		fmt.Println("源地址不合法")
		return
	}
	newPackageName := strings.ToLower(sourceUrl.Path[1:])

	// 判断配置文件中是否存在指定的包
	if customerPackage, ok := trovePackage.Custom[newPackageName]; ok {
		fmt.Println("已引入包:", newPackageName)
		_, err := os.Stat("vendor/" + newPackageName)
		if err != nil {
			fmt.Println(newPackageName + " 包未加载到本地")
			GitClone(customerPackage, newPackageName)
		}
		// 检测包版本信息
		GitVersion(newPackageName, customerPackage)

		return
	} else {
		// 如果没有则写入新记录到列表中 默认版本号为 * 默认类型为 git ，git默认版本控制方式为 commit:
		newPackage := &config.CustomerPackage{Version: "*", Type: "git", Source: versionType + "@" + source}
		trovePackage.Custom[newPackageName] = *newPackage
		err = config.Save(trovePackage)
		if err != nil {
			fmt.Println("保存配置文件失败")
		}
		fmt.Println("已保存配置")
		// 恢复包
		GitClone(*newPackage, newPackageName)
	}

}
