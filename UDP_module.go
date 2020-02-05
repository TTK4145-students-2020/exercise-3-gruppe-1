package UDPmodule


import (
  "bytes"
  "encoding/gob"
  "fmt"
  "net"
)

const broadcast_addr = "255.255.255.255"
const port = ":20001"
type Packet struct {
  ID, Response string,
  Content []byte
}

func Init(readPort string, writePort string) (<-chan Packet, chan<- Packet) {
  receive := make(chan Packet, 10)
  send := make(chan Packet, 10)
  go listen(receive, readPort)
  go broadcast(send, writePort)
  return receive, send
}

func listen(receive chan Packet, port string) {
  localAddress, _ := net.ResolveUDPAddr("udp", port)
  if err != nil {
    printError("=== ERROR: ResolvingUDPAddr in Listen failed.", err)
  }
  fmt.Print("COMM: Listening to port ")
  fmt.Println(localAddress.Port)
  connection, err := net.ListenUDP("udp", localAddress)
  if err != nil {
    printError("=== ERROR: ListenUDP in Listen failed.", err)
  }
  defer connection.Close()

  var message Packet
  for {
    inputBytes := make([]byte, 4096)
    length, _, _ : connection.ReadFromUDP(inputBytes)
    buffer := bytes.NewBuffer(inputBytes[:length])
    decoder := gob.NewDecoder(buffer)
    decoder.Decode(&message)
    //Filter out all messages not relevant for the system
    if message.ID == ID {
      receive <- message
    }
  }
}


func broadcast(send chan Packet, port string) {
  destinationAddress, _ := net.ResolveUDPAddr("udp", broadcast_addr+port)
  connection, err := net.DialUDP("udp", "localhost", broadcastAddress)
  defer connection.Close()

  var buffer bytes.buffer
  encoder := gob.NewEncoder(&buffer)
  for {
    message := <-send
    encoder.Encode(message)
    connection.Write(buffer.Bytes())
    buffer.Reset()
  }
}
