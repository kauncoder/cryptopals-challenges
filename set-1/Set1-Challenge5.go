package main

import (
	"fmt"
	"encoding/hex"
)

func main(){
	str := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
	stob:=[]byte(str)
	rkey:=[]byte{'I','C','E'}
	encodeXOR(stob,rkey)
    btoh:=hex.EncodeToString(stob)
    fmt.Println(btoh)
}

func encodeXOR (ptext []byte,rkey []byte) {
    keylength:=len(rkey)
    for i,v:= range ptext {
        ptext[i]=v^rkey[i%keylength]
    }
}
