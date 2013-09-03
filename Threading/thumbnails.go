package main

import (
    "github.com/nfnt/resize"
    "github.com/hahnicity/filetree"
    "github.com/hahnicity/filterutils"
    "image/jpeg"
    "log"
    "math/rand"
    "os"
    "strconv"
)

type Dimensions struct {
    height, width uint
}

const (
    FindImagePath string = "/home/greg/Pictures/"
    PutImagePath string = "/home/greg/Pictures/"
)

// XXX Arrays arent immutable by nature so you cant make them constant
var Choices = []*Dimensions{
    &Dimensions{75,75}, 
    &Dimensions{125,125}, 
    &Dimensions{150,150},
    &Dimensions{200,200},
}

// Get a list of all supported image files in a directory
func getImages(imagePath string) []string {
    // Currently only supports .jpg's
   imageTypes := [...]string{".jpg"}
   dir, err := filetree.GetDir(imagePath)
   if err != nil {
        panic(err)   
    }
    files, err := dir.GetFilePaths()
    if err != nil {
        panic(err)    
    }
    return filterutils.HasSuffix(files, imageTypes[0])

    // XXX Commented code is not a good practice, but this was good learning material
    //allFiles := make([]string,0)
    //for _, imageType := range imageTypes {
    //    filtered := filterutils.HasSuffix(files, imageType)
    //    allFiles = append(
    //        allFiles[:len(allFiles)], 
    //        append(
    //            make([]string,0), 
    //            filtered[:len(filtered)]...
    //        )...
    //    )
    //}
}

// Pick dimensions that we wish to resize our thumbnail to
func pickDimensions() *Dimensions {
    n := int32(len(Choices))
    return Choices[rand.Int31n(n)]
}

// Resize all images given to it
func resizeImages(path string, dim *Dimensions, progress chan int) {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)        
        return
    }
    img, err := jpeg.Decode(file)
    file.Close()
    if err != nil {
        log.Fatal(err)
        return   
    }
    res := resize.Resize(dim.height, dim.width, img, resize.Lanczos3)
    out, err := os.Create(PutImagePath + strconv.FormatInt(rand.Int63(),36) + ".jpg")
    if err != nil {
        log.Fatal(err)    
    }
    defer out.Close()
    jpeg.Encode(out, res, nil)
    progress <- 1
}

// Monitor the overall progress of the file conversion
func monitorProgress(progress chan int, n int) {
    totalFiles := 0
    for received := range progress {
        totalFiles = totalFiles + received
        // XXX What happens if theres an error? This file will loop infinitely
        if totalFiles >= n {
            close(progress)    
        }
    }
    
}

func main() {
    // XXX Move path to argument parser
    images := getImages(FindImagePath)
    // Implements a channel. They are simply conveyer belts.
    progress := make(chan int)
    for _, path := range(images) {
        // Every file gets its own thread! YAY!
        go resizeImages(path, pickDimensions(), progress)
    }
    monitorProgress(progress, len(images))
}
