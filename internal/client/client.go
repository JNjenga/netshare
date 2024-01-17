package client

import (
    "log"
    "net"
    "io"
    "encoding/binary"

    "github.com/JNjenga/netshare/internal/filesystem"
)

const (
    SERVER_HOST = "localhost:8080"
    SERVER_TYPE = "tcp"
)

func Ls() string {
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

func Cp(file_name string) {
    log.Println("Connecting to server at {0}", SERVER_HOST);

    conn, err := net.Dial(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    // Send command
    request := []byte("cp " + file_name)
    request_len := make([]byte, 4)
    binary.LittleEndian.PutUint32(request_len, uint32(len(request)))
    
    conn.Write(request_len)
    conn.Write(request)

    checkErr(err)

    response, err := io.ReadAll(conn)
    checkErr(err)

    // Save file

    if len(response) > 0 {
        log.Println("Writing file")
        filesystem.WriteFile(file_name + "_downloaded.txt", response)
    } else {
        log.Println("Null response")
    }
}

func Cd(path string) string {
    log.Println("Sending cd command path: {0}", path)

    conn, err := net.Dial(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    // Send command
    request := []byte("cd " + path)
    request_len := make([]byte, 4)
    binary.LittleEndian.PutUint32(request_len, uint32(len(request)))
    
    conn.Write(request_len)
    conn.Write(request)

    checkErr(err)

    response, err := io.ReadAll(conn)
    checkErr(err)

    response_str := string(response)
    if len(response) > 0 {
        log.Println("Response: ", string(response))
    } else {
        log.Println("Null response")
    }

    return response_str
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
