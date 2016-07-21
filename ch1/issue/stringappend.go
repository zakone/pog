package main

import (
	"bytes"
	"strings"
)

var list = []string{"aaa", "bbb", "ccc"}

func appendOperator() string {
	temp := ""
	for _, str := range list {
		temp += "," + str
	}
	return temp
}

func stringsJoin() string {
	return strings.Join(list[:], ",")
}

func appendHardCoding() string {
	return list[0] + "," + list[1] + "," + list[2]
}

func byteArray() string {
	var temp []byte
	for _, str := range list {
		temp = append(temp, str...)
		temp = append(temp, ',')
	}
	return string(temp)
}

func byteBuffer() string {
	temp := bytes.NewBuffer(make([]byte, 0))
	for _, str := range list {
		temp.WriteString(str)
		temp.WriteString(",")
	}
	return temp.String()
}
