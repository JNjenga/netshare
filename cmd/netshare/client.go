package main

import (
    "os"
    "fmt"
    "bufio"

    "github.com/JNjenga/netshare/internal/client"
)

func main() {
    should_exit := false
    var command string

    scanner := bufio.NewScanner(os.Stdin)

    for !should_exit {
        fmt.Print(">")

        scanner.Scan()

        command = scanner.Text()

        switch command {
            case "ls":
                response := client.Ls()
                fmt.Println("Response:", response)
            case "cd":
                client.Cd();
            case "cp":
                client.Cp();
            case "exit":
                should_exit = true
        }
    }

    // data := client.Download()

    // os.WriteFile("./downloaded_file.txt", data, 0666)
}
