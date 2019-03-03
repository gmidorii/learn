# golang.tokyo 6 (2017/06/01)

## 概要
[golang.tokyo #6](https://golangtokyo.connpass.com/event/57168/?utm_campaign=event_reminder&utm_source=notifications&utm_medium=email&utm_content=detail_btn) に参加しましたので、レポートを記載いたします。  
(会場は、DeNAさんでお寿司(食べ損ねた)とお酒をご用意いただいておりました。)  
下記は各セッションのまとめです。

## Gopher Fest 2017
[@tenntenn](https://twitter.com/tenntenn) さん  
  
スライド  
[https://www.slideshare.net/takuyaueda967/gopher-fest-2017:embed]

### 概要
サンフランシスコで開催されるGoSFが主催のイベントの参加レポートです。
セッションの中のひとつ [The state of Go](https://talks.golang.org/2017/state-of-go-may.slide) のお話

### The state of Go
Go1.9での仕様変更や標準ライブラリの改善についてのお話
(余談ですが、[tip.golang.org](https://tip.golang.org/)で最新のmasterのドキュメントが見れるそうです)  

- Go1.9は5/1にコードフリーズ済み
  - 残りはバグフィックスのみ
- リファクタリングを安全にする方法
  - [Codebase Refactoring](https://talks.golang.org/2016/refactor.article)
- Goで安全にリファクタできるのか?
  - 定数の場合は問題なし
  - 関数はラップして引き継げば問題なし
  - 型の場合→ 問題有り
    - 型を作る → メソッドが引き継げない
    - 埋め込み → キャストができない
    - Go1.9でAlias導入
- 言語仕様変更
  - Alias を作成できる
    - 型のエイリアスを定義できる
    - こんな感じで `type A = http.Client`
    - 完全に同じ型
    - キャスト不要
    - エイリアスの方ではメソッドを定義できない
- 標準ライブラリ変更
  - math/bits ビット演算に便利なライブラリ
  - `syc.Map` スレッドセーフなマップ
  - `html.template`がより安全に利用可能に
    - errorを返却するようになる
  - os.Exec
    - 環境変数が重複の際、一番後ろのものを優先する
      - ソース中で上書き可能
- go testの改善
  - *vendorを無視するようになった*

### 感想
Go1.9にあたっての改善点のお話がメインとなりました。  
@tenntennさんもおっしゃってましたが、 `go test`で`vendor`配下を  
みなくなる修正はなかなか良いですね。


## 初めてGolangで大規模Microservicesを作り得た教訓
Yuichi Murata さん  
  
スライド未公開(2017/06/01)

### 概要
GAE/Goで大規模なMicroservices(30くらい)を作った際の、教訓のお話  

- いきなりつかってみたため、失敗談が中心
- 構成
  - GAE/Go
  - Gin/Echo
  - Microservices構成
- 結論
  - 困ったときはGoの哲学にそったシンプルなアプローチを取る
  - Goを過信せずパフォーマンスに気を配る

### 教訓
#### 1. フレームワークにこだわらない
- Railsの経験から良いフレームワーク探しをこだわって行った
- [Gin](https://github.com/gin-gonic/gin)
  - 良い点
    - シンプル
    - App Engine とのインテグレーション有り
    - Framework Context vs App Engine Context
      - gin.ContextからApp Engine Contextに派生させられることで解消
  - 困った点
    - ワイルドカードパスを利用したルーティング
- Echo
  - 良かった点
    - ワイルドカードパス利用できる
    - コンポーネント開発が捗る
  - 困った点
    - 開発が盛んなことにより設計が書き換わる
      - バージョンを固定することで対応(アップデートできない..)
- まとめ
  - Goでは、net/httpパッケージで十分な部分もある
    - 足りない部分はライブラリを利用
  - フレームワークを利用しなくても統一的なコードとなる
    - 統一的なコードを書くためのフレームワークだったが、Goなら標準で統一的となる

#### 2. Interfaceを尊重する
- 独自のエラー型を定義
  - こちらの型にerrorを寄せる
    - 毎回キャストするのが面倒のため
- Nilに型がある問題発生
  - (参考記事) [絶対ハマる、不思議なnil](http://qiita.com/umisama/items/e215d49138e949d7f805)
  - 独自のエラー型を混在させた際に発生
    - 独自型とerrorインタフェースでnilの型が異なる
- まとめ
  - Interfaceが定義されたものはそのまま利用するほうが自然に実装できる

#### 3. regex compile/ reflectionが遅い
- バリデーション
  - JSON Schemaを利用
- GolangでのJSON Schema
  - [gojsonschema](https://github.com/xeipuuv/gojsonschema)
  - [go-jsval](https://github.com/lestrrat/go-jsval) (たしか..)
- パフォーマンステストで問題発覚
  - バリデーションの有無でパフォーマンスに顕著な差が出る
  - [pprof](https://golang.org/pkg/net/http/pprof/) を利用して実行を解析
- 問題点の解消
  - Regexコンパイルを毎回行う実装
    - キャッシュすることで解消
  - reflectionを多用していた
    - 利用しないように修正
  - go-jsvalを利用すればさらに早くなった
- まとめ
  - regex/reflectionのコスト高を意識する
  - regex compile / reflectionはとにかく重い

#### 4. 非対称暗号は遅い
- 認証・認可処理
  - 非対称鍵の処理が重い
    - opensslと比べcypto/rsaが貧弱
      - 一桁程度の差が生じる
    - cgoを利用したopensslのバインディングも存在
      - GAEでは利用不可
- 荒業で解消
  - 非対称鍵の処理署名処理をPHPにして外出し
  - httpリクエストのほうが早い

### 感想
GAE/Goでサービスを作る際の問題点がいくつか聞けました。  
サービスでの利用はこれからあるかもなので、参考になりました。  
またnilに型がある問題は、まだあたったことがないので知れて良かったです。


## [LT]ゲーム開発に欠かせないあれをしゅっとみる
ゴリラのアイコンの方([@Konboi](https://twitter.com/Konboi))  
  
スライド
<script async class="speakerdeck-embed" data-id="94b7f3f58efa43ed8b31e5c2e4219794" data-ratio="1.33333333333333" src="//speakerdeck.com/assets/embed.js"></script>

### 概要
- CSVの話
  - スマホゲームには欠かせない
- カヤックではデータの受け渡しとしてCSVを利用している
  - マスターデータ等
- CSV困る時
  - カラムとデータの関連性が見づらい
  - 空欄がつらい
  - DBにインポートする系はその際にわかる
  - 調査系で利用する時つらい
- [csviewer](https://github.com/Konboi/csviewer)
  - csvをDBのテーブルを検索するように参照できる
- ライブラリ
  - [sliceflag](https://github.com/soh335/sliceflag)
    - 同一オプションで複数の値を受け取りたい
    - 普段のflagの利用方法と変わらない
  - [tablewriter](https://github.com/olekukonko/tablewriter)
    - データをいい感じにテーブル表示してくれる


## [LT] Go Review Commentを翻訳した話
[@knsh14](https://twitter.com/knsh14) さん

スライド  
<script async class="speakerdeck-embed" data-id="1a12387747f3445f9bbd9bf7f8b0995c" data-ratio="1.77777777777778" src="//speakerdeck.com/assets/embed.js"></script>

### 概要
Go Code Review Commentを読もう  
(→ 読みます)

- [Go Code Review Comment](https://github.com/golang/go/wiki/CodeReviewComments)とは
  - コードレビューをする際に見るべき箇所をまとめたもの
  - Effective Goの簡略版
- 翻訳してみました([qitta](http://qiita.com/knsh14/items/8b73b31822c109d4c497))
- 良かった点
  - 正解のパターンを勉強できるところが効率が良い
- 内容
  - コードの見た目を改善
    - ツールでなんとかしよう
    - golintは優秀
  - コメントや文章の体裁
  - Tips系
    - より良いパターンを記載
  - 設計の指針
- まとめ
  - 良いドキュメントの翻訳は英語の勉強に最適


## [LT] ScalaからGo
James さん  
  
スライド未公開(2017/06/01)  

### 概要
- 関数型開発はGoでできるか?
  - できないですね
  - ジェネリクスがないので..
- 関数型開発のコンセプトは利用できるか？
  - 利用できる!!
- 関数型開発とは?
  - 副作用がない開発
  - 副作用あり
    - テストしにくい、バグの原因
  - 部分適用
- Goは初心者が入りやすく会社での導入がスムーズに行える点が良い


## [LT] Crypto in Go

### 概要
- Goにおける暗号アルゴリズムの利用
- AES
  - 固定長でしか利用できない
- AES + Padding + HMac
  - 行数があって面倒
- AES + GCM
  - 機密モード + 認証モードがひとつに
  - とってもシンプル


## まとめ
各セッションとLTについて、ざっくりと記載いたしました。  
詳細知りたい際は、スライドが公開されるはずですのでそちらをご参照ください。  
有意義な発表が聞けて楽しかったです。  
開催ありがとうございました。(次回も参加したいです)