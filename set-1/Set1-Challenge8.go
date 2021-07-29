package main

import (
	"fmt"
    "os"
    "bufio"
)

func main() {

    //can directly work on hex in this problem
    keySize:=32 //hex keysize is double of bytes
    fileName:="s1c8-data.txt"
    fmt.Println("String in ECB mode is: ",detectECB(fileName,keySize))
}

//to find a repeating substring of size 32
func detectECB(fileName string,keySize int) string {
    
    file,_:=os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        str:=scanner.Text()
        var repeatCount = map [string]int{}
        for i:=0;i<(len(str));i+=keySize {
            index:= str[i:i+keySize]
            if _,ok:=repeatCount[index];ok {
                return str
            } else {
                repeatCount[index] = 1
            }
        }
    }
    return "not ECB"
}
                  
