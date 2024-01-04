package main

import (
    "os"
    "log"

    "github.com/JNjenga/netshare/internal/server"
)

func main() {
    path := os.Args[1]

    dat, err := os.ReadFile(path)
    log.Println(string(dat))

    if err != nil {
        log.Fatal("Couldn't read file {}", path)
        panic(err)
    }

    server.Start(dat)
}
