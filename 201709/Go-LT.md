# Go-LT

## Go introduction

### Goの成り立ち
- 誰が
  - Googleのエンジニア
    - ロブ・パイク
    - ケン・トンプソン
    - etc
- なぜ
  - コンピュータ世界の変化(ここ10年)
    - コンピュータ自体の処理速度の向上
    - マルチコアコンピュータの出現
    - 依存関係解決による遅いコンパイル
    - 型システムへの不満→動的型付け言語の人気
- どのような目的で
  - ゴール
    - コンパイルが早く
    - 並列処理
    - ガベージコレクション
  - 目指す形
    - インタプリタ言語の容易さと動的型言語の効率、コンパイル言語の静的な方の安全性を併せ持つこと
  - ポイント
    - 高速コンパイル(数秒)
      - 依存関係の解析を容易化
    - 静的な型
      - 型間に階層構造はなし
    - 完全なガベージコレクション
      - 並列処理と通信を公式サポート
- 設計方針
  - typeとtyping量を減らした → ちょっと疑問
  - 型に階層を作らない
    - 型同士の関係性を記述する必要はなし
  - コンセプトに一貫性をもたせる

### Goの基本
- 言語仕様
  - 静的型付言語
    - 型推論あり (:=)
  - 特殊構造
    - array, slice, map, struct, interface, func
    - array=固定長配列
    - slice=可変長配列 (= 実際はarrayへの参照)
    - map= key, value
    - func= 関数
    - struct= 構造体
      - レシーバー(= メソッド)を持たせられる
    - interface 関数宣言
- 文法
  - `var` で変数宣言
  - `const` で定数宣言
  - `&` でアドレス取得
  - `*` でアドレス→変数の中身取得
  - 大文字始まりでエクスポート
    - 関数
    - struct
  - if
    - `if x > 10 {}`
  - for
    - `for i := 0; i < 10; i++ {}`
    - `for index, value := range [array, slice, string] {}`
    - `for key, value := range [map] {}`
  - switch
    - `switch hoge: case: default`
    - 各caseで自動でbreak
  - goroutine
      - `go hoge()` のように宣言することで軽量スレッド作成
      - `channel` を作成することでスレッド間でデータをやり取りする
      - `select` を利用
      - 生成コストは数キロバイトのメモリのみ
        - セグメントスタックをもつ
        - スタック拡張時には、3つのインストラクションを実行するのみ
        - 同一アドレス空間で10万レベルのgoroutineを生成可能

### Go言語にないもの
- Generics
  - よい実現方法を模索中
- exception
  - exceptionによりコードが複雑になることをさけるため
  - 代わりに `hoge, err := createHoge()`
    - `if err != nil {}`
- オブジェクト指向
  - 型階層がない
  - インターフェイスも思想が異なる
- assert
- オーバーロード
  - メソッドのシグネチャを調べるコスト
  - メソッドが名前だけで検索可能にするようシンプル化

### Go仕様
- 動的メソッド呼び出し
  - インターフェイス経由のみ
- インターフェイス (duck typing)
  - 型がある動作を満たせば、インターフェイスとみなす
    - 関数名とシグネチャが、インタフェースが持つそれらと完全一致するか?
  - 型を満たすかはコンパイル時に判断される
- nilに型が存在
  - 型と値をもつ
    - nil interface = (nil, nil)
    - 型付きのinterface = (MyError, nil)

### 学ぶには
- GOPATHの設定
    - 例: `export GOPATH=~/dev/src/`
- [Tour of Go](https://go-tour-jp.appspot.com/welcome/1)
- [The Go Playground](https://play.golang.org/)
  - ブラウザ上でGoを実行可能
  - リンク化可能
- [GoDoc(ドキュメント)](https://godoc.org/)
- [初心者が見ると幸せになれる場所](https://qiita.com/tenntenn/items/0e33a4959250d1a55045)

### 特徴的な部分
- buildしたバイナリ一つで実行可能
  - 各OS/アーキテクチャごとのバイナリを生成可能
    - 例: `GOOS=linux GOARCH=amd64 go build`
  - コンパイル速度が早い
- goroutineによりAPIサーバーが数行で立てられる
- 魔法が使えない
  - 愚直にコードを書くしかない
  - シンプルで読みやすさ抜群 (標準ライブラリのコードも普通に読める)
- GAE/Go
  - Google App Engineと合わせて利用されることが多い
    - 簡単に導入できるため

### Tips
- go getでライブラリ導入
  - `go get github.com/garyburd/redigo/redis`
- ライブラリの依存関係
  - [glidle](https://github.com/Masterminds/glide)
    - /vendor配下にライブラリをdownload
  - [dep](https://github.com/golang/dep)
    - 準公式
      - 後々 `go dep` のような形で公式に入る想定
- フォーマットが公式で決められている
  - `go fmt` で自動フォーマット (保存に自動で実施)
- auto import可能
  - `go import` でオートインポート + フォーマット (保存に自動実行)

### エディタ
- vscode
  - `Go` plugin導入でだいたい動く
  - `delve` を導入して、デバッグ実行可能
- Gogland
  - JetBrains製 IDE
  - あんまり使ってない(IDEが必要なほどかな..?)
- vim
  - `vim-go` で全てうまくいく
    - auto suggestion
    - build, run

### 導入事例
- ツール
  - Docker
  - kubernetes
  - [terraform](https://github.com/hashicorp/terraform)
  - [prometheus](https://github.com/prometheus/prometheus)
- 会社
  - Google
    - [vitess](https://github.com/youtube/vitess)
      - YouTubeのMySQLインフラで利用
  - Uber
  - メルカリ
    - [Push配信ミドルウェア](https://github.com/mercari/gaurun)
    - 
  - DeNA
  - エウレカ

### 界隈で有名な人
- @tenntenn
- @deeete
- @mattn
- 渋川さん
  - Real World Http著者
  - [Goならわかるシステムプログラミング連載](http://ascii.jp/elem/000/001/235/1235262/)
- 牧さん
  - [peco](https://github.com/peco/peco)
    - ※ peco + qhqにお世話になってます

### 参考
- http://tech.innova-jp.com/how-to-use-golang/
- https://mayonez.jp/topic/1504
- https://www.slideshare.net/yujiotani16/go-49082403
- https://qiita.com/Jxck_/items/02185f51162e92759ebe
- https://github.com/mitchellh/gox
- http://go-talks.appspot.com/github.com/lestrrat/go-slides/2014-yapcasia-go-for-perl-mongers/main.slide
