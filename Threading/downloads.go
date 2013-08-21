/*
Goal of this program is to create some kind of download tracker.
*/
package main

import (
    "fmt"
    "github.com/hahnicity/go-wget"
    "net/http"
    "os"
    "time"
)

const (
    bufSize   = 1024 * 8
    randomURL = "http://archive.ipython.org/release/1.0.0/ipython-1.0.0.zip"
)


func monitorFileSize(fileName string, downloadSize, timeout int64) {
    var (
        elapsed, start int64 = 0, 0
        progress float32 = 0
        ti *time.Time = &time.Time{}
    )
    start = ti.Unix()
    // XXX Only print when at least 1 % more has been added to file
    for progress <= 100.0 && elapsed < timeout{
        progress = getProgress(fileName, downloadSize)
        fmt.Printf("\rYour download is %f percent complete",progress)
        time.Sleep(time.Millisecond * 20)
        elapsed = ti.Unix() - start
    }
}

func getDownloadSize(url string) int64 {
    tr := new(http.Transport)
    client := &http.Client{Transport: tr}
    resp, err := client.Get(randomURL)
    if err != nil {
        panic(err)
    }
    return resp.ContentLength
}


//Monitor the progress of the download file
func getProgress(fileName string, downloadSize int64) float32 {
    file, _ := os.Open(fileName)
    stats, _ := file.Stat()
    progress := float32(stats.Size())/float32(downloadSize) * 100
    file.Close()
    return progress
}


// Main function
func main() {
    fileName := "foo"
    resp := wget.MakeRequest(randomURL)
    go monitorFileSize(fileName, resp.ContentLength, 100)
    wget.WriteToFile(fileName, resp)
    wget.Wget(randomURL, fileName)
    fmt.Print("\n")
}
