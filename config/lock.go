package config

func init() {
	troveLock, err := LoadLock()
	if err != nil {
		troveLock = TrovePackages{}
		// 没有锁定文件需要重新生成
	}
	TrovePackagesLock = troveLock.Packages
}
