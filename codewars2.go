package kata

//https://www.codewars.com/kata/54ff3102c1bad923760001f3/solutions/go
import "strings"

func GetCount(str string) (count int) {
  // Enter solution here
  parts := strings.Split(str, "")
  for i:=0;i<len(parts);i++{
      if parts[i]=="a" || parts[i]=="e" || parts[i]=="i" || parts[i]=="o"|| parts[i]=="u"{
          count++
      }
  }
  return count
}

package kata

//https://www.codewars.com/kata/586f6741c66d18c22800010a/train/go

func SequenceSum(start, end, step int) int {
  answer:=0
  if start>end{
    return 0
    }else if start==end{
    return end
  }else{
    for i:=start;i<=end;i+=step{
      answer+=i
    }
    return answer
  }
}

package kata

// https://www.codewars.com/kata/592a6ad46d6c5a62b600003f/train/go

import "strings"

func Encode(str string)string {
  word:=strings.Split(str, "")
  for i:=0;i<len(word);i++{
    if word[i]=="G"{
      word[i]="A"
    }else if word[i]=="A"{
      word[i]="G"
    }else if word[i]=="a"{
      word[i]="g"
    }else if word[i]=="g"{
      word[i]="a"
    }else if word[i]=="D"{
      word[i]="E"
    }else if word[i]=="E"{
      word[i]="D"
    }else if word[i]=="d"{
      word[i]="e"
    }else if word[i]=="e"{
      word[i]="d"
    }else if word[i]=="R"{
      word[i]="Y"
    }else if word[i]=="Y"{
      word[i]="R"
    }else if word[i]=="r"{
      word[i]="y"
    }else if word[i]=="y"{
      word[i]="r"
    }else if word[i]=="P"{
      word[i]="O"
    }else if word[i]=="O"{
      word[i]="P"
    }else if word[i]=="p"{
      word[i]="o"
    }else if word[i]=="o"{
      word[i]="p"
    }else if word[i]=="L"{
      word[i]="U"
    }else if word[i]=="U"{
      word[i]="L"
    }else if word[i]=="l"{
      word[i]="u"
    }else if word[i]=="u"{
      word[i]="l"
    }else if word[i]=="K"{
      word[i]="I"
    }else if word[i]=="I"{
      word[i]="K"
    }else if word[i]=="k"{
      word[i]="i"
    }else if word[i]=="i"{
      word[i]="k"
    }
  }
  str_encode:=strings.Join(word,"")
	return str_encode
}

func Decode(str string) string {
  word:=strings.Split(str, "")
  for i:=0;i<len(word);i++{
    if word[i]=="G"{
      word[i]="A"
    }else if word[i]=="A"{
      word[i]="G"
    }else if word[i]=="a"{
      word[i]="g"
    }else if word[i]=="g"{
      word[i]="a"
    }else if word[i]=="D"{
      word[i]="E"
    }else if word[i]=="E"{
      word[i]="D"
    }else if word[i]=="d"{
      word[i]="e"
    }else if word[i]=="e"{
      word[i]="d"
    }else if word[i]=="R"{
      word[i]="Y"
    }else if word[i]=="Y"{
      word[i]="R"
    }else if word[i]=="r"{
      word[i]="y"
    }else if word[i]=="y"{
      word[i]="r"
    }else if word[i]=="P"{
      word[i]="O"
    }else if word[i]=="O"{
      word[i]="P"
    }else if word[i]=="p"{
      word[i]="o"
    }else if word[i]=="o"{
      word[i]="p"
    }else if word[i]=="L"{
      word[i]="U"
    }else if word[i]=="U"{
      word[i]="L"
    }else if word[i]=="l"{
      word[i]="u"
    }else if word[i]=="u"{
      word[i]="l"
    }else if word[i]=="K"{
      word[i]="I"
    }else if word[i]=="I"{
      word[i]="K"
    }else if word[i]=="k"{
      word[i]="i"
    }else if word[i]=="i"{
      word[i]="k"
    }
  }
  str_decode:=strings.Join(word,"")
	return str_decode
}

package kata

// https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

func FindOdd(seq []int) int {
    numCount := make(map[int]int)
    for _, num := range seq {
        numCount[num]++
    }
    answer:=-1
    for num, count := range numCount {
        if count%2 != 0 {
            answer=num
        }
    }
    return answer
}

package kata

// https://www.codewars.com/kata/51b6249c4612257ac0000005

import "strings"

func Decode(roman string) int {
  answer:=0
  words:=strings.Split(roman,"")
  numbers:=make([]int,len(words))
  for i:=0;i<len(words);i++{
    if(words[i]=="I"){
      numbers[i]=1
    }else if(words[i]=="V"){
        numbers[i]=5
    }else if(words[i]=="X"){
        numbers[i]=10
    }else if(words[i]=="L"){
        numbers[i]=50
    }else if(words[i]=="C"){
        numbers[i]=100
    }else if(words[i]=="D"){
        numbers[i]=500
    }else if(words[i]=="M"){
        numbers[i]=1000
    }
  }
for i:=0;i<len(numbers);i++{
    if(i+1<len(numbers)){
        if(numbers[i]<numbers[i+1]){
            temp:=numbers[i+1]-numbers[i]
            answer+=temp
            i+=1
        }else{
            answer+=numbers[i]
        }
    }else{
        answer+=numbers[i]
    }
}
  return answer
}

package kata

// https://www.codewars.com/kata/554e4a2f232cdd87d9000038/train/go

import "strings"

func DNAStrand(dna string) string {
  // your code here
  words:=strings.Split(dna,"")
  for i:=0;i<len(words);i++{
    if(words[i]=="A"){
      words[i]="T"
    }else if(words[i]=="T"){
      words[i]="A"
    }else if(words[i]=="C"){
      words[i]="G"
    }else if(words[i]=="G"){
      words[i]="C"
    }
  }
  return strings.Join(words,"")
}

package kata

// https://www.codewars.com/kata/56747fd5cb988479af000028/train/go

import "strings"

func GetMiddle(s string) string {
  //Code goes here!
  words:=strings.Split(s,"")
  if(len(words)%2!=0){
    return words[(len(words)/2)]
  }else{
    a:=[]string{}
    a=append(a,words[(len(words)/2)-1])
    a=append(a,words[(len(words)/2)])
    return strings.Join(a,"") 
  }
}