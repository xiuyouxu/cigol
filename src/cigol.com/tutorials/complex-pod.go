// complex pod orchestration
// the input of the program is a json file defining the orchestration of
// the complex pod
// steps to create pod from the input json file:
// 1. parse the file
// 2. validate the orchestration definitions from the file
// 3. set default values needed for the pod
// 4. begin to create all the elements in the pod definitions

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"syscall"
)

type Container struct {
	Image      string `json:"image"`
	Memory     int    `json:"memory"`
	Ports      []int  `json:"ports"`
	Network    string `json:"network"`
	Cmd        string `json:"cmd"`
	Entrypoint string `json:"entrypoint"`
}

type ComplexPod struct {
	Name       string      `json:"name"`
	Containers []Container `json:"containers"`
}

func (cp *ComplexPod) String() string {
	b, e := json.Marshal(cp)
	if e != nil {
		panic("failed to marshal type" + reflect.TypeOf(cp).Name())
	}
	return string(b)
}

func (cp *ComplexPod) Exec() bool {
	if cp.Containers != nil && len(cp.Containers) > 0 {
		for _, c := range cp.Containers {
			e := Exec("docker", "run", "-d", "-it", c.Image, c.Cmd)
			if e != nil {
				fmt.Println(e)
				return false
			}
		}
	}
	return true
}

func Exec(args ...string) error {
	if len(args) > 0 {
		cmd := args[0]
		binary, e := exec.LookPath(cmd)
		if e != nil {
			return e
		}
		env := os.Environ()
		e = syscall.Exec(binary, args, env)
		if e != nil {
			return e
		}
	}
	return errors.New("must specify a command to exec")
}

func main() {
	args := os.Args[1:]
	if args == nil {
		fmt.Println("Usage: complex-pod <orchestration-file>")
		return
	}
	b, e := ioutil.ReadFile(args[0])
	if e != nil {
		fmt.Println("failed to read orchestration file", args[0])
		return
	}

	// use nil value object to unmarshal
	//	dat := &ComplexPod{}
	//	fmt.Println(reflect.TypeOf(dat))
	//	if e = json.Unmarshal(b, dat); e != nil {
	//		fmt.Println("failed to unmarshalling orchestration file", args[0], e)
	//		return
	//	}
	//	fmt.Println(dat)

	// use nil object to unmarshal
	var dat2 ComplexPod
//	fmt.Println(reflect.TypeOf(dat2))
	if e = json.Unmarshal(b, &dat2); e != nil {
		fmt.Println("failed to unmarshalling orchestration file", args[0], e)
		return
	}
//	fmt.Println(&dat2)

	fmt.Println("begin to exec the orchestration")
	dat2.Exec()
}
