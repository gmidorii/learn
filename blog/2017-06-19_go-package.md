# Gopackage

## Golangのパッケージ構成について
golangで開発するにあたって、パッケージ構成がよくわからなかったため、調べてみました。  

当記事では、下記に焦点をあてて調べてみました。
- packageの分け方
- pacakge内でのファイルの分割
- ファイル内のコードの書き方

## パッケージ構成
### Good Practice
下記の4点を考慮したアプローチ

1. Root package is for domain types
2. Group subpackages by dependency
3. Use a shared mock subpackage
4. Main package ties together dependencies

#### #1. Root package is for domain types
ドメインとは、データとプロセスがどのように相互作用を起こすかを記述する、高次元の言語のことです。  
ドメインは、技術的なバックグラウンドに依存することはありません。  
  
Root packageには、このドメインタイプを配置します。Root package内は、  
simpleな`struct`と`stuct`の振る舞いを定義した`interface`のみで構成されます。  
  
`The root package should not depend on any other package in your application!`  
→ Root packageはアプリケーション上の他のいかなるパッケージにも依存すべきでない

`ドメイン = User` の場合
```go
package app

type User struct {
  ID   int
  Name string
}

type UserService interface {
  User(id int) (*User, error)
  UpdateUser(id int, name string) error
}
```

#### #2. 


### (補足) イマイチなアプローチ
#### #1 モノリシックなパッケージ
全てのコードをひとつのパッケージに詰め込む方式  
→ 小さなアプリケーションに対しては十分にうまくいきます。

#### #2 機能(function)ごとにグルーピング
機能ごとに、処理をグルーピングすることで構成する方式  

構成例
- handler : ハンドラを集めたパッケージ
- model : モデルを集めたパッケージ
- controller : コントローラを集めたパッケージ

ただし、この手法には欠点が２点ほど存在しています。
1. 命名  
`handler.XXXHandler` のように名前が重複してしまう。
2. 循環参照  
パッケージ同士で参照しあい、循環参照(circular dependencies)を起こしてしまう。  
(1方向のみの依存関係をもつ場合のみしか、有効ではない手法)

#### #3 モジュールごとのグルーピング
モジュールごとにグルーピングをする方式  

こちらも同様に下記が欠点です
1. 命名  
`users.User` のように名前が重複してしまう
2. 循環参照
## 参考
[Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)  
[【翻訳】【Golang】標準的なパッケージのレイアウト](http://allishackedoff.hatenablog.com/entry/2016/08/23/015016)