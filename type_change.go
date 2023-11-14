package main
import "fmt"

func Type_change() {
	// var sum int /* sumをintで定義すると、下記のif文でエラー発生する */
	var sum float32
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 {
		fmt.Println("good")
	}
}