package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    listener, _ := net.Listen("tcp", ":5000")

    fmt.Printf("=> Server is started. Network: %s, Addr: %s\n",
        listener.Addr().Network(),
        listener.Addr().String())

    for {
        // Looking for the client
        conn, error := listener.Accept()

        if error != nil {
            fmt.Println("=> Can't connect!")
            conn.Close()
            continue
        }

        conn.Write([]byte("Hi!!!\n"))

        fmt.Println("=> New client is connected")

        bufReader := bufio.NewReader(conn)
        fmt.Println("=> Reading is started")

        go func(conn net.Conn) {
            defer func() {
                conn.Close()
                fmt.Println("=> Connection is closed")
            }()

            for {
                bytes, error := bufReader.ReadByte()

                if error != nil {
                    fmt.Println("=> Can't read!", error)
                    break
                }

                fmt.Print(string(bytes))
            }
        }(conn)

    }
}