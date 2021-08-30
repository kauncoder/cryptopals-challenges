package main

import (
	 "fmt"
     "encoding/base64"
     "crypto/aes"
     "os"
     "bufio"
)

//to CBC and decrypt
func main () {

    key:="YELLOW SUBMARINE"
    IV:="\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
    fileName:="s2c10-data.txt"
    stob:=decodeCipher(fileName)
    plaintext:=decrypt(stob,[]byte(key),[]byte(IV))
    fmt.Println(plaintext)
}

//use AES-CBC decryption to generate plaintext file
func decrypt (cipherBytes []byte,key []byte,IV []byte) string {
    keySize:=len(key)
    cipher, _ := aes.NewCipher(key)
    plaintext:=make([]byte,0)
    plainblock:=make([]byte,keySize)
    //decrypt
    for i:=0;i<len(cipherBytes);i+=keySize {
        //decrypt
        cipher.Decrypt(plainblock,cipherBytes[i:i+keySize])
        //XOR with IV
        for j,v:=range plainblock {
            plainblock[j]=v^IV[j]
        }
        //append to the plaintext slice
        plaintext=append(plaintext,plainblock...)
        //update IV value
        IV = cipherBytes[i:i+keySize]
    }
    return string(plaintext)
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
