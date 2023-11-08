package main
import "fmt"

func Basics() {
	// 変数宣言
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


	// 定数宣言
	const HTTPStatusOK = 200
	fmt.Println(HTTPStatusOK)

	// 数値の区切り
	var division int = 1_000_000
	fmt.Println(division)

	// 定数宣言においての右辺の省略
	const (
		test1 = 1 + 2 // 以降、全て3になる
		test2
		test3
	)
	fmt.Println(test1, test2, test3)

	// iota：組み込みで用意された特別な定数
	// 連番生成などに用いる
	const (
		i_0 = iota
		i_1 = iota
		i_2_1, i_2_2 = iota, iota
	)
	fmt.Println(i_0, i_1, i_2_1, i_2_2)

	/* 演算子について
	算術演算子：+,-,/,*,%
	代入演算子：=,:=,+=,-=,/=,*=,++,--
	
	ビット演算子について
	論理和：|
	論理積：&
	否定：^
	排他的論理和：^(どちらか一方が真ならば真、両方真や両方偽ならば偽)
	論理積の否定：&^
	左に算術シフト：<<
	右に算術シフト：>>

	論理演算子について
	または：||
	かつ：&&
	否定：!

	比較演算子について
	→==,!=,<,>,<=,>=

	アドレス演算について(ポイント演算はできない)
	ポインタを取得：&
	ポインタが指す値を取得：*
	*/

	// 条件分岐(if文)
	x := 1
	if x == 1 {
		println("x = 1")
	} else if x == 2 {
		println("x = 2")
	} else {
		println("xは1でも2でもない")
	}

	/* Goのif文の特徴について
	→Goはコンパイル時に後ろに改行があると、自動でセミコロンを付与する
	→改行の前が"{"などの文字だと差し込まない
	
	①()がいらない
	if a == 0 {
	}

	②NGパターン1
	if a == 0
	{
	}

	③NGパターン2
	if (a == 0) fmt.Println(a)

	④代入文を書く(aはifとelseのブロックで使える↓)
	if a := f() a > 0 {
		fmt.Println(a)
	} else {
		fmt.Println(2 * a)
	}
	*/

	//条件分岐(switch文)
	
	//breakがいらない(書いても良いが、普段は書かない)
	//caseをまたぐ際は"fallthrough"を使う
	y := 1
	switch y {
	case 1, 2:
		fmt.Println("y is 1 or 2") //何もしないとbreakになる
	default:
		fmt.Println("default")
	}

	/*
	・caseに式が使える
	→大量のif文をつなぐより見通しがよくなる
	*/
	switch {
	case y == 1:
		fmt.Println("y is 1")
	}

	/* 繰り返し：for 
	→ 繰り返しはforしかない
	→ do whileに対応する機能はない(コードが難解になりやすいため)
	*/

	// 初期値; 継続条件; 更新文
	for i := 0; i <= 10; i++ {
	}

	// 継続条件のみ
	i := 0
	for i <= 10 {
	}

	// 無限ループ
	for {
		break // breakによる無限ループの脱出
	}

	/* rangeを使った繰り返し
	for i, v := range []int{1, 2, 3} {
	}
	*/

	// ラベル指定のbreak
	LOOP: //ラベル
		for {
			break LOOP
		}
}