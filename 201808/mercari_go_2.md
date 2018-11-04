# mercari go 2
2018/08/10

## メルカリにおける開発環境/QA環境と、そこで使われるGoのツールについて @masudak
### SET
* software enginerr in Test チーム
  * 生産性向上/開発環境整備
  * SETに関する記事を書かれている
* 各チームの中にはいって実施している
* PMがJiraチケットを切って仕様をつめる
  * Product Mangerかな？
* Local Environment
  * docker-composeで構築される
* PRごとに環境が自動作成

## GoでGraphQLサーバを立てるぞ！ @vvakame
* GraphQLのサーバー側のお話
* Tree構造
  * Resolver
    * Queryの要求が刺さることによって起こる
    * **Resolverの集合がサーバーの実装**
  * Node
* 好むとこ
  * 単一エンドポイント
  * Introspection <- 
* Go
  * Resolverを並列に処理することができる
  * GAE/Goで動くので..
* GitHub
  * https://developer.github.com/v4/explorer/
* gqlgen
  * https://github.com/99designs/gqlgen
  * ErrorPresenterはあんまり使わないほうが良い
  * gqlgen/Lobby でチャット質問
  * Middleware
    * サンプルコード
* ベストプラクティス
  * Relayを勉強
    * UI側の設定のためServerの実装がある
    * 次サポートでTSありそう
    * 仕様を抑えておく
  * N+1
    * 処理の集約が難しい
  * RESTとの対比
    * REST => RPCなので正しいRequestの投げ方がある
    * GraphQL => 正しいリクエスト？
  * Schema設計
    * 1table = 1type

## Software Engineer, Infrastructure @cubicdaiya
* パフォーマンスが要求されているところで利用
* rpmで配布している
  * makeで作成
  * dockerを利用して対象OSに合わせたrpmを作成
  * yumサーバーをS3に配置
* slackboard
  * Webhook URL設定の集約ができる
  * https://github.com/cubicdaiya/slackboard
* GCP
  * 近いノードにバランシングを自動でしてくれる機能あり
* pvpool
  * 56000 req/secの要件
  * slotを利用してPoolingしている
    * 負荷軽減
  * 構成
    * storer
    * fluser
* go-httpstats
  * Mutexの競合が激しかったりする
* mfc
  * Fastly Client
  * サブコマンド
```go
switch args[1]{
  case "config":
  case "help":
  default:
}
```


