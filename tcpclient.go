package main

import (
  "net"
  "fmt"
  "log"
  "time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
  // connect to this socket
  hostName := "10.100.23.147"
  portNum := "33546"

  service := hostName + ":" + portNum
  local_service := "10.100.23.221:9999"
  remote_addr,_ := net.ResolveTCPAddr("tcp", service)

  conn, err := net.DialTCP("tcp", nil ,remote_addr)
  checkError(err)

  log.Printf("Established connection to %s \n", service)
  log.Printf("Remote TCP address : %s \n", conn.RemoteAddr().String())
  log.Printf("Local TCP client address : %s \n", conn.LocalAddr().String())

  messageIP := "Connect to:" + local_service + "\000"

  local_addr, err := net.ResolveTCPAddr("tcp", local_service)
  checkError(err)

  //listen to local port
  listener, err := net.ListenTCP("tcp", local_addr)
  checkError(err)

  //connnection order
  _, err = conn.Write([]byte(messageIP))
  checkError(err)

  client_conn, err := listener.AcceptTCP()
  checkError(err)

  for {
    // write
    msg := "cyka blyat"
    client_conn.Write([]byte(msg))

    // listen
    bugger := make([]byte,1024)
    client_conn.Read(bugger)

    message := string(bugger)
    fmt.Print("Message from server: "+ message)

    time.Sleep(1 * time.Second)
  }
}
