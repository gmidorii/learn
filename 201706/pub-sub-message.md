# Pub/Sub Messaging

## Pub/Sub Messagingモデル概要
- Publish-Subscribe Messaging Model

登場人物

| 名称        | 概要  | 説明          |
|:-----------|:-----|:-------------|
| Publisher | 配信者 | メッセージを配信する側 |
| Subscriver | 受信者 | メッセージを受信する側|
| Topic | 仲介者 | メッセージを受け付けて、配信する |

特徴
- 複数のPublisherがメッセージをトピックに登録できる
- 複数のSubscriberがメッセージをトピックから受信できる
- Subscriberはトピックに対して、配信を申し込む
  - 登録したトピックに対して全てのメッセージを受信できる
  - 下記条件の場合は受信しません
    - メッセージセレクターの設定に該当しないメッセージ
    - 有効期限が切れたメッセージ
- トピック
  - メッセージは送信した順序のままトピックに登録
  - Subscriberで受信する順序は下記をもとに決まる
    - メッセージの有効期限
    - 優先順位
    - メッセージセレクタの設定

## Bayeux プロトコル
- Webサーバー間でメッセージを非同期にやりとりするためのプロトコル
- メッセージは名前付きチャネルを経て、ルーティング及び送達される
  - サーバ <=> クライアント
- ajaxを利用したサーバープッシュ型のテクニック=「comet」
- 準拠要求
  - 全てのMUST及びREQUIREDを満たし、かつ、全てのSHOULDを満たすものは、"完全準拠(unconditionally compliant)"
  - １つでもSHOULDを満たさないモノは「条件付き準拠(conditionally compliant)」
- ワード定義
  - メッセージ
    - サーバーとクライアント間で交換されるJSONオブジェクトである
  - イベント
    - Bayeuxプロトコルを用いて送られるアプリケーション固有のデータである。
  - チャネル
    - 名前付きのイベント送達先（と送信元）。イベントはチャネルに対して投稿され、受信者はチャネルからこれを受け取る。

### HTTPプロトコル
- リクエスト/レスポンス型のプロトコル
- クライアント -> サーバー
  - リクエストメソッド(GET,POST)
  - URI
- サーバー -> クライアント
  - ステータスライン
  - プロトコルバージョン
- 基本的にサーバー -> クライアントへはクライアントの要求なしに通信を走らせない
- Bayeuxでは、双方向の非同期通信を走らせるためサーバー・クライアント間で複数コネクションをサポートしている
  - 通常アクセス用コネクション
  - ロングポーリング用コネクション
- (MUST NOT)Bayeuxプロトコルも、サーバーが全てのアプリケーションに対して処理を行うために3本以上のコネクションをはらない

### 様々なルール
- (MUST) クライアントはブラウザのシングルオンポリシーを尊守しなければならない
  - シングルオンポリシー = JSからはそれがダウンロードされたサーバー以外への接続は許可されない
- 2コネクション処理
  - (MUST) Bayeux実装はHTTPのパイプライン処理を制御し、リクエスト順を保持して実行しなければならない
- ポーリング
  - レスポンスを受け取った後、メッセージをサーバーに対して送信する
    - 次のメッセージを受け取れるようにする処理
  - Bayeux実装は、ロングポーリングと呼ばれる形式のポーリングをサポートする必要がある
- 接続ネゴシエーション
  - 接続の際、コネクション/認証/を交換して合意するネゴシエーションが行われる
  - その際、ハンドシェークメッセージがかわされる
