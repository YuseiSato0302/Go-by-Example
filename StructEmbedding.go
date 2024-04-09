package main

import "fmt"

// 基本構造体
type base struct {
	num int
}

// 基本構造体のメゾット
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// コンテナ構造体(基本構造体を埋め込んだ構造体)
// 他言語でいう継承に近いが、Go言語では継承という概念はなく「合成」に重点を置いている
type container struct {
	base
	str string
}

func main() {

	// コンテナ構造体の初期化
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// コンテナ構造体のフィールドへのアクセス
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	// 基本構造体のメゾットへのアクセス
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// インターフェイスの実装
	var d describer = co
	fmt.Println("describe:", d.describe())
}
