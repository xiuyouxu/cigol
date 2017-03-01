package main

import (
	"fmt"
	"os"
	"strconv"

	"cigol.com/mini-paas/node/status"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Usage: progress-mem-test <pid> <cpu-time-interval-in-second>")
		return
	}
	pid, _ := strconv.Atoi(args[1])
	interval, _ := strconv.Atoi(args[2])

	osMem := status.GetOsMem()
	fmt.Println(osMem)

	pMem := status.GetProgressMem(pid)
	fmt.Println(pMem.VmRSS)
	fmt.Println("mem usage:", status.GetMemUsageRate(pid))

	osCpu := status.GetOsCpu()
	fmt.Println(osCpu)

	pCpu := status.GetProgressCpu(pid)
	fmt.Println(pCpu)

	fmt.Println("cpu usage in", interval, "seconds:", status.GetCpuUsageRate(pid, interval))
}
