package main

import(
	"fmt"
	"encoding/hex"
    "os"
    "log"
    "bufio"
)

func main(){
        
    file, err := os.Open("s1c4-data.txt")
    if err != nil {
        log.Fatalf("failed to open")
  
    }
    defer file.Close()
    
    var str string
    var finalkey byte
    var finalcipher string
    var finalscore int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        str = scanner.Text()
        tempkey,tempscore:=getkey(str)
        if tempscore > finalscore {
            finalscore = tempscore
            finalkey   = tempkey
            finalcipher = str
        }        
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
        //decoding the cipher usign the key

   	htob,_ := hex.DecodeString(finalcipher)
     for i,v:= range htob {
        htob[i]=v^finalkey
    }
    fmt.Print(string(htob))
}

func getkey(str string) (byte,int){

	var freqlist =[]byte{'a','e','i','o','t','n','s','r','h',' '} //optimized
	htob,_ := hex.DecodeString(str)
  
    //find most frequent single character in the cipher
    maxchar:=findmaxchar(htob)
    
    //calculate keylist from maxchar and freqlist
    var maxscore int
    var finalkey byte
    for _,v:= range freqlist{
        k:=v^maxchar //k is the key
        ks:=keyscore(htob,k) //gives score for each key
        if ks>maxscore{
            maxscore = ks
            finalkey=k
        }
    }
    return finalkey,maxscore
    
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
