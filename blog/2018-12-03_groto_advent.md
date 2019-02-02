こちらは、[Go3 Advent Calendar 2018](https://qiita.com/advent-calendar/2018/go3) の 3 日目の記事です。
昨日は、[@takochuu](https://qiita.com/takochuu) さんの [標準パッケージから見るパッケージ構成のパターンの解説](https://qiita.com/takochuu/items/5fefa57418b7b2cc3bd1)でした。
(標準パッケージでも、様々な構成があって勉強になりました。)

## 概要

Go で非常に簡単なプロトコルを実装してみました。  
(アプリケーション層のプロトコルです。)

## 対象

- プロトコルってどうやって作るのか気になる方
- なんとなく読んでみたい方

## 内容

- 簡単なプロトコルの Go 実装
  - アプリケーション層のプロトコルを実装
    - ネットワーク層は TCP 通信

## 詳細

### プロトコル
簡単なユーザー認証が入ったプロトコルで、下記の通りに通信します。

1. 初期リクエストパケット (Clinet -> Server)  
対象リクエストか否かを判定するため、Clientから特定のbyte(ENQ)のリクエストを送ります。
```
|-----------------------------|
| 1byte | リクエスト通知 (0x05) |
|-----------------------------|
```

2. 初期レスポンスパケット (Server -> Client)  
対象の場合は、認証のための情報を返します。
今回はパスワードハッシュ化のためのキーを返しています。
```
|-----------------------------------------|
| 1byte  | リクエスト結果(OK=0x06, NG=0x06) |
| 3byte  | バージョン ex)1.0.0             |
| 10byte | リクエストID                    |
| 20byte | パスワードハッシュ化キー           |
|-----------------------------------------|
```

3. 認証リクエストパケット (Client -> Server)  
ユーザーIDとパスワードを送信します。パスワードは先程受け取ったハッシュ化キーを利用してハッシュ化します。
(※ ユーザーIDは10byte以上のデータは利用できないです)
```
|------------------------------|
| 3byte  | バージョン            |
| 10byte | リクエストID          |
| 10byte | ユーザーID            |
| 32byte | ハッシュ化済みパスワード |
|------------------------------|
```

4. 認証レスポンスパケット (Server -> Client)  
受信したユーザーID/パスワードから、結果を返します。
```
|------------------------------------|
| 3byte  | バージョン                  |
| 10byte | リクエストID                |
| 1byte  | 認証結果(OK=0x06, NG=0x06) |
|------------------------------------|
```

5. データ送/受信  
(※ ここからは、まだ未実装かつプロトコル未定)

### 実装

[midorigreen/groto](https://github.com/midorigreen/groto)

```sh
├── example         # サンプル
│   ├── client
│   │   └── client.go
│   └── server
│       └── server.go
├── README.md
├── go.mod
├── groto.go        # protocol定義
├── client.go       # protocol client実装
└── server.go       # protocol server実装
```


#### 使い方
[groto/example](https://github.com/midorigreen/groto/tree/master/example)  
プロトコルの利用方法としては、 `net.Conn` interfaceの実装を渡すだけで、認証まで実行するようにしています。  
Client
```go
// ユーザー/パスワードを渡してClient生成
cli := groto.NewClient(user, password)
// 初期→認証まで実行
if err := cli.Do(conn); err != nil {
  return err
}
```

Server  
```go
l, err := net.Listen("tcp", ":8080")
if err != nil {
  return err
}
for {
  conn, err := l.Accept()
  if err != nil {
    log.Println(err)
  }
  go func() {
    defer func() {
      conn.Close()
    }()
    // Server生成
    s := groto.NewServer()
    // 初期→認証まで実行
    if err := s.Do(conn); err != nil {
      return
    }

    // やりたい処理
    for {
      b := make([]byte, 2*1024)
      _, err := conn.Read(b)
    }
  }()
}
```
(書きながら気づいたんですが、このあたりの待受の処理等もプロトコル側の実装で持たせればよかったかなと思います。)


#### プロトコル側

基本的に各送信ごとのパケットを `struct` で定義しています。  
例えば、 `2. 初期レスポンスパケット` は下記の通りに定義しています。
```go

type PacketHandshake struct {
  status    Status
  version   []byte
  id        []byte
  pwHashKey []byte
}

```

struct to byteを`Marshal`、byte to structを`Unmarshal`として定義します。  
愚直な実装をしてます。
```go
func (i *PacketHandshake) Marshal() []byte {
  b := make([]byte, 0, initLen)

  b = append(b, byte(i.status))
  b = append(b, i.version...)
  b = append(b, i.id...)
  b = append(b, i.pwHashKey...)

  return b
}

func UnmarshalHandshake(b []byte) (PacketHandshake, error) {
  return PacketHandshake{
    status:    Status(b[0]),
    version:   b[1:4],
    id:        b[4:14],
    pwHashKey: b[14:initLen],
  }, nil
}
```

各パケットを上記のように定義して、後は決められた通りにパケットをやり取りする実装を書きました。
パケットの交換は `net.Conn` interfaceの`Read/Write`メソッドを利用して実装しています。

サーバー側
```go
type Server struct {
  hashKey []byte
}

func NewServer() *Server {
  return &Server{}
}

func (s *Server) Do(conn net.Conn) error {
  // 初期リクエストの処理
  if err := s.stepHandshake(conn); err != nil {
    return err
  }
  // 認証リクエストの処理
  if err := s.stepAuthN(conn); err != nil {
    return err
  }
  return nil
}
```

クライアント側
```go
type Client struct {
  user      string
  password  string
  id        []byte
  pwHashKey []byte
}

func NewClient(user, password string) *Client {
  return &Client{
    user:     user,
    password: password,
  }
}

func (c *Client) Do(conn net.Conn) error {
  // 初期レスポンス処理
  if err := c.stepHandshake(conn); err != nil {
    return err
  }
  // 認証レスポンス処理
  if err := c.stepAuthN(conn); err != nil {
    return err
  }
  return nil
}
```

各パケットの送受信を実装しているメソッドは下記で、`net.Conn` を引数に取ります。
```go
// Client側 メソッド
func (c *Client) stepHandshake(conn net.Conn) error {}

// Server側メソッド
func (s *Server) stepHandshake(conn net.Conn) error {}
```

パケットの流れベースで見ると、処理は下記の通りになっています。  
例) 初期パケット

① Client (初期リクエストパケット送信)
```go
// conn = net.Connを実装したstruct
// 今回は *TCPConn

// 初期パケット送信
_, err := conn.Write([]byte{0x05})
if err != nil {
  return err
}
```
↓  
② Server (初期リクエストパケット受信)
```go
b := make([]byte, 2*1024)
// 初期パケットを読み込み
_, err := conn.Read(b)
if err != nil {
  return fmt.Errorf("failed read connection: %v", err)
}

// 対象パケットか否かを判定
if b[0] != 0x05 {
  // エラー処理
}
// 結果を返すパケットのstructを作成
i, err := NewPacketHandshake(OK)
if err != nil {
  return fmt.Errorf("failed create init proto: %v", err)
}
// 結果をconnectionへ書き込み
_, err = conn.Write(i.Marshal())
if err != nil {
  return err
}
```
↓  
③ Client
```go
b := make([]byte, 33)
// 初期リクエスト結果を読み込み
_, err = conn.Read(b)
if err != nil {
  return err
}
// structへ変換
i, err := UnmarshalHandshake(b)
if err != nil {
  return err
}
// 結果を確認
if i.status != OK {
  return errors.New("failed init")
}
```

このような処理を、各ステップごとにClient/Server側で実装しています。

## まとめ
非常に簡易で実用性はほぼなさそうですが、プロトコル実装してみました。このプロトコルを拡張して、作ろうとしているツールの初期認証に利用したいと考えています。  
作る流れとしては、プロトコルの定義を作ってから実装の想定でしたが、とりあえず動くものが見たくなったため、ソケット通信を先に実装して動かしながら作りました。


## おわりに
ここまで読んでいただきありがとうございました。  
明日は、[@r-fujimoto](https://qiita.com/r-fujimoto)さんです。

## 参考

- [MySQLのプロトコル解説 - とみたまさひろ - Rabbit Slide Show](https://slide.rabbit-shocker.org/authors/tommy/mysql-protocol/)
- [Go でバイナリ処理 - Qiita](https://qiita.com/Jxck_/items/c64d9ae0e910762eab37)
- [日本語は 1 文字何バイト？ - Sanwa Systems Tech Blog](http://tech.sanwasystem.com/entry/2017/11/13/102531)
- [Go 言語で TCP の Echo サーバーを実直に実装する - Qiita](https://qiita.com/kawasin73/items/3371d35166af733c2ce4)
