package config

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"
)

func Save(trovePackage TrovePackage) error {
	// 将结构体转义成JSON
	v1, err := json.Marshal(trovePackage)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var out bytes.Buffer
	// 格式化Json到缓存区域
	_ = json.Indent(&out, v1, "", "\t")
	//创建文件
	file, err := os.Create(TrovePackagePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 写入数据
	_, err = file.WriteString(out.String())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func SaveLock() error {
	var trovePackages TrovePackagesLock
	trovePackages.GenerationTime = time.Now().Format("2006-01-02 15:04:05")
	trovePackages.Packages = TrovePackages
	// 将结构体转义成JSON
	v1, err := json.Marshal(trovePackages)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var out bytes.Buffer
	// 格式化Json到缓存区域
	_ = json.Indent(&out, v1, "", "\t")
	//创建文件
	file, err := os.Create(TrovePackagesLockPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 写入数据
	_, err = file.WriteString(out.String())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
