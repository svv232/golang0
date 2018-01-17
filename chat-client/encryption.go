package main

import (
    "fmt"
    "crypto/sha1"
)

func stringHash(message string){
    h := sha1.New()
    h.Write([]byte(message))
    bs := h.Sum([]byte{})
    fmt.Println(bs)
}

func main(){
    stringHash("test")
}
