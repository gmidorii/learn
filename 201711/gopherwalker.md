# GopherWalker

## Go Friday
### switchのススメ
- iotaは値に意味がない場合のみ利用する
- caseの値が重複しているとコンパイルエラー

### subtest and table driven test
- `t.Run(string, func(t *testing.T) {})`
- table driven testの各実行をサブテストで行うことで2点の利益
	- エラー時にどのdataで落ちたかがわかる
	- 落ちたテストのみを再実行可能

### Test Helper
- helper func内で落とす(Fatal)
- t `*testing.T` を引数に取る
- t.Helper() をhelper funcの先頭で宣言
	- エラーが起きた場合は、helper func内の行が示される

### Temporary Error Handling
```go
type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}
```
- 条件
	- temporary interfaceを実装しているか
	- te.Temporary()がtrueを返すか
- error interfaceを満たすように実装する
- interfaceを実装した型はpackage内に隠匿すると良い
	- 利用側が具体的な実相を知ることなく使える
	- 変更時に利用側への影響がない
- エラーレベルの判定に利用可能

## 認証	

## Protocol Buffers
### What?
- バイナリフォーマット
	- encode/decodeが高速
- protoファイルを元にclient/serverの実装ができる
	- コンパイル時に不整合を見つけられる
- 仕様
	- 通信プロトコル
		- HTTP/HTTP2
	- API実行ルール
		- gRPC
		- JSON-RPC
		- REST
	- 構造化データ仕様
		- XML
		- JSON
		- Protocol Buffers

## 新しい型定義
- メリット
	- 新しい型になる
		- 明示的に型の利用を実施する
		- 可読性アップ
	- メソッドを実装できる
- 値に関連する振る舞いを値自身に持たせることができる
	- string -> UserName
- interfaceの実装
- 基本型に具体性を持たせられる
	- 具体型を生成するfuncを作ると良い
-	sliceに対する振る舞い実装
	- ex) []User -> Users 
		- `func (u Users) FilterAdmin() Users {}`
- First Class Collection
	- sliceをラップして新しい型を定義すること(↑ だね)
	- sliceとsliceに対する振る舞いをセットで管理できる
	- sliceに名前をつけることができる



