package main

import (
    "os"

    "github.com/JNjenga/netshare/internal/server"
)

func main() {
    cmdArgs := os.Args

    if len(cmdArgs) < 2 {
        panic("Specify repository path")
    }

    repo_path := cmdArgs[1]

    server.Start(repo_path)
}
