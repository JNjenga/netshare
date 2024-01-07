package server

import (
    "log"
    "net"
    "io"
    "encoding/binary"
    // "bufio"
)

const (
    SERVER_HOST = "localhost:8080"
    SERVER_TYPE = "tcp"
)

func Start() {
    log.Println("Starting server at {}", SERVER_HOST);

    ln, err := net.Listen(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    defer ln.Close()

    for {
        conn, err := ln.Accept()

        log.Println("Accepted connection");

        checkErr(err)

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    log.Println("Processing connection...")

    var command string
    var command_len_bytes [4]byte
    var command_len int
    
    // Read command size
    _, err := io.ReadAtLeast(conn, command_len_bytes[:], 4)
    checkErr(err)

    command_len = int(binary.LittleEndian.Uint32(command_len_bytes[:]))

    // Read command
    command_buffer := make([]byte, command_len)
    _, err = io.ReadAtLeast(conn, command_buffer, command_len)
    checkErr(err)

    command = string(command_buffer)

    log.Println("Command:", command)

    switch command {
        case "ls":
            files := [] string {
                "statement.pdf",
                "img001.png",
                "readme.txt",
            }

            response := make([]byte, 4)
            binary.LittleEndian.PutUint32(response, uint32(len(files)))
            // response = append(response, '\n')

            for i := 0; i < len(files); i++ {
                bvalue := []byte(files[i])
                response = append(response, bvalue...)
                response = append(response, '\n')
            }
            _, err = conn.Write(response)
            log.Println("test")
            break;
        // case "cd":
        // case "cp":
    }

    conn.Close()

    log.Println("Done Writing file")
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
