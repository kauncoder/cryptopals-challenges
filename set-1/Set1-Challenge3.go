package main

import(
	"fmt"
	"encoding/hex"
)

func main(){

	var freqlist =[]byte{'a','e','i','o','t','n','s','r','h',' '} //optimized
    str:="1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	htob,_ := hex.DecodeString(str)
  
    //find most frequent single character in the cipher
    maxchar:=findmaxchar(htob)
    
    //calculate keylist from maxchar and freqlist
    var maxscore int
    var finalkey byte
    for _,v:= range freqlist{
        k:=v^maxchar //k is the key
        ks:=keyscore(htob,k) //returns weighted score for each key
        if ks>maxscore{
            maxscore = ks
            finalkey=k
        }
    }
    
    //decoding the cipher using the key
    for i,v:= range htob{
        htob[i]=v^finalkey
    }
    fmt.Println(string(htob))
}


func findmaxchar(htob []byte) byte{
    
    freqscore := map[byte]int{}
    var maxfreq int
    var maxchar byte
    
    for _,v:= range htob{
        j, _ := freqscore[v]
        freqscore[v]=j+1
    }
    for k,v:=range freqscore{
        if (v>maxfreq){
            maxfreq=v
            maxchar=k
        }
    }
    return maxchar
}

func keyscore(htob []byte,keyx byte) int{
    
    var decodedstring []byte
    var score int
    freqweight :=map [byte]int{'a':82,'e':130,'i':70,'o':75,'t':91,'n':67,'s':63,'r':60,'h':61,' ':240} //weights based on text analysis research
    freqscore := map [byte]int{'a':0,'e':0,'i':0,'o':0,'t':0,'n':0,'s':0,'r':0,'h':0,' ':0}
    
    //decoding the original sentence using key
    for _,v:= range htob {
        decodedstring=append(decodedstring,v^keyx)
    }
    
    //calculate frequent-wqords occurence for the sentence   
    for _,v:= range decodedstring{
             freqscore[v]+=1
    }
    
    //final score calculation
    for k,v:=range freqscore{
        score=score+v*freqweight[k]
    }
    return score
}
