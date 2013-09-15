package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net"
)

var (
    DotcomServer string = "whois.verisign-grs.com"
    Domain       string 
)

func MakeRequest() {
    conn, err := net.Dial("tcp", DotcomServer + ":43")
    if err != nil {
        panic(err)    
    }
    defer conn.Close()
    fmt.Fprintf(conn, "%s \n", flag.Arg(0))  // I guess this is similar to conn.Write
    resp, err := ioutil.ReadAll(conn)
    if err != nil {
        panic(err)    
    }
    fmt.Printf("%s \n", resp)
}

func help() {
    fmt.Println("A whois client for go.\n\nA domain name needs to be entered for the "+
                "program to function properly\n\nExample: go run whois.go google.com")
    flag.PrintDefaults()
}

func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        help()
        return  
    }
    MakeRequest()
}
