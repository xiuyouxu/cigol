package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	fmt.Println("I will kill myself after 10s...")
	for i := 10; i > 0; i-- {
		time.Sleep(time.Second)
		fmt.Println(i, "...")
	}

	binary, lookErr := exec.LookPath("pkill")
	if lookErr != nil {
		fmt.Println(lookErr)
	}

	args := []string{"pkill", "-9", "suicide"}

	env := os.Environ()

	//fmt.Println(env)

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		fmt.Println(execErr)
	}

	fmt.Println("This line should never be printed...")
}
