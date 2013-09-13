package main

import (
    "flag"
    "fmt"
    "net"
    "strconv"
    "time"
)

var (
    ports     = []string{"443", "80"}
    protocols = []string{"tcp"}
    Ip          string
    StartPort   int
    EndPort     int
    OpenPorts   []int
)

func ParseArgs() {
    flag.StringVar(&Ip, "ip", "127.0.0.1", "Input the IP Address you'd like to test")
    flag.IntVar(&StartPort, "start", 1, "Port to start evaluating")
    flag.IntVar(&EndPort, "end", 1000, "Port to end the evaluation")
    flag.Parse()
}

func main() {
    ParseArgs()
    for _, protocol := range protocols {
        for port := StartPort; port <= EndPort; port++ {
            _, err := net.DialTimeout(protocol, Ip+":"+strconv.Itoa(port), time.Millisecond * 100)
            if err == nil {
                OpenPorts = append(OpenPorts, port)
            }
        }
    }
    fmt.Println(OpenPorts)
}
