package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Load(filePath string) (TrovePackage, error) {
	var trovePackage TrovePackage
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("配置文件不存在", err)
		return trovePackage, err
	}

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("读取包配置错误", err)
		return trovePackage, err
	}
	defer file.Close()

	err = json.Unmarshal(fileByte, &trovePackage)
	if err != nil {
		log.Fatal("配置转结构体失败", err)
		return trovePackage, err
	}

	return trovePackage, nil
}
