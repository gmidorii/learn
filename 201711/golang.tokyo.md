# Golang.Tokyo 10

## テキスト審査 (Eureka)
### 登壇者
後藤さん

### 概要
- 業者対策の話
- はじく対象
	- 写真
	- 投稿
- ワードフィルタ
	- 形態素解析による単語一致
	- ブラックワードリストを作っている
- go-jp-text-ripper
	- https://github.com/evalphobia/go-jp-text-ripper
	- kagome利用
		- https://github.com/ikawaha/kagome
- 文章をスコアリングして、上限値を上回るとNG
- 特徴的なワードに関して性別ごとに登録
- 怪しいテキスト -> slack通知
- 3ヶ月スパンで既存フィルタが意味をなさなくなる
- 機械学習
	- 特徴量は、手作業で登録もある
	- エラーになるデータが少ないため、学習データが足りない
- 怪しいものは、一旦人間がフィルタする


## interfaceとの付き合い方
### 登壇者

### 概要
- 自由なjson構造をパースしたい 
	- `{"1": ["hoge", "hoge", 1]}`
	- map[string]interface 型でjson.Marshal()で可能
	- 型変換を強引に実装
-	データ移行
- Nullな項目を扱う時
	- 扱い方
	- sql.Null系
- database/sql
- Optional型
	- https://github.com/tenntenn/optional


## メルカリカウルマスタデータ更新
### 登壇者
tenntennさん

### 概要
#### マスタデータバッチ処理
- 製品のマスタデータを持っている
	- データ提供元からもらってバッチで取得
		- JSON,TSV,etc
	- Cloud SQL(MySQL)にて管理
		- スケールする必要がないため
		- スケールするならDatastore利用
- マスタデータTSVはCloud Storage
- Cloud Storage -> Task Queue
- 大きなファイルを分割して処理
- GAEのcronジョブで更新
- 問題点
	- URL Fetch
		- 32MBまでしか読めない
		- GAE -> GCS
		- 固定長で分割するしかない
	- Task Queue
		- 更新順序が適当
		- 外部キー制約はとてつもなく遅くなる
		- 冪等性[<65;110;30Mを担保
			- リトライOKにする
- 分割する方法
	- 固定長で分割[<64;110;30M
		- 1レコードの最大サイズ+バッファ
	- 担当バイト+改行まで読み込むようにする
	- お互いにはみ出し合って処理する
- Data flow

#### マスタデータ変換[<64;110;30M
- データ提供元のデータを改変はNG
	- 動的に変換
- x/text/encoding/japanese
	- SJIS/EUP-JPに対応
	- x/text/transform
		- transform.Transform inteface
		- Reader	
			- 読み込みながら変換できる
		- Write
			- 同じく
		- filter的に利用可能
		- 結構実装は難しい
- バイト列の変換は難しい

## LT

### RubyエンジニアがGo
#### 概要
- gracefulにしたい
- kami (Web App Framework)
	- Gunosyは基本kamiを利用している
	- einhornを利用してgracefulを実現
- einhorn
	- 新プロセス → Ack
- FileListenerからnetwork listnerをコピー
- einhornにAck

### 高精度名寄せシステム
#### 概要
- オープンデータの加工
- データクレンジングに利用
- JSONキャッシュ
	- シリアライズ/デシリアライズが遅い
- strings.Index
	- rune非対応
	- byte数で計算している
	- utf8パッケージ
- 正規表現
	- unicode
- 内製テスト
	- testify/assert 
