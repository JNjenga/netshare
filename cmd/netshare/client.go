package main

import (
    "os"

    "github.com/JNjenga/netshare/internal/client"
)

func main() {
    data := client.Download()

    os.WriteFile("./downloaded_file.txt", data, 0666)
}
