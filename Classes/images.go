package main

import (
    "flag"
    "github.com/hahnicity/filetree"
    "github.com/hahnicity/filterutils"
    "os"
    "os/exec"
    "strings"
)


// XXX Allow for different image viewers to be used
const imageViewer string = "eog"

// Gallery abstract class //
type Gallery interface {
    ShowImages(imagePath string)
}

// PNG class //
type PNG struct {}

func (p *PNG) ShowImages(imagePath string) {
    /* Find all files in path with .png and shows them*/
    files := findFiles(imagePath, ".bmp")
    runCmd(files)
}

// BMP class //
type BMP struct {}

func (b *BMP) ShowImages(imagePath string) {
    /*Find all files in path with .bmp and shows them*/
    files := findFiles(imagePath, ".bmp")
    runCmd(files)
}

//Helper functions
func runCmd(files []string) {
    /* Run the image viewer */
    cmd := exec.Command(imageViewer, files...)
    cmd.Run()
}

func findFiles(imagePath, fileType string) []string {
    dir, err := filetree.GetDir(imagePath)
    if err != nil {
        panic(err)    
    }
    files, err := dir.GetFilePaths()
    if err != nil {
        panic(err)    
    }
    return filterutils.Filter(
        files, 
        func(i int) bool { return strings.HasSuffix(files[i], fileType) },
    )
}

// Argument Parser
func parseArgv() (string, string) {
    var imageType string
    var path string
    flag.StringVar(&path, "path", os.Getenv("GOPATH"), "A path to look for images")
    flag.StringVar(&imageType, "imageType", "bmp", "The type of image we wish to display")
    flag.Parse()
    return imageType, path
}

// main block
func main() {
    /*Show some images*/
    imageType, path := parseArgv()
    t := map[string]Gallery{"png": new(PNG), "bmp": new(BMP)}
    g := t[imageType]
    g.ShowImages(path)
}
