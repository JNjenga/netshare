package client

import (
    "log"
    "net"
    "io"
)

const (
    SERVER_HOST = "localhost:8080"
    SERVER_TYPE = "tcp"
)

func Download() [] byte {
    log.Println("Connecting to server at {0}", SERVER_HOST);

    conn, err := net.Dial(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    // reader := io.NewReader(conn)
    b, err := io.ReadAll(conn)

    log.Print(string(b))

    checkErr(err)

    return b
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
