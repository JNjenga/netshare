package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"

    "github.com/JNjenga/netshare/internal/client"
)

func main() {
    should_exit := false
    var command []string

    scanner := bufio.NewScanner(os.Stdin)

    for !should_exit {
        fmt.Print(">")

        scanner.Scan()

        command = strings.Split(scanner.Text(), " ")

        switch command[0] {
            case "ls":
                response := client.Ls()
                fmt.Println(response)
            case "cd":
                client.Cd();
            case "cp":
                if len(command) < 2 {
                    fmt.Println("Error: Specify file name")
                    break
                }

                client.Cp(command[1]);
            case "exit":
                should_exit = true
        }
    }

    // data := client.Download()

    // os.WriteFile("./downloaded_file.txt", data, 0666)
}
