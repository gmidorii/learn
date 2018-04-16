# Golang Tokyo

## Goroutine
### 基礎
- CPUが複数あれば並列処理される
- コールスタック
  - 呼び出した時点でコールスタックが分離する
    - goで呼び出した先のpanicは分岐後のdeferでリカバーする必要がある
  - panicはgoroutineのコールスタックへ戻っていく
- channel
  - 型がある(R/W, RO, WO)
  - close後に書き込むとpanicする
  - RO
    - `chan←`
  - WO
    - `←chan`
- select
  - 1つのcaseのみが実行される

### 注意点
- Wait処理
  - chan(struct{}{})
  - sync.WaitGroupがベスト
    - errgroupを利用するのも有り
- 競合状態
  - 各goroutine間で同一変数を扱うと競合してしまう
  - sync.Mutex()を利用して排他制御する
  - sync.RWMutex()で読み書きを別に排他制御できる
  - 数値系sync.atomicパッケージを利用
- Goroutine leak
  - 片方のchannelだけを利用すると書き込みロックがかかってリークする

### ツール
- gotrace
  - Web上で可視化できるツール
  - runtime/trace
  - dockerを利用している
    - カスタムされたGoのランタイムパッケージを利用する必要があるため
    - Image: divan/golang:gotrace

### 使いどこ
- sliceの各値に対して処理を実行したい時
- APIコール
  - 片方のGoroutineでキャッシュ
  - もう片方でレスポンス処理


------------------------------------------------

## チャネルの仕組み
- GopherConの発表が元ネタ
- ヒープ領域にバッファのデータが配置される

### バッファ有りチャネルの基本
- 重要(hchan)
  - buf
  - lock
  - sendx(送信用index)
  - recivex(受取用index)
- 基本的にまずlockをとって処理を行っている
- goroutineスケジューリング
  - M:N threadモデル
  - M個のOSスレッドがN個のgoroutineを扱う
- chanがいっぱいになった場合、OSスレッドがgoroutineをwaitする
- sudog: LinkedListへgoroutineを追加する
  - 値とgoroutineを持っている
  - sudogからバッファを経由せずに値を書き込む処理を行う
  - channelが持っている
- selectは各channelのsudogを掴んでいる

-------------------------------------

## vgo
- goに持たせるパッケージ管理機能
  - depではないのね..
- Go1.12で正式導入
- 依存関係の解決
  - ビルド時に行う
  - ソースコードのダウンロードも行う
  - go.modファイルで管理
- vendor不要
  - $GOPATH/src/v 配下にパッケージが配置される
- Import Compatibility Rule
  - 破壊的な変更をする場合はimport pathを変える
- Minimal Version Selection
  - 可能な限り過去のバージョンを取りに行く

-------------------------------------

## linter
- 静的解析
  - go配下のパッケージを利用する
- 行番号を管理しないといけない
