package main

import(
	"fmt"
	"encoding/base64"
    "os"
    "bufio"
    "math/bits"
)

func main(){
    
    file,_:=os.Open("s1c6-data.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var cipherBytes []byte
    for scanner.Scan() {
            stob,_ := base64.StdEncoding.DecodeString(scanner.Text())
            cipherBytes = append(cipherBytes,[]byte(stob)...)
    }
    key:=getRotateKey(cipherBytes,getKeySize(cipherBytes))
    decrypt(cipherBytes,key)
    
    fmt.Println("KEY: ",string(key)) // "Terminator X: Bring the noise"
    fmt.Println("PLAINTEXT:\n",string(cipherBytes))
}

func getKeySize (cipherBytes []byte) int {

    var distMap = map [int] int {1:1000000}
    var bestKeySize int
    var minDistance int = 100000 //arbitrary very high value

    //calcualte normalised hamming distance for possible keysizes
    for keysize:=2;keysize <40;keysize ++ {
        for i:=0;i<len(cipherBytes)-(2*keysize);i+=keysize {
            if  _,ok:=distMap[keysize]; ok {
            distMap[keysize]+=(getHammingDistance(cipherBytes[i:i+keysize],cipherBytes[i+keysize:i+(2*keysize)]))
            } else {
                distMap[keysize]=getHammingDistance(cipherBytes[i:i+keysize],cipherBytes[i+keysize:i+(2*keysize)])
            }
        }
    }
    //get key from minimum distance
    for i,v:= range distMap {
        if v < minDistance {
                bestKeySize = i
                minDistance = v
            }
    }
    return bestKeySize
}

func getHammingDistance(stobA [] byte, stobB [] byte) int {
        var distance int = 0
        for i,_ := range stobA {
            //for each byte we need to calculate the differences in bits
            xorb:=stobA[i]^stobB[i]
            distance += bits.OnesCount(uint(xorb)) //counting 1s in xorb to get hamming distance
        }
        return distance
}

func getRotateKey (cipherBytes []byte, keySize int) []byte {
    
    var rotateKey []byte
    for i:=0;i<keySize;i++ {
        //calcualte score to find key for each location
        var transposeBytes [] byte
        var score int
        var key byte
        for j:=i;j<len(cipherBytes);j+=keySize {
            transposeBytes=append(transposeBytes,cipherBytes[j])
        }
        tempKey,tempScore:=getKey(transposeBytes)
        if tempScore>score {
            score = tempScore
            key   = tempKey
        }
        rotateKey=append(rotateKey,key)
    }
    return rotateKey
}

func decrypt (text []byte ,rotateKey []byte) {
    keyLength:=len(rotateKey)
    for i,v:= range text {
        text[i]=v^rotateKey[i%keyLength]
    }
}

func getKey(cipherBytes []byte) (byte,int){

	var freqList =[]byte{'a','e','i','o','t','n','s','r','h',' '} //optimized for use instead of going through all characters
	//find most frequent single character in the cipher
    maxChar:=findMaxChar(cipherBytes)
    //calculate keylist from maxchar and freqlist
    var score int
    var key byte
    for i:=0;i<len(freqList);i++{
        k:=freqList[i]^maxChar //k is the key (byte)
        ks:=keyScore(cipherBytes,k) //gives score(int) for each key
        if ks>score{
            score = ks
            key=k
        }
    }
    return key,score
    
}

//from Set1-Problem3
func findMaxChar(cipherBytes []byte) byte{
    
    freqScore := map[byte]int{}
    var maxFreq int
    var maxChar byte
    
    for _,v:= range cipherBytes{
        j, _ := freqScore[v]
        freqScore[v]=j+1
    }
    for k,v:=range freqScore{
        if (v>maxFreq){
            maxFreq=v
            maxChar=k
        }
    }
    return maxChar
}

//from Set1-Problem3
func keyScore(cipherBytes []byte,key byte) int {
    
    var decodedString []byte
    var score int
    freqWeight :=map [byte]int{'a':82,'e':130,'i':70,'o':75,'t':91,'n':67,'s':63,'r':60,'h':61,' ':240} //weights based on text analysis research
    freqScore := map [byte]int{'a':0,'e':0,'i':0,'o':0,'t':0,'n':0,'s':0,'r':0,'h':0,' ':0}
    
    //decoding the original sentence using key
    for _,v:= range cipherBytes {
        decodedString=append(decodedString,v^key)
    }
    
    //calculate frequent-wqords occurence for the sentence   
    for _,v:= range decodedString{
             freqScore[v]+=1
    }
    
    //final score calculation
    for k,v:=range freqScore{
        score=score+v*freqWeight[k]
    }
    return score
}
