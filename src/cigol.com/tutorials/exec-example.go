package main

import (
	"syscall"
	"os"
	"os/exec"
	"fmt"
)

func main() {
	binary, lookErr:=exec.LookPath("telnet")
	if lookErr != nil {
		fmt.Println(lookErr)
	}

	args:=[]string{"telnet", "10.126.3.161", "22"}

	env:=os.Environ()

	//fmt.Println(env)

	execErr:=syscall.Exec(binary, args, env)

	if execErr != nil {
		fmt.Println(execErr)
	}
}