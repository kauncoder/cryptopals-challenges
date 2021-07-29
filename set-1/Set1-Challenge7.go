package main

import (
	"fmt"
	"encoding/base64"
    "os"
    "bufio"
    "crypto/aes"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
    fileName:="s1c7-data.txt"
    cipherBytes:=decodeCipher(fileName)
    plaintext:=decrypt(cipherBytes,key)
    fmt.Println("PLAINTEXT:\n",plaintext) 
}

//copy file contents and decode them from base64 to bytes
func decodeCipher(fileName string) [] byte{
    
    file,_:=os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var cipherBytes []byte
    for scanner.Scan() {
            stob,_ := base64.StdEncoding.DecodeString(scanner.Text())
            cipherBytes = append(cipherBytes,[]byte(stob)...)
    }
    return cipherBytes
}

//take keysized bytes and use AES decryption on them
func decrypt (cipherBytes []byte,key []byte) string {
    keySize:=len(key)
    cipher, _ := aes.NewCipher(key)
    for i:=0;i<len(cipherBytes);i+=keySize {
        cipher.Decrypt(cipherBytes[i:i+keySize], cipherBytes[i:i+keySize])
    }
    return string(cipherBytes) //can return this since decryption was done in place
}
                  
