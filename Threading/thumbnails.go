package main

import (
    //"github.com/nfnt/resize"
    "github.com/hahnicity/filetree"
    "github.com/hahnicity/filterutils"
    "math/rand"
    "fmt"
    "strings"
)

const (
    imagePath string = "/home/greg/Pictures"
)

func getImages(imagePath string) []string {
   imageTypes := [...]string{".bmp", ".png", ".jpg"}
   dir, err := filetree.GetDir(imagePath)
   if err != nil {
        panic(err)   
    }
    files, err := dir.GetFilePaths()
    if err != nil {
        panic(err)    
    }
    allFiles := make([]string,0)
    for _, imageType := range imageTypes {
        filtered := filterutils.Filter(
            files,
            func(i int) bool { return strings.HasSuffix(files[i], imageType) },
        )
        allFiles = append(allFiles[:len(allFiles)], append(make([]string,0), filtered[:len(filtered)]...)...)
    }
    return allFiles
}

// Pick dimensions that we wish to resize our thumbnail to
func pickDimensions(bound int32) (uint32, uint32){
    var (
        height, width uint32 = uint32(rand.Int31n(bound) + 1), uint32(rand.Int31n(bound) + 1)
    )
    return height, width
}

func main() {
    fmt.Println(rand.Uint32())
}
