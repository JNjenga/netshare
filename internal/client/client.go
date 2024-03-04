package client

import (
    "log"
    "net"
    "io"
    "strings"
    "errors"
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
    err = writeRequest(conn, "ls")
    checkErr(err)

    response, err := readResponse(conn)
    checkErr(err)

    return response
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

    response, err := readResponse(conn)
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
    err = writeRequest(conn, "cd", path)

    checkErr(err)

    response, err := readResponse(conn)
    checkErr(err)

    response_str := string(response)
    if len(response) > 0 {
        log.Println("Response: ", string(response))
    } else {
        log.Println("Null response")
    }

    return response_str
}

func writeRequest(conn net.Conn, command string, args... string) error {
    // Build request
    var sb strings.Builder
    sb.WriteString(command + "\n")

    for i := 0; i < len(args); i++ {
        sb.WriteString(args[i] + "\n")
    }

    request_string := sb.String()

    request := []byte(request_string)
    request_len := make([]byte, 4)

    binary.LittleEndian.PutUint32(request_len, uint32(len(request)))

    _, err := conn.Write(request_len)
    _, err  = conn.Write(request)

    return err
}

func readResponse(conn net.Conn) (string, error) {
    var response_len_bytes [5] byte
    var response_len int
    var status_code byte

    _, err := io.ReadAtLeast(conn, response_len_bytes[:], 5);
    if err != nil {
        return "", err
    }

    response_len = int(binary.LittleEndian.Uint32(response_len_bytes[:4]))
    status_code  = response_len_bytes[4]

    if status_code != 0x00 {
        log.Println("Non-Zero Status code: ", status_code);
        return "", errors.New("Status code not 0x00");
    }

    // Read request
    response := make([]byte, response_len)
    _, err = io.ReadAtLeast(conn, response, response_len)

    if err != nil {
        return "", err
    }

    if len(response) == 0 {
        return "", nil
    }

    return string(response), nil
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
