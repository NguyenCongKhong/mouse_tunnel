package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
    log.Printf("Start tunnel\n")
    var cmd *exec.Cmd
    var err error
    var stdoutStderr []byte
    current_dir, err := os.Getwd()
    if err != nil {
        log.Printf("get cwd error %v\n", err)
        return
    }
    log.Printf("current dir %s\n", current_dir)
    cmd = exec.Command("/bin/sh", "-c", "GOPATH=/tmp/gotty go get github.com/yudai/gotty")
    stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
        log.Println(string(stdoutStderr))
		log.Printf("run shell ls return error %v\n", err)
        return
	}
    log.Println(string(stdoutStderr))
    cmd = exec.Command("/bin/sh", "-c", "/tmp/gotty/bin/gotty --port 8000 --permit-write --reconnect --credential hello:1 /bin/bash")
    stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
        log.Println(string(stdoutStderr))
		log.Printf("run shell ls return error %v\n", err)
        return
	}
    log.Println(string(stdoutStderr))
    return
    cmd = exec.Command("/bin/sh", "-c", "cp ../chisel .")
    stdoutStderr, err = cmd.CombinedOutput()
    if err != nil {
        log.Println(string(stdoutStderr))
        log.Printf("run shell cp return error %v\n", err)
        return
    }
    
    cmd = exec.Command("/bin/sh", "-c", "ls -al .")
    stdoutStderr, err = cmd.CombinedOutput()
	if err != nil {
        log.Println(string(stdoutStderr))
		log.Printf("run shell ls return error %v\n", err)
        return
	}
    log.Println(string(stdoutStderr))
    chisel := filepath.Join(current_dir, "..", "chisel")
    cmd = exec.Command("sh", "-c", fmt.Sprintf("chmod +x %s", chisel))
    err = cmd.Run()
    if err != nil {
            log.Printf("run shell chmod return error %v\n", err)
        return
    }
   
    cmd = exec.Command(chisel, "--port", "8000", "--backend", "--socks5", "--reverse")
    err = cmd.Run()
    if err != nil {
        log.Printf("start chisel return error %v\n", err)
    }
    return
}