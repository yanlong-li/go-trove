package create

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"trove/config"
)

// init 为特殊关键字
func Create() {

	_, err := os.Stat(config.TrovePackagePath)

	if !os.IsNotExist(err) {
		log.Fatal("Configuration already exists")
	}

	userName := getUserName()
	pathName := getRuntimePathName()

	trovePackage := &config.TrovePackage{Keywords: make([]string, 0, 1), Authors: make([]config.Author, 0, 1), Require: map[string]string{}, Custom: map[string]config.CustomerPackage{}, Support: map[string]interface{}{}}
	trovePackage.Name = input(fmt.Sprintf("Package name (<vendor>/<name>) [%s/%s]:", userName, pathName), userName+"/"+pathName)
	trovePackage.Description = input("Description []:", "")
	gitUserName := gitConfig("user.name")
	gitUserEmail := gitConfig("user.email")
	authorString := input(fmt.Sprintf("Author [%s/<%s>], n to skip:", gitUserName, gitUserEmail))
	if (len(authorString) == 0 || authorString != "n") && (gitUserEmail != "" && gitUserName != "") {
		author := &config.Author{Name: gitUserName, Email: gitUserEmail}
		trovePackage.Authors = append(trovePackage.Authors, *author)
	}
	trovePackage.Type = input("Package Type (module, project) []:")
	trovePackage.License = input("License []:")
	ifRequire := strings.ToLower(input("Would you like to define your dependencies (require) interactively [yes]?", "yes"))
	if ifRequire == "yes" || ifRequire == "y" {
		fmt.Println("Not yet supported")
	}

	jsonByte, err := json.Marshal(trovePackage)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	// 格式化Json到缓存区域
	_ = json.Indent(&out, jsonByte, "", "\t")
	fmt.Println(out.String())
	ifConfirm := strings.ToLower(input("Do you confirm generation [yes]", "yes"))
	if ifConfirm == "yes" || ifConfirm == "y" {
		err = config.Save(*trovePackage)
		if err != nil {
			log.Fatal("Save failed", err)
		}
		fmt.Println("Successful Preservation")
	}
}

func input(tips string, defaultValue ...string) string {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	fmt.Printf(tips)
	var Input string
	Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
	Input = strings.Replace(Input, "\n", "", -1)
	Input = strings.Replace(Input, "\r", "", -1)
	if len(Input) == 0 {
		Input = fmt.Sprint(defaultValue)
	}
	return strings.Trim(Input, "[\r\n]")
}

func getUserName() string {
	userHomeDir, err := os.UserHomeDir()
	var userName string
	if err != nil {
		userName = "user"
	}
	switch runtime.GOOS {
	case "windows":
		pos := strings.LastIndex(userHomeDir, "\\")
		userName = strings.ToLower(userHomeDir[pos+1:])
	case "linux":
		pos := strings.LastIndex(userHomeDir, "/")
		userName = strings.ToLower(userHomeDir[pos+1:])
	default:
		userName = "user"
	}
	return userName
}

func getRuntimePathName() string {
	runtimeDir, err := os.Getwd()
	var pathName string
	if err != nil {
		pathName = "name"
	}
	switch runtime.GOOS {
	case "windows":
		pos := strings.LastIndex(runtimeDir, "\\")
		pathName = strings.ToLower(runtimeDir[pos+1:])
	case "linux":
		pos := strings.LastIndex(runtimeDir, "/")
		pathName = strings.ToLower(runtimeDir[pos+1:])
	default:
		pathName = "name"
	}

	return pathName
}

func gitConfig(config string) string {
	cmd := exec.Command("git", "config", config)
	buf, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.Trim(string(buf), "[\r\n]")
}
