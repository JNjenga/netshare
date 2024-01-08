package server

import (
    "log"
    "net"
    "io"
    "encoding/binary"
    "strings"
    // "bufio"
    "github.com/JNjenga/netshare/internal/filesystem"
)

const (
    SERVER_HOST = "localhost:8080"
    SERVER_TYPE = "tcp"
)

func Start(repo_path string) {
    log.Println("Starting server at {}", SERVER_HOST);

    ln, err := net.Listen(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    defer ln.Close()

    for {
        conn, err := ln.Accept()

        log.Println("Accepted connection");

        checkErr(err)

        go handleConnection(conn, repo_path)
    }
}

func handleConnection(conn net.Conn, repo_path string) {
    log.Println("Processing connection...")

    var command_len_bytes [4]byte
    var command_len int
    
    _, err := io.ReadAtLeast(conn, command_len_bytes[:], 4)
    checkErr(err)

    command_len = int(binary.LittleEndian.Uint32(command_len_bytes[:]))

    // Read command
    command_buffer := make([]byte, command_len)
    _, err = io.ReadAtLeast(conn, command_buffer, command_len)
    checkErr(err)

    if len(command_buffer) == 0 {
        conn.Close()
        return
    }

    command := strings.Split(string(command_buffer), " ")

    log.Println("Command:", command)

    switch command[0] {
        case "ls":
            dirEntries, err := filesystem.ListDir(repo_path, ".")
            checkErr(err)

            var files string

            for _, dirEntry := range dirEntries {
                files += dirEntry.Name() + "\n"
            }
            log.Printf("Files:\n%s", files)

            writeResponse(conn, []byte(files))

            break;
        // case "cd":
        case "cp":
            if len(command) < 2 {
                log.Println("Error: File path not provided")
                conn.Close()
                return
            }

            log.Printf("path: %s\n", repo_path + "/" + command[1])

            // TODO: Does file exist?
            data, err := filesystem.ReadFile(repo_path + "/" + command[1])
            checkErr(err)

            writeResponse(conn, data)

            break
    }

    conn.Close()

    log.Println("Done Writing file")
}

func writeResponse(conn net.Conn, data []byte){
    response := make([] byte, 4)
    binary.LittleEndian.PutUint32(response, uint32(len(data)))

    response = append(response, data...)

    _, err := conn.Write(response)
    checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
