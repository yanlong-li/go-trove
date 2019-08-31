package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	Help "trove/help"
)

var TrovePackagePath string = "trove.json"

type TrovePackage struct {
	Name        string // 包名称
	name        string
	Version     string          // 版本号
	Description string          // 描述
	Homepage    string          // 网站页面
	Type        string          // 项目类型
	License     string          // 许可协议
	Require     PackageRequire  // Trove管理包
	Custom      CustomerRequire // 自定义包
}
type PackageRequire map[string]string
type CustomerRequire map[string]struct {
	Version string
	Source  string
	Type    string
}

func main() {

	// 如果没有传递任何参数
	if len(os.Args) == 1 {
		Help.Header()
		return
	}

	//检测当前文件夹是否存在管理文档
	file, err := os.Open(TrovePackagePath)
	if err != nil {
		log.Fatal("配置文件不存在", err)
	}

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("读取包配置错误", err)
	}
	file.Close()

	var person TrovePackage
	err = json.Unmarshal(fileByte, &person)
	if err != nil {
		log.Fatal("配置转结构体失败", err)
	}

	// 将结构体转义成JSON
	v1, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	// 格式化Json到缓存区域
	_ = json.Indent(&out, v1, "", "\t")
	//创建文件
	file, err = os.Create(TrovePackagePath)
	if err != nil {
		log.Fatal(err)
	}
	// 写入数据
	_, err = file.WriteString(out.String())
	if err != nil {
		log.Fatal(err)
	}
	// 程序结束

}
