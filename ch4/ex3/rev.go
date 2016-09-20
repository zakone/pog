package main
import "fmt"

func reverse(ptr *[5]int) {
	for i, j:= 0, 4; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main() {
	a := [...]int{1,2,3,4,5}
	reverse(&a)
	fmt.Println(a)
}