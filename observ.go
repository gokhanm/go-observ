package main

import (
    "fmt"
    "os/exec"
)

func RunCommand(cmd string)string {
    args := []string{"list", "--outdated", "--format=columns"}

    out, err := exec.Command(cmd, args...).Output()

    if err != nil {
        fmt.Printf("%s command not found. Passing..\n", cmd)
        return ""
    }

    return string(out)
}

func main() {
    pipCommands := []string{"pip2", "pip3"}

    for _, cmd := range pipCommands{
        pipOut := RunCommand(cmd)
        if pipOut != "" {
        fmt.Printf("%s Output:\n%s", cmd, pipOut)
        }
    }
}
