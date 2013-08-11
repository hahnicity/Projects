/*
Goal of this program is to create some kind of download tracker.
*/
package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)

const (
    bufSize   = 1024 * 8
    randomURL = "http://archive.ipython.org/release/1.0.0/ipython-1.0.0.zip"
)

// Make the GET request to a server, return the response
func getResponse() *http.Response {
    tr := new(http.Transport)
    client := &http.Client{Transport: tr}
    resp, err := client.Get(randomURL)
    errorChecker(err)
    return resp
}

func monitorFileSize(fileName string, downloadSize, timeout int64) {
    var (
        elapsed, start, size int64 = 0, 0, 0
        ti *time.Time = &time.Time{}
    )
    start = ti.Unix()
    // XXX Only print when at least 1 % more has been added to file
    for size < downloadSize && elapsed < timeout{
        file, _ := os.Open(fileName)
        stats, _ := file.Stat()
        size := stats.Size()
        progress := float32(size)/float32(downloadSize) * 100
        fmt.Println("Your download is %", progress, "percent finished")
        elapsed = ti.Unix() - start
        file.Close()
    }
}

// Write the response of the GET request to file
func writeToFile(fileName string, resp *http.Response) {
    // Credit for this implementation should go to github user billnapier
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
    defer file.Close()
    bufferedWriter := bufio.NewWriterSize(file, bufSize)
    errorChecker(err)
    _, err = io.Copy(bufferedWriter, resp.Body)
    errorChecker(err)
}

// Check if we received an error on our last function call
func errorChecker(err error) {
    if err != nil {
        panic(err)
    }
}

// Main function
func main() {
    resp := getResponse()
    fileName := "foo"
    go monitorFileSize(fileName, resp.ContentLength, 100)
    writeToFile(fileName, resp)
}
