package main

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
    "log"
)

func main() {
	var str string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    //convert hex string to bytes
	x,errx := hex.DecodeString(str)
    if (errx!=nil){
        log.Fatal(errx)
	}
	//convert byte to base64
	y := base64.StdEncoding.EncodeToString(x)
	fmt.Println(y)
}
