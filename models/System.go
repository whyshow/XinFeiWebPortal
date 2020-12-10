package models

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"math"
	"runtime"
	"strconv"
)

/**
 * 系统信息显示数据模型
 */
type SystemStatus struct {
	DiskAll              string //全部磁盘大小
	DiskUsed             string //磁盘已使用量
	DiskFree             string //磁盘剩余量
	DiskPercentage       string //磁盘剩余占比
	MemoryAll            string //全部内存量
	MemoryUsed           string //内存使用
	MemoryFree           string //内存剩余
	MemorySelf           string //程序占用内存
	MemoryUsedPercentage string //内存使用百分比
	CpuLoad              string //CPU 负载百分比
	CpuCores             string //CPU核心数
	CpulName             string //CPU信息名等
	GolangVersion        string //Go SDK 版本
	Hostname             string //主机名
	Uptime               string //正常运行时间
	Procs                string //程序
	OS                   string //操作系统
	PlatformVersion      string //系统版本
}

func GetSystemStatus() (system SystemStatus) {
	//磁盘
	/**	disk := DiskStat()
	system.DiskAll = strconv.FormatUint(disk.All, 10) + "." + strconv.FormatUint(disk.Rem, 10)
	system.DiskUsed = strconv.FormatUint(disk.Used, 10)
	system.DiskFree = strconv.FormatUint(disk.Free, 10)
	system.DiskPercentage = strconv.FormatUint(disk.Percentage, 10)**/
	//内存
	memory := MemStat()
	system.MemoryAll = strconv.FormatUint(memory.All, 10) + "." + strconv.FormatUint(memory.Rem, 10)
	system.MemoryFree = strconv.FormatUint(memory.Free, 10)
	system.MemorySelf = strconv.FormatUint(memory.Self/1024, 10) + "." + strconv.FormatUint(memory.Self%1024, 10)
	system.MemoryUsed = strconv.FormatUint(memory.Used, 10)
	system.MemoryUsedPercentage = strconv.Itoa(int(Wrap(memory.UsedPercent, 2) / 100))
	//CPU
	cpu := CpuStat()
	system.CpuCores = strconv.Itoa(int(cpu.CpuCores))
	system.CpuLoad = strconv.Itoa(int(cpu.Load))
	system.CpulName = cpu.CpulName
	system.GolangVersion = runtime.Version()

	info, _ := host.Info()
	system.Hostname = info.Hostname
	system.Uptime = strconv.FormatUint(info.Uptime/60/60, 10) + "." + strconv.FormatUint(info.Uptime/60%60, 10) //小时
	system.Procs = strconv.FormatUint(info.Procs, 10)
	if info.OS == "darwin" {
		system.OS = "MacOS"
	} else {
		system.OS = info.OS
	}
	system.PlatformVersion = info.PlatformVersion
	return
}

type DiskStatus struct {
	All        uint64 `json:"all"`
	Used       uint64 `json:"used"`
	Free       uint64 `json:"free"`
	Rem        uint64 `json:"rem"`
	Percentage uint64 `json:"percentage"`
}

/**
// 磁盘总量、使用量、剩余量
func DiskStat() (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	dir, _ := os.Getwd()
	err := syscall.Statfs(dir, &fs)
	if err != nil {
		return
	}
	if runtime.GOOS == "darwin" {
		disk.All = fs.Blocks * uint64(fs.Bsize) / 1000 / 1000 / 1000
		disk.Free = fs.Bfree * uint64(fs.Bsize) / 1000 / 1000 / 1000
		disk.Used = disk.All - disk.Free
		disk.Percentage = disk.Used * 100 / disk.All
		//取余值
		disk.Rem = fs.Blocks * uint64(fs.Bsize) / 1000 / 1000 % 1000 % 100
	} else {
		disk.All = fs.Blocks * uint64(fs.Bsize) / 1024 / 1024 / 1024
		disk.Free = fs.Bfree * uint64(fs.Bsize) / 1024 / 1024 / 1024
		disk.Used = disk.All - disk.Free
		disk.Percentage = disk.Used * 100 / disk.All
		//取余值
		disk.Rem = fs.Blocks * uint64(fs.Bsize) / 1024 / 1024 % 1024 % 100
	}

	return
}
**/
type MemStatus struct {
	All         uint64  `json:"all"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	Self        uint64  `json:"self"`
	Rem         uint64  `json:"rem"`
	UsedPercent float64 `json:"usedpercent"`
}

//项目运行占用、内存总量、内存使用量、内存剩余量、内存使用百分比
func MemStat() (memstatus MemStatus) {
	v, _ := mem.VirtualMemory()
	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	memstatus.Self = memStat.Alloc / 1024
	memstatus.All = v.Total / 1024 / 1024 / 1024
	memstatus.Rem = v.Total / 1024 / 1024 % 1024
	memstatus.Free = v.Available / 1024 / 1024 / 1024
	memstatus.Used = v.Used / 1024 / 1024 / 1024
	memstatus.UsedPercent = v.UsedPercent
	return
}

type CpuStatus struct {
	Load     int64  // 负载率
	CpuCores int32  //CPU核数
	CpulName string //CPU 名称

}

//CPU信息
func CpuStat() (cpustat CpuStatus) {
	infoSta, _ := cpu.Info()
	for _, info := range infoSta {
		//CPU核数
		cpustat.CpuCores = info.Cores
		//CPU名称信息
		cpustat.CpulName = info.ModelName
	}
	//获取CPU负载值
	loa, _ := load.Avg()
	//计算 CPU平均负载率
	cpustat.Load = Wrap(loa.Load1, 2) / int64(cpustat.CpuCores) / 2

	return
}

//将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}
