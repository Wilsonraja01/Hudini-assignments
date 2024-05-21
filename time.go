package main

import (
	"fmt"
)
func average(array []int) int{
    avg:=0
    for i:=0;i<len(array);i++{
        avg+=array[i]
    }
    return (avg/len(array))
}
func checkAge(number int) string{
    if(number<18){
        return "Minor"
    }else if(number<=30){
        return "A Young Adult"
    }else{
        return "Adult"
    }
}
func reverseString(words string) string {
    var result string
    for _, v := range words {
        result = string(v) + result
    }
    return result
}
func largestNumber(arr []int) int {
    max:=arr[0]
    for i:=1;i<len(arr);i++{
        if(arr[i]>max){
            max=arr[i]
        }
    }
    return max
}
type Counter struct{
	count int;
}
func (counter Counter) add()Counter{
	counter.count++
	return counter
}
func (counter Counter) subtract()Counter{
	counter.count--
	return counter
}
func (counter Counter) display()Counter{
	fmt.Println(counter)
	return counter
}
func (counter Counter) reset()Counter{
	counter.count=counter.count*0
	return counter.display()
}
func main() {
	pg1 := []int{10, 20, 30, 40, 50}
  	fmt.Println(average(pg1))
	pg2:=25
	fmt.Println(checkAge(pg2))
	pg3:="hello"
	pg3_reverse:=reverseString(pg3)
	fmt.Println(pg3_reverse)
	pg4 := []int{10, 20, 30, 40, 50}
	fmt.Println(largestNumber(pg4))
	counter:=Counter{count:0}
	counter=counter.add()
	counter=counter.add()
	counter.display()
	counter=counter.subtract()
	counter.display()
	counter=counter.reset()
	counter.display()
}
