# Gopackage

## Golangのパッケージ構成について
golangで開発するにあたって、パッケージ構成がよくわからなかったため、調べてみました。  

当記事では、下記に焦点をあてて調べてみました。
- packageの分け方
- pacakge内でのファイルの分割
- ファイル内のコードの書き方

## パッケージ構成
[@benbjohnson](https://twitter.com/benbjohnson) さんの[記事](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)を参考とした例です。  

### Good Practice (by [@benbjohnson](https://twitter.com/benbjohnson))
下記の4点を考慮したアプローチ

1. Root package is for domain types
2. Group subpackages by dependency
3. Use a shared mock subpackage
4. Main package ties together dependencies

#### 1. Root package is for domain types
ドメインとは、データとプロセスがどのように相互作用を起こすかを記述する、高次元の言語のことです。  
ドメインは、技術的なバックグラウンドに依存することはありません。  
  
Root packageには、このドメインタイプを配置します。Root package内は、  
simpleな`struct`と`stuct`の振る舞いを定義した`interface`のみで構成されます。  
  
`The root package should not depend on any other package in your application!`  
→ Root packageはアプリケーション上の他のいかなるパッケージにも依存すべきでない

`ドメイン = User` の場合
```go
package app

type User struct {
  ID   int
  Name string
}

type UserService interface {
  User(id int) (*User, error)
  UpdateUser(id int, name string) error
}
```

#### 2. Group subpackages by dependency
Root packageにて、外部ドメインとの依存性を持てない  
→ Subpackageにて依存性をもたせます  
Subpackageは、ドメインと実装のアダプターとして存在しています。  
このように実装することで、テストの単純性と他のDBが追加された場合の実装の容易性につながります。  

  
Userデータの保持が`MySQL` の場合
```go
package mysql

import (
  "database/sql"

  "github.com/midorigreen/app"
  _ "github.com/go-sql-driver/mysql"
)

// implemented app.UserService
type UserService struct {
  DB *sql.DB
}

func (s *UserService) User(id int) (*app.User, error) {
  var u app.User
  row := db.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
  if row.Scan(&u.ID, &u.Name); err != nil {
    return nil, err
  }
  return &u, nil
}
```

#### 3. Use a shared mock subpackage
ドメイン`interface`利用により、依存性を分離することができました。  
ドメイン `interface` の実装をモックとして、作成することで処理ごとに  
実装をinjectすることが可能となります。

モック作成例
```go
package mock

import "github.com/midorigreen/app"

// UserService is mock implementation of app.UserService
type UserService struct {
  UserFn      func(id int) (*app.User, error) // 処理ごとに実装をInject
  UserInvoked bool
}

func (u *UserService) User(id int) (*app.User, error) {
  // mark function invoked
  UserInvoked = true
  return u.UserFn(id)
}
```

#### 4. Main package ties together dependencies
全ての依存性をもつパッケージを孤立化させられたので、後は結びつける必要があります。  
この役割をはたすのが、`main` pacakgeです。  
`main` packageでは、objectにどの依存性を入れるのかを選ぶ役割があります。

```go
package main

import (
  "log"
  "os"

  "github.com/midorigreen/app"
  "github.com/midorigreen/mysql"
)

func main() {
  db, err := mysql.Open(os.Getenv("DB"))
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // create service
  us := &mysql.UserService{DB: db}

  // etc
}
```

### (補足) イマイチなアプローチ
#### #1 モノリシックなパッケージ
全てのコードをひとつのパッケージに詰め込む方式  
→ 小さなアプリケーションに対しては十分にうまくいきます。

#### #2 機能(function)ごとにグルーピング
機能ごとに、処理をグルーピングすることで構成する方式  

構成例
- handler : ハンドラを集めたパッケージ
- model : モデルを集めたパッケージ
- controller : コントローラを集めたパッケージ

ただし、この手法には欠点が２点ほど存在しています。
1. 命名  
`handler.XXXHandler` のように名前が重複してしまう。
2. 循環参照  
パッケージ同士で参照しあい、循環参照(circular dependencies)を起こしてしまう。  
(1方向のみの依存関係をもつ場合のみしか、有効ではない手法)

#### #3 モジュールごとのグルーピング
モジュールごとにグルーピングをする方式  

こちらも同様に下記が欠点です
1. 命名  
`users.User` のように名前が重複してしまう
2. 循環参照
## 参考
[Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)  
[【翻訳】【Golang】標準的なパッケージのレイアウト](http://allishackedoff.hatenablog.com/entry/2016/08/23/015016)