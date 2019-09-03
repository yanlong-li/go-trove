package version

import (
	"fmt"
	"os/exec"
	"strings"
	"trove/config"
)

// GIT 版本控制 分流 目前支持 commit@ tag@
func GitShunt(url string) (source, versionControl string) {
	versionControlPos := strings.Index(url, "@")
	if versionControlPos < 0 || url[:versionControlPos] == "git" {
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
	//fmt.Println(cmd.Dir)
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

func GitClone(customerPackage config.CustomerPackage, packageName string) {

	fmt.Println("正在恢复:" + packageName)
	source, _ := GitShunt(customerPackage.Source)
	cmd := exec.Command("git", "clone", source, "vendor/"+packageName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("恢复失败", err)
		return
	}
	fmt.Println("下载成功")
	// 切换分支
	GitCheckoutBranch(customerPackage, packageName)
	if customerPackage.Version != "*" {
		// 切换版本
		GitCheckoutVersion(customerPackage, packageName)
	}

}
func GitUpdate(customerPackage config.CustomerPackage, packageName string) {
	fmt.Println("正在更新:" + packageName)
	cmd := exec.Command("git", "pull", "--all")
	cmd.Dir = "vendor/" + packageName
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("更新失败", err)
		return
	}
	fmt.Println("更新成功")
}

func GitCheckoutBranch(customerPackage config.CustomerPackage, packageName string) {
	// 切换分支
	cmd := exec.Command("git", "checkout", "-b", "_trove")
	cmd.Dir = config.VendorPath + "/" + packageName
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("分支切换失败", err)
		return
	}
	//git branch --set-upstream-to=origin/dev
	cmd = exec.Command("git", "branch", "--set-upstream-to=origin/master")
	cmd.Dir = config.VendorPath + "/" + packageName
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("分支切换关联失败", err)
		return
	}
	fmt.Println("分支切换成功")
}

func GitCheckoutVersion(customerPackage config.CustomerPackage, packageName string) {

	if customerPackage.Version == "*" {
		fmt.Println("无需切换版本")
		return
	}

	_, versionType := GitShunt(customerPackage.Source)
	var cmd *exec.Cmd
	if versionType == "commit" {
		cmd = exec.Command("git", "reset", "--hard", customerPackage.Version)
	} else if versionType == "tag" {
		cmd = exec.Command("git", "reset", customerPackage.Version)
	} else {
		fmt.Println("不支持的版本控制方式")
		return
	}

	cmd.Dir = config.VendorPath + "/" + packageName
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("版本切换失败", err)
		return
	}
	fmt.Println("版本切换成功")
}
