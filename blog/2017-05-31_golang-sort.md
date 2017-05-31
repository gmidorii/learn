# GolangのSort処理について

## GolangのSort処理
GolangのSort処理について、まとめました。  
(`// package sort` と記載があるサンプルコードは、Golang本体のソースコードです)

## Sort
### Sort Interface
Golangでは、`struct`のソートを行うため、`sort.Interface`を実装する必要があります。  
(実際はGo1.8以降、下記の `sort.Slice()` を利用して、ソートすることができるようになりました。)  
  
```golang
// sort.Interface
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```
| メソッド                | 特徴              |
|---------------------|-----------------|
| Len() int           | 配列・リストの要素数      |
| Less(i, j int) bool | i < j を満たすかの真偽値 |
| Swap(i,j int)       | iとjの要素を入れ替える    |

### Sortの方法
標準パッケージで対応している型は下記になります。
- float64
- int
- string

また、ソートの方法をintをもとにみていきます。

#### sort.Ints(a int[])
`Ints()` メソッドは、int[]型をソートするメソッドです。  
これは、昇順(increacing)にソートされます。  
```golang
nums := []int{4, 3, 2, 10, 8}
sort.Ints(nums)
fmt.Println(nums)

// output: [2 3 4 8 10]
```

内部では、 `IntSlice` 型のstructにつめかえを行い、Sort関数を呼んでいます。  
(`sort.Interface`を実装するルールは同じ)
```golang
// package sort

// Ints sorts a slice of ints in increasing order.
func Ints(a []int) { Sort(IntSlice(a)) }


// IntSlice
// Len,Less,Swapを実装したstruct

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
```

#### sort.Reverse(data Interface) Inteface
`Reverse(data Interface)` は、sort.Interface型のソート方法を逆順にする修正する関数です。  
実際に内部で行っていることは、下記の２点です。  
1. reverse structにstructを詰め替える  
`reverse struct` は、 `sort.Interface` を埋め込んでいます。  
```golang
// package sort

type reverse struct {
	// This embedded Interface permits Reverse to use the methods of
	// another Interface implementation.
	Interface
}

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
	return &reverse{data}
}
```

2. Lessメソッドの引数をスワップ  
`reverse struct` は、実装を参照したところ、
`Less(i,j int) bool` のみを再実装しています。  
これにより、比較条件を逆にすることで、逆順ソートを実現しています。
```golang
// Less returns the opposite of the embedded implementation's Less method.
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}
```

#### sort.Search(n int, f func(int) bool) int
nの長さの配列内で`f()=true` となる最小のindexを返却します。  
→ 配列内で指定のvalueがあるかどうかを調べる関数です。  

特徴
- 条件を満たさなかった場合、nが返却
- 一致を探すことも、ある数値以上の値を探すことも可能
- binary search(二分探索)で検索
  - ソートされた配列に対して実行しないと正確な結果が返却されない
- 下記は`Search`のラッパー関数
  - func SearchFloat64s(a []float64, x float64) int
  - func SearchInts(a []int, x int) int
  - func SearchStrings(a []string, x string) int

```golang
// package sort

// binary sort の実装
func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

// Int ラッパー関数
func SearchInts(a []int, x int) int {
	return Search(len(a), func(i int) bool { return a[i] >= x })
}
```

利用サンプル
```golang

nums := []int{4, 3, 2, 10, 8}
x := 8

// 事前にソートしておく
sort.Ints(nums)
// search asc
index = sort.Search(len(nums), func(i int) bool {
  return nums[i] >= x
})

if index < len(nums) && nums[index] == x {
	// nums[index] = 8 = x
} else {
	// x は nums[]中に存在しなかった
  // だが、indexはnums[]の中にxをインサートする位置にある
}
```

- sort.Slice(slice interface{}, less func(i, j int) bool)
  - less関数で定義された順にソートされる(go1.8.1以上)
  - sort.Interfaceを実装していなくてもソートが可能となっている
  - stable sortを保証していない
    - sort.SliceStableを利用することで保証される

- sort.Sort(data sort.Interface) は渡されるデータによってソートの種別変える
  - ヒープソート
    - 条件: log(n+1)の整数値の2乗よりも深い場合
      - 右に1ビットずつシフト演算して0になるまでの回数の2乗が最大の深さ
      - この値が0となったときに、ヒープソートに切り替わる
    ```go
    // Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
    ```
  - シェルソート
    - 条件: スライスが12より短い場合
    ```go
    // Use ShellSort for slices <= 12 elements
    ```
  - クイックソート
    - 条件: 上記条件に当てはまらない場合