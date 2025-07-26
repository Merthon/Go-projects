package probe

import (
	"os"
	"runtime"
)

//SystemInfo 封装服务器基础信息
type SystemInfo struct {
	OS string  //操作系统
	Arch string //架构
	GoVersion string //Go版本
	NumCPU int //CPU核心数
	Hostname string //主机名
	ProcessID int //当前进程
}

// GetSystemInfo 获取系统信息
func GetSystemInfo() (*SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	info := &SystemInfo{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		NumCPU:    runtime.NumCPU(),
		Hostname:  hostname,
		ProcessID: os.Getpid(),
	}

	return info, nil
}
