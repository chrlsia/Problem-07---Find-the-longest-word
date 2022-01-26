package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
  

 words := splitSentIntoWords()

 max:=findMaxWord(words)

 fmt.Println(max)

}

func  splitSentIntoWords()[]string{
	fmt.Println("Give me a sentence ->")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	words := strings.Split(strings.TrimSpace(line), " ")
	return words
}

func findMaxWord(words []string)string{
	max:=words[0] //string

 for i:=1;i<len(words);i++{
   if len(words[i])>=len(max){
     max = words[i]
   }
  }
  return max
}