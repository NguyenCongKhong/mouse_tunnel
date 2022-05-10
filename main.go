package main

import (
	"log"
	"os/exec"
)

func main() {
    log.Printf("Start tunnel\n")
    var cmd *exec.Cmd
    var err error
    var stdoutStderr []byte
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
	}
    log.Println(string(stdoutStderr))
    // cmd = exec.Command("sh", "-c", "chmod +x ./chisel")
    // err = cmd.Run()
    // if err != nil {
    //     log.Printf("run shell chmod return error %v\n", err)
    //     return
    // }
    cmd = exec.Command("./chisel", "--port", "8000", "--backend", "--socks5", "--reverse")
    err = cmd.Run()
    if err != nil {
        log.Printf("start chisel return error %v\n", err)
    }
    return
}