package main

import (
	 "fmt"
     "crypto/aes"
     "crypto/rand"
)

//one function to generate AES keySize
//have two functions one uses cbc and the other ecb
//next function to genrate key and encrypt; this one takes the plaintext, appends random bytes to the plaintext, flips a coin to chose cbc or ecb mode and encrypts in that mode with AES
//last is a blockmode detection function
//have added a function to generate reasonably detectable ciphertexts

var globalkeysize int = 16

func main () {

    basetext:=[]byte("YELLOW SUBMARINE")
    plaintext:=GenerateRepetitivePlainTexts(basetext)
    howmanyciphers:=10  //how many times should the plaintext be encrypted
    for i:=0;i<howmanyciphers;i++ {
        ciphertext,ciphermode:=EncryptAES(plaintext)
        modedetected:=DetectCipherMode(ciphertext)
        if modedetected==ciphermode {
            fmt.Println("correct detection")
        } else {
            fmt.Println("wrong detection")
        }
    }
}

//generate repetitive texts for better chances of detecting ECB (only for demo reasons)
func GenerateRepetitivePlainTexts (basetext []byte) []byte {

    reps:=GetRandomValueUnder256(40)+60 //take reps value bw 60-100
    plaintext:=make([]byte,0)
    for i:=0;i<int(reps);i++ {
        getnum:=GetRandomValueUnder256(5)+5 //takes 5-10 bytes
        addbytes:=make([]byte,getnum)
        _, errRand := rand.Read(addbytes)
        if errRand != nil {
            fmt.Println("error in byte generation:", errRand)
        }
        addedblock:=append(basetext,addbytes...)
        plaintext=append(plaintext,addedblock...)
    }
    return plaintext
}

//function to encrypt using AES with one of two block modes
func EncryptAES (originaltext []byte) (string,string) {
    
    keySize:=globalkeysize
    key:=GenerateRandomAESKey(keySize)
    //appending "random" values to both ends 
    appendedtext:=AddRandomBytes(originaltext,7,7)
    plaintext:=padBlock(appendedtext,keySize)
    //generate IV value (random)
    IV:=make([]byte,keySize)
    _, errIV := rand.Read(IV)
    if errIV != nil {
        fmt.Println("error in Initialization Vector generation:", errIV)
    }
    
    flipcoin:=GetRandomValueUnder256(2)
    if flipcoin == 0 {
        ciphertext:=AESECBEncrypt(plaintext,key)
        return ciphertext,"ECB"
    } else {
        ciphertext:=AESCBCEncrypt(plaintext,IV,key)
        return ciphertext,"CBC"
    }
}

//function for detecting ciphermode
func DetectCipherMode(ciphertext string) string{
    keySize:=globalkeysize
    var repeatCount = map [string]int{}
    
    for i:=0;i<(len(ciphertext));i+=keySize {
        index:= ciphertext[i:i+keySize]
        if _,ok:=repeatCount[index];ok {
            return "ECB"
        } else {
            repeatCount[index] = 1
        }
    
    }
    return "CBC"
}

func GenerateRandomAESKey (keySize int) []byte{
    key:=make([]byte,keySize)
    _, errRand := rand.Read(key)
    if errRand != nil {
        fmt.Println("error in key generation:", errRand)
    }
    return key
}

func GetRandomValueUnder256 (max int) int {
    b:=[]byte{0}
    _, errreps := rand.Read(b)
    if errreps != nil {
        fmt.Println("error in random value generation:", errreps)
    }
    return int(b[0]%byte(max))   //return random int value under max
}

func AddRandomBytes (plaintext []byte,frontadding int, backadding int) []byte {
    
    addedtext:=make([]byte,0,len(plaintext)+backadding+frontadding)
    
    frontadd:=make([]byte,frontadding)
     _, errRand := rand.Read(frontadd)
    if errRand != nil {
        fmt.Println("error in byte generation:", errRand)
    }
    addedtext = append(addedtext,frontadd...)
    addedtext = append(addedtext,plaintext...)    
    backadd:=make([]byte,backadding)
     _, errRand = rand.Read(backadd)
    if errRand != nil {
        fmt.Println("error in byte generation:", errRand)
    }
    addedtext = append(addedtext,backadd...)
    return addedtext
    
}

func padBlock(plaintext []byte,paddedBlockSize int) []byte {
        
        textLength:=len(plaintext)
        var padByte byte = '\x04'   //padding characters
        padSize:=paddedBlockSize-(textLength%paddedBlockSize)
        var padding = make([]byte,padSize,padSize)
        for i,_:=range padding {
            padding[i]=padByte
        }
        plaintext=append(plaintext,padding...)
        return plaintext
        
}

//use AES-ECB encryption
func AESECBEncrypt (plaintext []byte,key []byte) string {
    keySize:=len(key)
    ciphertext:=make([]byte,0)
    cipherblock:=make([]byte,keySize)
    cipher, errAES := aes.NewCipher(key)
    if errAES != nil {
        fmt.Println("error in AES initialization:", errAES)
    }
    //encrypt
    for i:=0;i<len(plaintext);i+=keySize {
        //encrypt
        cipher.Encrypt(cipherblock,plaintext[i:i+keySize])
        //append to the ciphertext slice
        ciphertext=append(ciphertext,cipherblock...)
    }
    return string(ciphertext)
}

//use AES-CBC encryption to
func AESCBCEncrypt (plaintext []byte,IV []byte,key []byte) string {
    keySize:=len(key)
    ciphertext:=make([]byte,0)
    cipherblock:=make([]byte,keySize)
    cipher, errAES := aes.NewCipher(key)
    if errAES != nil {
        fmt.Println("error in AES initialization:", errAES)
    }
    
    //encrypt
    for i:=0;i<len(plaintext);i+=keySize {
        //XOR plaintext with IV
        for j,v:=range IV {
             plaintext[i+j]=v^plaintext[i+j]
        }
        //encrypt
        cipher.Encrypt(cipherblock,plaintext[i:i+keySize])
        //append to the ciphertext slice
        ciphertext=append(ciphertext,cipherblock...)
        //update IV value
        IV = cipherblock
    }
    return string(ciphertext)
}






