# みんなのGo言語

## テスト
### Table Driven Test
- 特になし

### reflect.DeepEqual
- mapやsturctの比較で利用できる
```go
type Hoge struct {
  A int
  B string
}

func TestHoge(t *testing.T) {
  h1 := Hoge {
    A: 1,
    B: "string",
  }
  h2 := Hoge {
    A: 1,
    B: "string",
  }

  if !reflect.DeepEqual(h1, h2) {
    t.Errorf("want %#v got %#v", h1, h2)
  }
}
```

### Race Detector
- 並列処理の際に競合状態が発生しないかチェック
- 競合状態
  - 1つのデータを複数の並行コードが読み書き
  - イレギュラーな値になること
- 使用方法
```sh
go run -race pkg
go build -race pkg
go install -race pkg
go get -race pkg
```
- 実行時に競合状態を検出→ `Data Race` を報告
- Race Detector有効時
  - メモリ5~10倍
  - 実行時間2~20倍
- ライブラリ実装時にはスレッドセーフな実装を心がける

### TestMain
- 書き方
```go
func TestMain(m *testing.M) {
  setup() // 任意の名称でOK
  exitCode := m.Run()
  shutdown() // 任意の名称でOK
  os.Exit(exitCode)
}
```
- 各テスト実行時に、自動でsetup/shutdownが走る
- 各テストは、 `m.Run()`

### Build Constraints
- build tag
```go
// +build integration
package main
```
```sh
# integrationがついたテストも実行される
go test -tags=integration

# vetによってbuild tagの検証ができる
go tool vet -buildtags main.go
```