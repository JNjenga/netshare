package client

import (
    "log"
    "net"
    "io"
    "encoding/binary"
)

const (
    SERVER_HOST = "localhost:8080"
    SERVER_TYPE = "tcp"
)

func Ls() string {
    log.Println("Connecting to server at {0}", SERVER_HOST);

    conn, err := net.Dial(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    // Send command
    request := []byte("ls")
    request_len := make([]byte, 4)
    binary.LittleEndian.PutUint32(request_len, uint32(len(request)))
    
    conn.Write(request_len)
    conn.Write(request)

    checkErr(err)

    response, err := io.ReadAll(conn)

    checkErr(err)

    return string(response)
}

func Cp() [] byte {
    log.Println("Connecting to server at {0}", SERVER_HOST);

    conn, err := net.Dial(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    // reader := io.NewReader(conn)
    b, err := io.ReadAll(conn)

    log.Print(string(b))

    checkErr(err)

    return b
}

func Cd() {
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
