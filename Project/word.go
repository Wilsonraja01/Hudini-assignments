package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("--------------------------------------")
        fmt.Println("1. Enter a block of text: ")
        fmt.Println("2. Exit")
        fmt.Println("Enter choice: ")
        input, _ := reader.ReadString('\n')
        choice, _ := strconv.Atoi(strings.TrimSpace(input))
		fmt.Println("You have choosen ",choice)
		fmt.Println("--------------------------------------")
		switch choice{
		case 1:
			fmt.Println("Enter the Block of Text: ")
			input, _ = reader.ReadString('\n')
			words:=strings.Fields(input)
			wordCount := make(map[string]int)
			for _, word := range words {
				if wordCount[word]==0{
					wordCount[word]+=1
				}else{
				wordCount[word]++
				}
			}
			fmt.Println("--------------------------------------")
			for word, count := range wordCount {
				fmt.Printf("words : %s and the Count is : %d\n", word, count)
			}
		case 2:
			fmt.Println("You Pressed exit...")
			fmt.Println("--------------------------------------")
			os.Exit(0)
		default:  
			fmt.Println("Invalid Choice")
		}
	}
}