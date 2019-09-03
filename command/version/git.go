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
	cmd.Dir = config.VendorPath + packageName
	//fmt.Println(cmd.Dir)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("Version Detection Failed", err)
		return
	}
	if len(buf) == 9 {
		commitId := string(buf[1:8])
		fmt.Println("current version:" + commitId)
	}
}

func GitClone(customerPackage config.CustomerPackage, packageName string) {

	fmt.Println("Restoring:" + packageName)
	source, _ := GitShunt(customerPackage.Source)
	cmd := exec.Command("git", "clone", source, config.VendorPath+packageName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("restore failed", err)
		return
	}
	fmt.Println("Download Successful")
	// 切换分支
	GitCheckoutBranch(customerPackage, packageName)
	if customerPackage.Version != "*" {
		// 切换版本
		GitCheckoutVersion(customerPackage, packageName)
	}

}
func GitUpdate(customerPackage config.CustomerPackage, packageName string) {
	fmt.Println("Updating in progress:" + packageName)
	cmd := exec.Command("git", "pull", "--all")
	cmd.Dir = config.VendorPath + packageName
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Update failed", err)
		return
	}
	fmt.Println("Update Successful")
}

func GitCheckoutBranch(customerPackage config.CustomerPackage, packageName string) {
	// 切换分支
	cmd := exec.Command("git", "checkout", "-b", "_trove")
	cmd.Dir = config.VendorPath + packageName
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Branch handover failure", err)
		return
	}
	//git branch --set-upstream-to=origin/dev
	cmd = exec.Command("git", "branch", "--set-upstream-to=origin/master")
	cmd.Dir = config.VendorPath + packageName
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Branch switching Association failure", err)
		return
	}
	fmt.Println("Successful branch switching")
}

func GitCheckoutVersion(customerPackage config.CustomerPackage, packageName string) {

	if customerPackage.Version == "*" {
		fmt.Println("No need to switch versions")
		return
	}

	_, versionType := GitShunt(customerPackage.Source)
	var cmd *exec.Cmd
	if versionType == "commit" {
		cmd = exec.Command("git", "reset", "--hard", customerPackage.Version)
	} else if versionType == "tag" {
		cmd = exec.Command("git", "reset", customerPackage.Version)
	} else {
		fmt.Println("Unsupported versioning")
		return
	}

	cmd.Dir = config.VendorPath + packageName
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Version switching failed", err)
		return
	}
	fmt.Println("Successful version switching")
}