- クライアントの状態
```
-------------++------------+-------------+----------- +------------
State/Event  || handshake  | Timeout     | Successful | Disconnect
            ||  request   |             |   connect  |  request
            ||   sent     |             |  response  |   sent    
-------------++------------+-------------+----------- +------------
UNCONNECTED  || CONNECTING | UNCONNECTED |            |
CONNECTING   ||            | UNCONNECTED | CONNECTED  | UNCONNECTED
CONNECTED    ||            | UNCONNECTED |            | UNCONNECTED
-------------++------------+-------------+------------+------------
```
- Bayeuxにおける名前や識別子に用いられる文字はBNF記法
```
// BNF記法
alpha    = lowalpha | upalpha
lowalpha = "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" |
           "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" |
           "s" | "t" | "u" | "v" | "w" | "x" | "y" | "z"
upalpha  = "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" |
           "J" | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" |
           "S" | "T" | "U" | "V" | "W" | "X" | "Y" | "Z"
digit    = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" |
           "8" | "9"
alphanum = alpha | digit
mark     = "-" | "_" | "!" | "~" | "(" | ")" | "$" | "@" 
string   = *( alphanum | mark | " " | "/" | "*" | "." )
token    = ( alphanum | mark ) *( alphanum | mark )
integer  = digit *( digit )
```
- チャネル
  - URIの絶対パス部分の形式で記述される名称
  - /meta/で始まるチャネル名は予約されている
    - (SHOULD NOT) リモートのクライアントは購読するべきでない
  - パターン指定
    - 単一のセグメントを表す"*"
    - 複数のセグメントにマッチする"**"
  - /service/チャネル
    - サービスチャネルに投稿されたメッセージは他のクライアントに伝送されない
    - (SHOULD) サービスチェネルに行われた購読要求を記録管理すべきでない
  - メッセージ
    - フィールド名と値名が並んだ順不同なJSONオブジェクトとしてて定義される
    - 必ず一つの"channel"フィールドを持たねばならない
```
channel_name     = "/"  channel_segments
channel_segments = channel_segment *( "/" channel_segment )
channel_segment  = token
```

### メッセージフィールド定義
|名 | 説明|
|:--|:--|
|channel| - メッセージの配送先および配送元を示す <br> - リクエスト: 配送先 <br> - レスポンス: 配送元|
|supportedConnectionTypes| - `/meta/handshake`へ送受信の際に利用 <br> - どの転送タイプを用いるか決める <br> - (MUST) long-polling,callback-polling等が存在|
|clientId| - クライアントをユニークに識別するもの <br> - `/meta/handshake`,投稿メッセージ以外の全てのメッセージに付与が必須|
|id | 全てのメッセージに付与可能 |
|timestamp | - ISO 8601形式(`YYYY-MM-DDThh:mm:ss.ss`) <br> - GMTで表記|
|successful | -リクエストの成否 <br> -`/meta/*`系のレスポンスに必須|
|error| エラー|

### メタメッセージ定義
#### handshake
クライアントは、`/meta/handshake`channelに対してメッセージを送ることで接続ネゴシエーションを開始する

handshake request (MUST field)
|field| 説明|
|:--|:--|
|channel| `/meta/handshake` という値|
|version| クライアントで処理されるバージョン|
|supportedConnectionTypes| 接続タイプの文字列|

handshake response
|field| 説明|
|:--|:--|
|channel| `/meta/handshake` という値|
|version| クライアントで処理されるバージョン|
|supportedConnectionTypes| クライアント・サーバー共にサポートする接続タイプ|
|clientId| クライアント固有のID|
|successful| true/false|

#### connect
クライアントは、`/meta/connection` channelにメッセージを送ることでコネクションを開始する。

connection request
|field| 説明|
|:--|:--|
|channel| `/meta/handshake` という値|
|clientId| クライアント固有のID|
|connectionType| 接続タイプ|

connection response
|field| 説明|
|:--|:--|
|channel| `/meta/handshake` という値|
|successful| true/false|
|clientId| クライアント固有のID|

#### subscribe
チャネルへの登録を行うために、`/meta/subscribe`にリクエストを行い、配信登録を行う。  

subscribe request
|field| 説明|
|:--|:--|
|channel| `/meta/subscribe` という値|
|clientId| クライアント固有のID|
|subscription| 購読するチェネル名/チャネルパターン/配列|

subscribe response
|field| 説明|
|:--|:--|
|channel| `/meta/subscribe` という値|
|successful| true/false|
|clientId| クライアント固有のID|
|subscription| 購読するチェネル名/チャネルパターン/配列|


## 参考
[Pub/Subメッセージングモデル](http://itdoc.hitachi.co.jp/manuals/link/cosmi_v0870/APKC/EU070377.HTM)  
[Bayeux プロトコル日本語訳](http://d.hatena.ne.jp/takaxi/20111128/1322485586)  
[怖くないバッカスナウア記法(BNF)入門](http://qiita.com/h_sakurai/items/3cc328a6db8941ac6336)  