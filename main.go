package main

import (
	"log"
	"os"
	"os/exec"
)

func runShell(command string){
    cmd := exec.Command("/bin/sh", "-c", command)
    stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
        log.Printf("Run shell %s return error: %v\n", command, err)
        log.Println(string(stdoutStderr))
        return
	}
    log.Printf("Run %s result: %s\n", command, string(stdoutStderr))
}

func main() {
    log.Printf("Start tunnel\n")
    var cmd *exec.Cmd
    var err error
    current_dir, err := os.Getwd()
    if err != nil {
        log.Printf("get cwd error %v\n", err)
        return
    }
    log.Printf("current dir %s\n", current_dir)    
    runShell("chmod +x ./chisel")
    cmd = exec.Command("./chisel", "server", "--port", "8000", "--backend", "--socks5", "--reverse")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    err = cmd.Run()
    if err != nil {
        log.Printf("start chisel return error %v\n", err)
    }
    return
}
