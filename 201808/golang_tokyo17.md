# golang.tokyo 
2018/08/21

## テーマ
テスト

## tour of testing

### メモ
* 命名
  * TestStruct_Method
* テストデータ
  * testdataディレクトリとしておけばpkgとして認識されない
  * .golden拡張子 ※ 調べる
    * https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3
* go vetが自動で実行されている(Go 1.10以降)
* デバッグ
  * dlv test -- -test.run
* CI
  * JUnit形式に変換する
* TB
  * Errorは処理は続く
  * Fatalは処理がお亡くなりになる
  * Helper
    * tb.Helper()を呼んでおくと、呼び出し元のファイル等がログに出る
* Table Driven Test
  * t.Run()
    * 成否を各テストケースで見せれる
  * t.Parallele()
    * ローカル変数を捕捉する必要あり
* 変数名
  * want/got
* %#v => は\nが出力される

## 非公開なxx
tenntenn

### メモ
* 外部パッケージの割合は4割くらい
* 外部パッケージメリット
  * 使いづらい部分が早期に発見できる
* メソッド式
  * (*Counter).reset
* 非公開な型
  * type Exporttype = hoge

## 外部パッケージへの依存をテスト
@duck8823

### メモ
* duci
  * CIサーバー
* http
  * go-chiでcontextを渡してテスト
* gock
  * http requestを奪ってくれる
  * 特定のURL

## developer frindly testing

### メモ
* テストは落ちたときに直しやすいようにしたい
* 完走させる
  * panicにならないテストを書く
    * panicだと、どのテストで落ちたかわからなくなる
* deepEqual
  * go-cmp/cmp
* Golden File test
  * ファイルを作っておいて、そのファイルと比較する

## 止めたいのに止められないテストの話
@knsh14

### メモ
* go test -failfast
  * go 1.11でバグ修正