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

// Make the GET request to a server, return the response
func getResponse() *http.Response {
    tr := new(http.Transport)
    client := &http.Client{Transport: tr}
    resp, err := client.Get(randomURL)
    errorChecker(err)
    return resp
}

// Write the response of the GET request to file
func writeToFile(resp *http.Response) {
    // Credit for this implementation should go to github user billnapier
    file, err := os.OpenFile("foo", os.O_CREATE|os.O_WRONLY, 0777)
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
    writeToFile(resp)
}
