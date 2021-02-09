package cpu

import (
	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

//CPU is the structure of a cpu
type CPU struct {
	Mhz int
}

//GetCPUUsage returns the cpu usage
func (c *CPU) GetCPUUsage() int {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	mhz := int(math.Ceil(percent[0]))
	cpu := &CPU{
		Mhz: mhz,
	}
	return cpu.Mhz
}
