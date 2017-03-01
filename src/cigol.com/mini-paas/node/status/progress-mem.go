package status

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"cigol.com/mini-paas/common/utils"
)

type ProgressMem struct {
	Name    string
	State   string
	Pid     int
	PPid    int
	VmPeak  int
	VmSize  int
	VmLck   int
	VmHWM   int
	VmRSS   int
	VmData  int
	VmStk   int
	VmExe   int
	VmLib   int
	VmPTE   int
	VmSwap  int
	Threads int
}

var fields *utils.Set = utils.NewSet("Name:", "State:", "Pid:", "PPid:", "VmPeak:", "VmSize:", "VmLck:", "VmHWM:", "VmRSS:", "VmData:", "VmStk:", "VmExe:", "VmLib:", "VmPTE:", "VmSwap:", "Threads:")

func process(pm *ProgressMem, line string) {
	vs := strings.Split(line, "\t")
	if fields.Contains(vs[0]) {
		tmp := reflect.ValueOf(pm).Elem()
		field := vs[0][:len(vs[0])-1]
		v := tmp.FieldByName(field)
		t := v.Type().Name()
		var vv reflect.Value
		vs[1] = strings.Trim(vs[1], " \n\t")
		switch t {
		case "int":
			if strings.Contains(vs[1], " ") {
				vs[1] = strings.Split(vs[1], " ")[0]
			}
			r, _ := strconv.Atoi(vs[1])
			vv = reflect.ValueOf(r)
		default:
			vv = reflect.ValueOf(vs[1])
		}
		v.Set(vv)
	}
}

// get progress meminfo by reading /proc/<pid>/status
func GetProgressMem(pid int) *ProgressMem {
	p := "/proc/" + strconv.Itoa(pid) + "/status"
	f, err := os.Open(p)
	if err != nil {
		fmt.Println("failed to open file: " + p)
		return nil
	}
	defer f.Close()

	pm := &ProgressMem{}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		process(pm, line)
	}
	return pm
}

// get os meminfo by reading /proc/meminfo
func GetOsMem() int {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println("failed to open file:", "/proc/meminfo")
		return 0
	}
	defer f.Close()

	mem := 0
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		if strings.HasPrefix(line, "MemTotal:") {
			s := strings.Split(line, "\t")
			if len(s) < 2 {
				s = strings.Split(line, " ")
				for _, v := range s[1:] {
					if v != "" {
						mem, _ = strconv.Atoi(v)
						break
					}
				}
			} else {
				s[1] = strings.Trim(s[1], " ")
				i := strings.Index(s[1], " ")
				mem, _ = strconv.Atoi(s[1][:i])
			}

			break
		}
	}
	return mem
}

// get progress memory usage rate
func GetMemUsageRate(pid int) float64 {
	osMem := GetOsMem()
	pMem := GetProgressMem(pid)
	return float64(pMem.VmRSS) / float64(osMem) * 100
}
