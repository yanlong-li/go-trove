package config

func init() {
	troveLock, err := LoadLock()
	if err != nil || troveLock.Packages == nil {
		TrovePackages = make(map[string]TroveLockPackage)
		// 没有锁定文件需要重新生成
	} else {
		TrovePackages = troveLock.Packages
	}
}
