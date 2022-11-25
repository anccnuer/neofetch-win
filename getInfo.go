package main

import (
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func getCpu() string {
	cpu, _ := cpu.Info()
	return cpu[0].ModelName
}

func getMem() (uint64, uint64) {
	m, _ := mem.VirtualMemory()
	return m.Total / 1024 / 1024, m.Used / 1024 / 1024

}

func getHost() string {
	u, _ := user.Current()
	userName := strings.Split(u.Username, "\\")
	h, _ := host.Info()
	return userName[1] + "@" + h.Hostname
}

func GetOS() (string, string) {
	h, _ := host.Info()
	return h.OS, h.PlatformVersion
	// sys := runtime.GOOS
	// arch := runtime.GOARCH
	// return sys, arch
}

func GetDisk() (uint64, uint64) {
	d1, _ := disk.Usage("C:")
	return d1.Total / 1024 / 1024 / 1024, d1.Used / 1024 / 1024 / 1024
}

func getUptime() (uint64, uint64, uint64, uint64) {
	t, _ := host.Uptime()
	// print(t)
	s := t % 60
	m := t / 60 % 60
	h := t / 60 / 60 % 24
	d := t / 60 / 60 / 24
	return s, m, h, d
}

// func getIp() string {
// 	ip, _ := net.InterfaceAddrs()
// }
