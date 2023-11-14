package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Omikuzi() {
	seed := time.Now().UnixNano() // 現在時刻を数値で取得
	rand.New(rand.NewSource(seed)) // 乱数の種を設定
	val := rand.Intn(6) + 1 // 1~6の範囲で乱数生成
	
	switch {
	case val == 1:
		fmt.Println("凶")
	case val == 2 || val == 3:
		fmt.Println("吉")
	case val == 4 || val == 5:
		fmt.Println("中吉")
	case val == 6:
		fmt.Println("大吉")
	}
}