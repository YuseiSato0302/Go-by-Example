package main

import (
	"fmt"
	"math"
)

/*
Go言語における「インターフェイス」は、特定のメソッドの集まりを定義する方法です。
これはちょっとした約束みたいなもので、「このメソッドを実装しているなら、このインターフェイスを実装しているとみなすよ」という約束です。
インターフェイスは具体的な動作を定義しないで、メソッドの形（名前、引数、返り値の型）だけを指定します。
*/

// インターフェイスの定義
type geometry interface {
	area() float64
	perim() float64
}

// 構造体1の定義
type rect struct {
	width, height float64
}

// 構造体2の定義
type circle struct {
	radius float64
}

// 構造体1のメソッド1
func (r rect) area() float64 {
	return r.width * r.height
}

// 構造体1のメソッド2
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 構造体2のメソッド1
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 構造体2のメソッド2
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// インターフェイスを引数に取る関数
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// メイン関数
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
