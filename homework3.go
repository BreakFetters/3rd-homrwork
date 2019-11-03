package main

import (
	"fmt"
	"os"
)

func main() {
	file1,_:=os.Create("proverb.txt")
	file1.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	file2,_:=os.Open("proverb.txt")
	b:=make([]byte,1024)
	file2.Read(b)
	fmt.Println(string(b))
}