package system

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// GetCpuPercent 获取CPU利用率
func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

// GetMemPercent 获取内存利用率
func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

// GetDiskPercent 获取磁盘利用率
func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

// GetComputerInfo 获取本机信息
func GetComputerInfo() *host.InfoStat {
	ComputerInfo, _ := host.Info()
	return ComputerInfo
}

// GetNetConnectionInfo 获取当前网络连接信息
func GetNetConnectionInfo() []net.ConnectionStat {
	ConnectInfo, _ := net.Connections("all") // 可填入tcp、udp、tcp4、udp4等等
	return ConnectInfo
}

// GetNetFlowInfo 获取网络读写字节／包的个数
func GetNetFlowInfo() []net.IOCountersStat {
	flowInfo, _ := net.IOCounters(false)
	return flowInfo
}

// GetAllProcessID 获取进程信息
func GetAllProcessID() []int32 {
	processInfo, _ := process.Pids() // 获取当前所有进程的pid
	return processInfo
}
