package main
import (
	"math"
	"strings"
)

func monkeyCount(n int) []int {
	// Your code here, happy coding!
   a:=make([]int,n)
   for i:=0;i<n;i++{
	 temp:=i+1
	 a[i]=temp
	 temp=0;
   }
	return a
 }

 func MakeNegative(x int) int {
	if(x<0){
	  return x
	}else{
	  return 0-x
	}
  }

  func FindMultiples(integer, limit int) (res[]int) {
	for i:=integer;i<=limit;i++{
	  if(i%integer==0){
		res=append(res,i)
	  }
	}
	 return
   }

   func CountBy(x, n int) (res []int) {
	total:=x*n
	for i:=1;i<=total;i++{
	  if(i%x==0){
		res=append(res,i)
	  }
	}
	return
  }


func PowersOfTwo(n int) (res []uint64) {
  // your code goes here // your code goes here
  var i int
  for i=0;i<=n;i++{
    temp:=math.Pow(2,float64(i))
    res=append(res,uint64(temp))
  }
  return 
}

func GetSum(a, b int) int {
	var temp int
	if(a==b){
	  temp=a
	  return temp
	}else if(a<b){
	  for i:=a;i<=b;i++{
		temp+=i
	  }
	  return temp
	}else{
	  for i:=b;i<=a;i++{
		temp+=i
	  }
	  return temp
	}
  }

  func Points(games []string) int {
	// your code here!
	points:=0
	for _,a:=range (games){
	  if(a[0]>a[2]){
		points+=3
	  }else if(a[0]==a[2]){
		points+=1
	  }
	}
	return points
  }

  func FindShort(s string) int {
	// your code
	a:=strings.Fields(s)
	length:=len(s)
	for i:=0;i<len(a);i++{
	  if(length>len(a[i])){
		length=len(a[i])
	  }
	}
	return length
  }

  func SortMyString(s string) string {
	odd:=""
	even:=""
  //   var odd string
  //   var even string
	for i:=0;i<len(s);i++{
	  if(i==0){
		even+=string(s[i])
	  }else if(i%2==0){
		even+=string(s[i])
	  }else{
		odd+=string(s[i])
	  }
	}
	return even+" "+odd
  }

  func OddCount(n int) int{
	//your code here
	return n/2
  }

  func SquareSum(numbers []int) int {
    sum:=0
  for i:=0;i<len(numbers);i++{
    sum+=numbers[i]*numbers[i]
  }
  return sum
}

func LeastLarger(a []int, i int) int {
	//define the smallest number as big as possible
	  small:=100000
	//define index as -1 
	  index:=-1
	  for j:=0;j<len(a);j++{
		//checking the given index value is greater than the user index value
		if(a[j]>a[i]){
		  //checking the given value is smaller than the small value
		  if(a[j]<small){
			small=a[j]
			index=j
		  }
		}
	  }
	  return index
	}

