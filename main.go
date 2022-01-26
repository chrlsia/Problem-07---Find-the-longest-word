package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
  fmt.Println("Give me a sentence ->")
  r := bufio.NewReader(os.Stdin)
  line, err := r.ReadString('\n')
  if err != nil {
	panic(err)
 }

 words := strings.Split(strings.TrimSpace(line), " ")

 max:=words[0] //string

 for i:=1;i<len(words);i++{
   if len(words[i])>=len(max){
     max = words[i]
   }
  }

 fmt.Println(max)

}