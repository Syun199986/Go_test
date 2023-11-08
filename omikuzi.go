package main
import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		str_i := strconv.Itoa(i)
		if i % 2 == 0 {
			fmt.Println(str_i + "-偶数")
		} else {
			fmt.Println(str_i + "-奇数")
		}
	}
}