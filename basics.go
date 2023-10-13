package main
import "fmt"

func main() {
	var firstName string = "sakai"
	var first, name string = "sakai", "shun"
	var z int = 1
	var (
		f_name string = "sakai"
		l_name string = "shun"
		age int = 24
	)
	var (
		f = "s"
		l = "s"
		ag = 24
	)
	var (
		fn, ln, a = "shun", "sakai", 24
	)

	b, c := "shun", "sakai"
	tosi := 24

	fmt.Println(firstName, first, name, z, f_name, l_name, age, f, l, ag, fn, ln, a, b, c, tosi)

	const HTTPStatusOK = 200

	fmt.Println(HTTPStatusOK)
}
