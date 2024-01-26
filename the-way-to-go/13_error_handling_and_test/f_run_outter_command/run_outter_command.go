package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 1. 通过os.StartProcess启动外部进程
func t1() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}

	fmt.Printf("The process id is %v", pid)
}

// 2. 通过exec.Command 和 cmd.Run 启动外部进程
func t2() {
	cmd := exec.Command("gedit")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}

func main() {
	//t1()
	t2()
}
