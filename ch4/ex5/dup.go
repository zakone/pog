package main
import "fmt"

func dup(strings []string) []string{
	for i:= 0; i < len(strings)-1; {
		if strings[i] == strings[i+1] {
			strings = remove(strings, i)
		} else{
			i++
		}
	}
	return strings
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}


func main() {
	s := []string{"aa", "aa", "bbb", "c", "ddddd", "ddddd"}
	fmt.Println(dup(s))
}