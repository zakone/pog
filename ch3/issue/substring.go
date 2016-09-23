package main

import "fmt"

func main() {
	s := "left foot"
	sub := s[2:5]
	//sub[1] = 'y' //部分文字列も編集できない
	//sub[0], sub[1] = sub[1], sub[0] //交換もできない
	b := []byte(s)
	sub_b := b[:3]
	sub_b[0] = 'y'
	fmt.Println(sub)
}
