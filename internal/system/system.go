package system

type Info struct {
	OS      string
	Kernel  string
	Hashcat string

	CPU    CPUInfo
	Memory MemoryInfo
	Disk   DiskInfo
	GPU    GPUInfo

	Tools []ToolInfo
}
