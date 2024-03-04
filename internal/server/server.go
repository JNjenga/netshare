package server

import (
    "log"
    "net"
    "io"
    "encoding/binary"
    "strings"
    // "fmt"
    // "bufio"
    "github.com/JNjenga/netshare/internal/filesystem"
)

const (
    SERVER_HOST = "localhost:8080"
    SERVER_TYPE = "tcp"
)

func Start(repo_path string, cwd *string) {
    log.Println("Starting server at {}", SERVER_HOST);

    ln, err := net.Listen(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    defer ln.Close()

    for {
        conn, err := ln.Accept()

        log.Println("Accepted connection");

        checkErr(err)

        go handleConnection(conn, repo_path, cwd)
    }
}

func handleConnection(conn net.Conn, repo_path string, cwd *string) {
    log.Println("Processing connection...")

    request,_ := readRequest(conn);

    command := strings.Split(request, "\n")

    log.Println("Command:", command)

    switch command[0] {
        case "ls":
            dirEntries, err := filesystem.ListDir(repo_path, *cwd)
            checkErr(err)

            var files string

            for _, dirEntry := range dirEntries {
                if dirEntry.IsDir() {
                    files += "./"
                }
                files += dirEntry.Name() + "\n"
            }

            log.Printf("Files:\n%s", files)

            writeResponse(conn, []byte(files))

            break;
        case "cd":
            if len(command) < 2 {
                log.Println("Error: File path not provided")
                conn.Close()
                return
            }

            log.Printf("Path: %s\n", command[1]);

            path := repo_path + "/" + *cwd + "/" + command[1];

            is_dir, err := filesystem.IsDir(path)

            checkErr(err)

            if is_dir {
                // TODO: Handle multiple kinds of paths
                // relative, absolute etc
                *cwd = command[1]
                writeResponse(conn, []byte(*cwd));
            } else {
                writeResponse(conn, []byte("Path not dir or not found\n"));
            }
        case "cp":
            if len(command) < 2 {
                log.Println("Error: File path not provided")
                conn.Close()
                return
            }

            log.Printf("path: %s\n", repo_path + "/" + *cwd + "/" + command[1])

            // TODO: Does file exist?
            data, err := filesystem.ReadFile(repo_path + "/" + *cwd + "/" + command[1])
            checkErr(err)

            writeResponse(conn, data)

            break
    }

    conn.Close()

    log.Println("Done Writing file")
}

func readRequest(conn net.Conn) (string, error) {
    var request_len_bytes [4] byte
    var request_len int

    _, err := io.ReadAtLeast(conn, request_len_bytes[:], 4);
    if err != nil {
        return "", err
    }

    request_len = int(binary.LittleEndian.Uint32(request_len_bytes[:]))

    // Read request
    request := make([]byte, request_len)
    _, err = io.ReadAtLeast(conn, request, request_len)
    if err != nil {
        return "", err
    }

    if len(request) == 0 {
        return "", nil 
    }

    return string(request), nil
}

func writeResponse(conn net.Conn, data []byte){
    response := make([] byte, 4)
    binary.LittleEndian.PutUint32(response, uint32(len(data)))

    var status_code byte = 0b00000000

    response = append(response, status_code)
    response = append(response, data...)

    _, err := conn.Write(response)
    checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
