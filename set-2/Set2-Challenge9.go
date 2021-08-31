package main

import (
	 "fmt"
)

func main () {
    	plaintext:="YELLOW SUBMARINE"
    	paddedBlockSize:=20
    	paddedBlock:=PadBlock(plaintext,paddedBlockSize)
    	TestPadding (plaintext, paddedBlock)
}

func PadBlock(plaintext string,paddedBlockSize int) string {
        
        textLength:=len(plaintext)
        var padByte byte = '\x04' //changes length with invisible padding
        padSize:=paddedBlockSize-(textLength%paddedBlockSize)
        var padding = make([]byte,padSize,padSize)
        for i,_:=range padding {
            padding[i]=padByte
        }
        plaintext=plaintext+string(padding)
        return string(plaintext)
        
}


func TestPadding ( plaintext string, paddedBlock string) {
	
        fmt.Println("orignal block length: ",len(plaintext))
        fmt.Println("padded block length: ",len(paddedBlock))

}

