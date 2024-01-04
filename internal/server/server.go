package server

import (
    "log"
    "net"
)

const (
    SERVER_HOST = "localhost:8080"
    SERVER_TYPE = "tcp"
)

func Start(dat []byte) {
    log.Println("Starting server at {}", SERVER_HOST);

    ln, err := net.Listen(SERVER_TYPE, SERVER_HOST)

    checkErr(err)

    defer ln.Close()

    for {
        conn, err := ln.Accept()

        log.Println("Accepted connection");

        checkErr(err)

        go handleConnection(conn, dat)
    }
}

func handleConnection(conn net.Conn, dat []byte) {
    log.Println("Writing file...")
    
    _, err := conn.Write(dat)
    checkErr(err)
    conn.Close()

    log.Println("Done Writing file")
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
