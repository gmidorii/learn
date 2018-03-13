# Goならわかるシステムプログラミング

## 第2章 低レベルアクセスへの入り口 Writer
- POSIX系OSではすべてのものをファイルとして抽象化
  - ファイルディスクリプタが割り当てられる
  - ファイルディスクリプタは数値
- プロセス
  - 生成時に3つfdを作成
    - 標準入力0
    - 標準出力1
    - 標準エラー出力2
  - その後ファイルをopenするたびに1ずつ増える
- io.Writerをラップするデコレータ
  - フィルタ的な役割を持たせられる
  - gzip.NewWriter(), bufio.NewWriter()
  - 標準入出力時に変換処理 or 追加処理をかましたい場合に利用できそう
- `godoc -http ':6060' -analysis type`
  - godoc起動 + implementsチェック

## 第3章 低レベルアクセスへの入り口 Reader
```go
// io.Closerインターフェースを満たしていないio.Readerを
// io.ReadCloserインターフェースへキャストできる
// Close()時は何もしない
var reader io.Reader = strings.NewReader("テスト")
var readCloser io.ReadCloser = ioutil.NopCloser(reader)
```

- エンディアン
  - バイトの格納順
  - リトルエンディアン
    - 小さい桁から格納
    - CPU
  - ビッグエンディアン
    - 大きい桁から格納
    - ネットワーク
- テキスト読み込み
  - bufio.Reader('\n') は区切り文字も含む
  - bufio.Scanner() は区切り文字は含まない
- zip.NewWriter(io.Writer)でwriterを作成
  - zipWriterはio.Writer IFを実装していない
  - `w, e := writer.Create("tmp.txt")` でarchiveするファイル作成
  - io.Copyでarchive writerへreaderから引き渡す

## 第4章 チャネル
- queue + 並列機構 = channel
- 性質
  - ① データを順序よく受け渡すデータ構造
    - ランダムアクセス不可
  - ② 並列処理の場合でも正しく動作する
    - 整合性が壊れない
    - 同時に投入の場合も、1つのgoroutineからしか順に投入となる
  - ③ 読み書きの準備ができるまでブロックする
    - データが無い場合、データが投入されるまでブロック
    - バッファに空きがない場合、空きが出るまでブロック
- channelは1, 2 returnできる
  - channel内を通ったデータ + close判定
- 終了通知はcloseを利用するのでなく別channelを開く
  - 各channelはcloseしない (closeせずともGCされる)
  - 終了通知用のchannelはcloseする