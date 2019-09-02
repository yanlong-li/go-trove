package require

import (
	"fmt"
	"os/exec"
	"strings"
	"trove/config"
)

// 版本控制

func Control(url string) (source, versionControl string) {
	versionControlPos := strings.Index(url, "@")
	if versionControlPos < 0 {
		versionControl = "commit"
		source = url
	} else {
		versionControl = url[:versionControlPos]
		source = url[versionControlPos+1:]
	}
	return source, versionControl
}

// 获取git代码仓库的最后一次提交ID 短的
func GitVersion(packageName string, customerPackage config.CustomerPackage) {

	cmd := exec.Command("git", "log", "--pretty=format:\"%h\"")
	cmd.Dir = "vendor/" + packageName
	fmt.Println(cmd.Dir)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("版本检测失败", err)
		return
	}
	if len(buf) == 9 {
		commitId := string(buf[1:8])
		fmt.Println("当前版本:" + commitId)
	}
}

func gitClone(customerPackage config.CustomerPackage, packageName string) {

	fmt.Println("正在恢复")
	source, _ := Control(customerPackage.Source)
	cmd := exec.Command("git", "clone", source, "vendor/"+packageName)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("恢复失败", err)
		return
	}
	fmt.Printf("恢复成功: %s\n", buf)
}
