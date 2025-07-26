package probe

import "runtime"

// MemoryInfo 内存使用情况
type MemoryInfo struct {
	Alloc      uint64 // 当前已分配内存（字节）
	TotalAlloc uint64 // 总分配内存
	Sys        uint64 // 系统申请的内存
	NumGC      uint32 // GC 次数
}

// GetMemoryInfo 获取当前内存使用信息
func GetMemoryInfo() *MemoryInfo {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &MemoryInfo{
		Alloc:      m.Alloc,
		TotalAlloc: m.TotalAlloc,
		Sys:        m.Sys,
		NumGC:      m.NumGC,
	}
}
