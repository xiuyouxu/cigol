package status

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// get progress cpu usage by reading /proc/<pid>/stat
func GetProgressCpu(pid int) int {
	p := "/proc/" + strconv.Itoa(pid) + "/stat"
	f, err := os.Open(p)
	if err != nil {
		fmt.Println("failed to open file: " + p)
		return 0
	}
	defer f.Close()

	cpu := 0
	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	if err != nil || err == io.EOF {
		return 0
	}
	i := strings.LastIndex(line, ") ")
	if i >= 0 {
		line = line[i+2:]
		s := strings.Split(line, " ")
		utime, _ := strconv.Atoi(s[11])
		stime, _ := strconv.Atoi(s[12])
		cutime, _ := strconv.Atoi(s[13])
		cstime, _ := strconv.Atoi(s[14])
		cpu = utime + stime + cutime + cstime
	}

	return cpu
}

// get os cpu total usage by reading /proc/stat
func GetOsCpu() int {
	f, err := os.Open("/proc/stat")
	if err != nil {
		fmt.Println("failed to open file:", "/proc/stat")
		return 0
	}
	defer f.Close()

	cpu := 0
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		if strings.HasPrefix(line, "cpu") {
			s := strings.Split(line, " ")
			for _, v := range s[1:] {
				c, _ := strconv.Atoi(v)
				cpu += c
			}
			break
		}
	}
	return cpu
}

// get cpu usage rate for the pid in the interval(unit in seconds)
func GetCpuUsageRate(pid, interval int) float64 {
	t1 := GetOsCpu()
	p1 := GetProgressCpu(pid)
	time.Sleep(time.Duration(interval) * time.Second)
	t2 := GetOsCpu()
	p2 := GetProgressCpu(pid)
	fmt.Println("p2", p2, "\np1", p1, "\nt2", t2, "\nt1", t1)
	if t1 != t2 {
		return 100 * (float64(p2 - p1)) / (float64(t2 - t1))
	}
	return 0
}
