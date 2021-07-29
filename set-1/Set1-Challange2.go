package main

import (
	"fmt"
	"encoding/hex"
	"log"
)

func main(){

	//convert hex to byte
	var s1 string = "1c0111001f010100061a024b53535009181c"
	var s2 string = "686974207468652062756c6c277320657965"
	hextobyte1,err := hex.DecodeString(s1)
	hextobyte2,err := hex.DecodeString(s2)
	if (err!=nil){
	log.Fatal(err)
	}
	
	finalbyte := make([]byte,len(hextobyte1))
    for i,_:= range finalbyte {
		finalbyte[i] = hextobyte1[i] ^ hextobyte2[i] //XOR values in bytes
	}
	bytetohex := hex.EncodeToString(finalbyte)
	fmt.Println(bytetohex)

}
