package cpu

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

//CPU is the structure of data cpu
type CPU struct {
	Pourcent int
}

//getCPUUsage get current cpu mhz in %
func getCPUPourcent() int {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	mhz := int(math.Ceil(percent[0]))
	return mhz
}

//ReturnCPUUsage return CPU mhz in %
func ReturnCPUUsage() *CPU {
	value := getCPUPourcent()
	cpu := &CPU{
		Pourcent: value,
	}
	return cpu
}

func getCPUFrequency() int {
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n", cpuStat)
	return 0
}
