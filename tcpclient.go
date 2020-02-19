package main

import (
  "net"
  "fmt"
  "bufio"
  "os"
)

func main() {
  // connect to this socket
  hostName := "10.100.23.147"
  portNum := "34933"

  service := hostName + ":" + portNum
  addr, err ResolveTCPAddr("tcp", servce)
  conn, err := net.Dial("tcp", addr)

  if err != nil {
    log.Fatal(err)
  }

  log.Printf("Established connection to %s \n", service)
  log.Printf("Remote TCP address : %s \n", conn.RemoteAddr().String())
  log.Printf("Local TCP client address : %s \n", conn.LocalAddr().String())

  defer conn.Close()

  message := []byte("Hello UDP server!\000")

  for {
    // send to socket
    fmt.Fprintf(conn, message)
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}
