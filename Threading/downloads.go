/*
Goal of this program is to create some kind of download tracker. I am no where close to
done with this. However I have accomplished the part requiring the creation of a wget
type utility
*/
package main

import (
    "bufio"
    "io"
    "net/http"
    "os"
)

const (
    bufSize   = 1024 * 8
    randomURL = "http://archive.ipython.org/release/1.0.0/ipython-1.0.0.zip"
)

//Main function
func main() {
    tr := new(http.Transport)
    client := &http.Client{Transport: tr}
    resp, err := client.Get(randomURL)
    if err != nil {
        panic(err)    
    }
    file, err := os.OpenFile("foo", os.O_CREATE|os.O_WRONLY, 0777)
    if err != nil {
        panic(err)
    }
    bufferedWriter := bufio.NewWriterSize(file, bufSize)
    if err != nil {
        panic(err)
    }
    _, err = io.Copy(bufferedWriter, resp.Body)
    if err != nil {
        panic(err)
    }
}
