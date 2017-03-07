# NewsChecker

## 概要
アプリ側(Swift)とサーバー側(Golang)を一緒に作成してみた備忘録  
あまり技術的な内容はないです。

## 成果物
毎日のニュースをなんとか(自分用に)カスタマイズしてみやすくしたい  

## 概要
- 自分専用(自分が良ければOK)
- 短期間で作る(3日)
- アプリ&サーバーどちらも書く
- Swiftを使ってみる(初心者)
- サーバー側はGoで
- Qiita記事に書く(あまりやったことない)

### Swift事始め
- [Swift初心者向けスライドまとめ5選](http://www.sejuku.net/blog/5029)
- [iPhoneアプリ開発入門 (全13回)](http://dotinstall.com/lessons/basic_iphoneapp_v2)

### Swift側(Client)
`UITableView` を作成して、リスト表示ができるようにしました。  
合わせて、`CustomView` を作って、好きなデータを配置できるようにしました。

### Golang側(Server)
`echo` ライブラリを利用して、APIサーバーを構築。  
GitHub APIを利用して、APIリクエスト→パース→返却までできるように。  
Dockerfileを作成して、dockerが載っている環境でどこでも動作するようにひとまずしました。  

### GCP
- 初登録
- `Container Registry` にDockerイメージをpush
- `kubernetes` 利用する
   - [GKE を使って golang アプリケーションコンテナを稼働させる](http://blog.kaneshin.co/entry/2016/12/15/133943)
   - Deployment, Serviceのみで起動
- ログをコンソールから確認可能

### 詰まったこと
- Swift
    - Xcodeの使い方分からない。。
    - ライブラリの導入(最終的にCocoapodsを使用)
        - Xcodeに認識されない([Make sure you are opening .xcworkspace and not .xcodeproj file.](http://stackoverflow.com/questions/40601961/xcode-and-cocoapods-no-such-module-error))
        - buildしたら(Alomofire)compilerエラー →　バージョンUPで解決(swift3にあう)
    - webviewの上部のタブバー
        - タブバーがwebviewに覆われて隠される  
          →　最前面に `self.myWebView.bringSubview(toFront: myToolbar)`
        - viewに戻るボタンを配置するも押下のイベントをキャッチしない
    - webViewの配置方法
        - WebViewControllerを利用すると全画面に配置される
        - UiViewControllerを配置し、その上にWebViewを配置する

- Golang
    - Dockerのpullに詰まった
        - [Docker for mac v1.13でdocker pullできない問題](http://matsnow.hatenablog.com/entry/docker-for-mac/docker-pull-problem)
        - リンクと同様にbeta版を導入することで解消

- GKE
    - ポート指定がうまくいかず返ってこない。。
    - spec.template.spec.containers.pors.containerPort = コンテナ内のポート
    - `docker run -it -p 80:[containerPort]`