package main

import "fmt"

// Interface
type Bird interface {
	Bark() string
}

// 構造体
type Duck struct {
	Cry string
}

// メソッド
func (d Duck) Bark() string {
	return d.Cry
}

func main() {
	// 型推論
	cry := "ga-ga"

	duck := Duck{
		Cry: cry,
	}
	bark(duck)
}

func bark(bird Bird) {
	fmt.Println(bird.Bark())
}
