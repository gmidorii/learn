# GoCon 2017/03/25

## Keynote
- Container need
    - Encapsulation
    - Isolation
    - Lightwight
- LetsEncrypt: /autocert package
- net/http Server graceful shutdown


## 条件式評価機~管理ツール
@tenntenn

### Overview
- 管理ツール抽象化
    - 汎用的に使えるようにする
    - 難易度が高い
- コアな機能を式で書けるように
- 条件式評価器
    - go/parserで式のASTを取得
    - constant.Valueを利用
- JSON Scheme
- Goでは式の評価をすることができる


## Concurrency
@niconegoto

### Overview
- プロセス間通信IPCの必要性
- Concurrency model
    - shared-memory
    - message-passing
- goroutine
    - 2048byte
    - スレッドのようなもの
    - M:Nスケジューリング
    - スケジューラーに追加するところをよしなにしてくれる
- channel
- select
    - timeoutや定期実行なども利用できる


## Context
maki さん
@lestrrat

### Overview
- 以前の方式
    - 安全な設計にするためにロックをかけないといけない
```go
obj.Start()
obj.Stop()
```
- mutexを使わないような実装を心がける
- context.Context
    - Explicit cancellation
- Run(ctx)として定義する
    - contextを引数として渡す設計
- blockする処理はcontextを渡す
- contextは第一引数で渡すのがお約束
- contextをキャンセル以上に使うのは難しいところ
- migrateする際は、+buildタグできる    
- 下記がパターンになる
```go
func Loop(ctx context.Context) {
    for {
        select {
        case: <-ctx.Done():
            return
        default:
        }

        // 処理
    }
}
```


## 分散オブジェクトストレージ
Goto san
@ono_matope

### Overview
- Dragon    
    - 分散オブジェクトストレージ
    - S3みたいなもの
    - PUTでとりこみ
    - cassandraをメタDBとして利用している
- 耐障害性
    - Circuit Breaker
        - 接続先ごとにCircuit Breakerを用意
        - DialContextが監視する
        - 数回上手く行かなかったら対象から外す
    - Timeout for Stream
        - io.Copy
            - read,write処理を内部で行っている
        - read,write処理のどちらがブロックしても動きが止まる
        - net.conn.SetDeadline()
        - http.TimeoutHandler
        - ○ タイムアウト付きのio.Reader,io.Writerを定義


## Testable HTTP API Server
@_achiku

### Overview
- routeing
    - gorilla/mux
- Model: SQLAlchemy + Alembic + dgw
- advanced testing with go (slide)
- 一度APIアクセスしてファイル保存
    - リポジトリにコミット
    - 次回以降はファイルを使う
    - テスト時はServer, Clientをテスト内で生成する


## GAE/DDD
@__timakin__

### Overview
- GAE/Goで詰まった点
    - フォルダ構成
        - ちゃんと方式に習う
        - vendoring配下でパッケージ名衝突
    - 秘密情報の管理
        - ConfigurationRepository
        - Datastoreに情報を入れておく
    - Datastoreに移行
        - CloudSQLはちょっとお金がかかる
- DDDはパッケージ名の衝突を避けられてよい
- Goの実装と相性が良いのではないか


## 汎用言語処理系
@Lind_jp

### Overview
- GoCaml言語を実装
- コンパイラ
    - `徐々に抽象度を落とした実装` を行う
- Lexerを参考に


## net/http package
@kaneshin0120

### Overview
- 初心者から中級者へ
- GoDocを読みましょう
- ResponseWriter
    - Header -> WriteHeader -> Writeの順
- DefaultClientを使う場合はinitで設定する
```go
http.DefaultClient.Timeout = 30 
```
- net/http.Constant積極的に利用する
- BodyはバッファになっているのでフラッシュしてCloseする
- http.Clientをsettableにする


## GoImagick
よやさん @

### Overview
- import対象が各バージョンにより異なる
- 基本的にCを直接呼ぶ形式 
- net/http内で/を縮める処理が入る
- Reverse proxyでガードしないと危険


## LT

### net/http context
- WithValue
    - データがひっつけれられる
    - mapではないlist
- 前処理の結果をContextの中に突っ込む
- リクエストに関する情報を入れるべき
- withValueにはint,string等は怒られる(golint)
    - 自前で型を定義して渡す必要あり

### goa
- goa: マイクロサービス向けフレームワーク
- DSL -> Swaggerを生成できる
- 再生成すると差分が出る
- plugin Generators
- 別ファイルに分割して生成する
- plugin作って対応する

### net/http timeout Go1.8
- 1.8になって細かくtimeoutを設定できるようになった 
- Go1.8になってhttp serverが順分安定してきている

### cgo cross compile

### Go contribution

### デスクトップアプリ

### dep管理
- 2018には入りそう
- `rm -rf vendor/` を消さないとupdateできない
- `dep ensure`

### インフラ
- Openconfig
    - オープンソース
    - YANG
- リアルタイムストリーミング方式
    - ネットワーク側から随時変更を通知する
- ネットワークインフラを効率的に管理するAPIが必須
- Goと親和性の高い技術が使われている

### カンスト

