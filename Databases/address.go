package main

import (
    "fmt"
    "github.com/eaigner/hood"
)


type NewContacts struct {
    Id      hood.Id
    Name    string `validate:"presence"`
    Address string 
}


func main() {
    hd, err := hood.Open("postgres", "user=greg password=greg dbname=computerblue")    
    if err != nil {
        panic(err)    
    }
    
    hd = hd.Begin()
    fmt.Println(hd.IsTransaction())
    err = hd.CreateTable(&NewContacts{})
    if err != nil {
        panic(err)    
    }

    contacts := []NewContacts{
        NewContacts{Name: "Baz", Address: "1012 Horton St. Emeryville CA"},
        NewContacts{Name: "Bill", Address: "2040 Hahn St. Berkeley CA"},
    }
    _, err = hd.SaveAll(&contacts)
    if err != nil {
        panic(err)    
    }
    err = hd.Commit()
    if err != nil {
        panic(err)    
    }
}


