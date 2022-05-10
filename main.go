package main

import (
	"log"
	"os/exec"
)

func main() {
    log.Printf("Start tunnel\n")
    var cmd *exec.Cmd
    var err error
    cmd = exec.Command("sh", "-c", "chmod +x ../chisel")
    err = cmd.Run()
    if err != nil {
        log.Printf("run chmod return error %v\n", err)
        return
    }
    cmd = exec.Command("../chisel", "--port", "8000", "--backend", "--socks5", "--reverse")
    err = cmd.Run()
    if err != nil {
        log.Printf("start chisel return error %v\n", err)
    }
    return
}