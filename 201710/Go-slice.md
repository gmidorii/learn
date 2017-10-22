# Goのsliceって結局どう書けばいいの？

## 概要
色々な資料を見返しながら、sliceの実装を書いてる気がするので、
一旦自分で調べてまとめます。

## 現状理解

### Array
- 固定長配列
  - lenを変える事はできない
- 型に配列長も含まれる
	- ex) `[3]int`

### Slice
- 可変長配列
  - Arrayへの参照を持っている
- len
  - sliceの長さ
- cap
  - 確保したメモリの長さ
- 変更時
  - capを上回っている場合は、allocしてcapを確保してlen追加

## サンプル実装と測定
### Slice生成
#### 測定
宣言時に追加
```go
func BenchmarkInitDeclation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		b.StopTimer()
		fmt.Sprint(s)
		b.StartTimer()
	}
}
```

len=0で初期化
```go
func BenchmarkInitMakeLen0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for i := 0; i < 10; i++ {
			s = append(s, i)
		}
	}
}
```

len=10で初期化
```go
func BenchmarkInitMakeLen10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 10)
		for i := 0; i < 10; i++ {
			s[i] = i
		}
	}
}
```

len=0,cap=10で初期化
```go
func BenchmarkInitMakeLen0Cap10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 10)
		for i := 0; i < 10; i++ {
			s = append(s, i)
		}
	}
}
```

len=10,cap=10で初期化
```go
func BenchmarkInitMakeLen10Cap10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 10, 10)
		for i := 0; i < 10; i++ {
			s[i] = i
		}
	}
}
```

### 結果
```sh
% go test -bench . -benchmem
BenchmarkInitDeclation-4                 3000000               433 ns/op              80 B/op          1 allocs/op
BenchmarkInitMakeLen0-4                 10000000               284 ns/op             248 B/op          5 allocs/op
BenchmarkInitMakeLen10-4                200000000                7.43 ns/op            0 B/op          0 allocs/op
BenchmarkInitMakeLen0Cap10-4            100000000               15.4 ns/op             0 B/op          0 allocs/op
BenchmarkInitMakeLen10Cap10-4           200000000                7.19 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/midorigreen/learn/201710/slice       118.526s
```

#### 考察

## 参考
- [Go のスライスの内部実装](http://jxck.hatenablog.com/entry/golang-slice-internals)
- [golang でパフォーマンスチューニングする際に気を付けるべきこと](https://mattn.kaoriya.net/software/lang/go/20161019124907.htm)