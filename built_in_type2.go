package main
import "fmt"

func Built_in_type2() {
	var a, b, c bool
	// 「「aかつbがtrue(論理積がtrue)」または(OR)「cがfalse」」の時のみ"true"が表示される
	if a && b || !c {
	fmt.Println("true")
	} else {
	fmt.Println("false")
	}
}