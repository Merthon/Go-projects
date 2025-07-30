package probe

import (
    "math"
    "syscall"
)

// DiskInfo 表示磁盘使用情况（单位：MB）
type DiskInfo struct {
    TotalMB     float64 // 总容量（MB）
    FreeMB      float64 // 可用容量（MB）
    UsedMB      float64 // 已用容量（MB）
    UsedPercent float64 // 使用率（百分比）
}

// GetDiskInfo 获取根目录磁盘使用情况，返回 MB 单位数据
func GetDiskInfo() (*DiskInfo, error) {
    var stat syscall.Statfs_t
    if err := syscall.Statfs("/", &stat); err != nil {
        return nil, err
    }

    // 原始字节数
    totalBytes := float64(stat.Blocks) * float64(stat.Bsize)
    freeBytes  := float64(stat.Bavail) * float64(stat.Bsize)
    usedBytes  := totalBytes - freeBytes

    // 转换到 MB 并四舍五入到两位小数
    totalMB := roundFloat(totalBytes/1024/1024, 2)
    freeMB  := roundFloat(freeBytes/1024/1024, 2)
    usedMB  := roundFloat(usedBytes/1024/1024, 2)
    usedPct := roundFloat(usedBytes/totalBytes*100, 2)

    return &DiskInfo{
        TotalMB:     totalMB,
        FreeMB:      freeMB,
        UsedMB:      usedMB,
        UsedPercent: usedPct,
    }, nil
}

// roundFloat 将浮点数四舍五入到指定小数位
func roundFloat(val float64, precision int) float64 {
    factor := math.Pow(10, float64(precision))
    return math.Round(val*factor) / factor
}
