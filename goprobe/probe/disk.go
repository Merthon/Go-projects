package probe

import (
	"syscall"
)

// DiskInfo 表示磁盘使用情况
type DiskInfo struct {
	Total uint64  // 总容量（字节）
	Free  uint64  // 可用容量（字节）
	Used  uint64  // 已用容量（字节）
	Usage float64 // 使用率（百分比）
}

// GetDiskInfo 获取根目录磁盘使用情况
func GetDiskInfo() (*DiskInfo, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		return nil, err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	used := total - free
	usage := float64(used) / float64(total) * 100

	return &DiskInfo{
		Total: total,
		Free:  free,
		Used:  used,
		Usage: usage,
	}, nil
}
