# AWS

## 情報源
- [公式ドキュメント][https://aws.amazon.com/jp/documentation/]
- [JAWS-UG CLI ハンズオン一覧][http://qiita.com/tcsh/items/b55eee599ae2c8806e4f]

## IAM (Identity and Access Management)
- 使用料無料
- IAMユーザー : 認証
  - AWSアカウントに紐づくユーザー
  - 個別のアクセスキーを作成できる
  - AWSのアクセスキーを取得するために作成したりもする
  - デフォルトでは何にもアクセスできない
  - IAMグループにまとめてポリシーをアタッチする
- IAMポリシー : 承認
  - 許可されているアクションは何かの定義
  - ユーザーベースポリシー
    - ユーザーがアクセスできるリソースを定義
  - リソースベースポリシー
    - 許可されるアクションと、影響を受けるリソースを指定
    - リソースへアクセスできるユーザーも合わせて定義
    - Principalエレメントを利用
  - ポリシー(JSON)は最適化変換をされる可能性があるため文字列比較をしない
  - Actions, Resources, Effect
- IAMロール : 委任
  - アクションおよびリソースへのアクセスを許可する一連のアクセス権限
  - 利用可能者
    - ロールと同じAWSアカウントのIAMユーザー
    - ロールと異なるAWSアカウントのIAMユーザー →　これにより異なるアカウントからアクセスできるようにする
  - 委任するために
    - IAMロールに2つのポリシーをアタッチ
      - アクセス権限ポリシー(どのリソースにアクセス許可をさせるか) and 信頼ポリシー(どのアカウントに委任させるか)
    - 別アカウントのIAMユーザーにロールを引き受けられるアクセス制限ポリシーをアタッチ
  - ロールを受けると一時的にもとのアクセス権限は失う
- Principal
  - アクションを実行してリソースにアクセスできるエンティティ
  - root, IAMユーザー, IAMロールを設定できる
- アクセスキー
  - AWS CLIを使用するために利用する
  - コンソールで作成
  - `aws configure` を実行して、AccessKeyとSecretAccessKeyをセット
  - IAMユーザーのポリシーに則りアクセスできる
- AWS CloudTrailにてすべての動作が記録されている

## EC2(Elastic Compute Cloud)
- 仮想コンピューティング: インスタンス
- インスタンスストアボリューム
  - 一時データ用のストレージボリューム
  - 停止・削除で消える
- EBS(Elastic Block Store)
  - 永続的なストレージボリューム
- セキュリティグループを利用してインスタンスのアクセス制限
- タグ
- インスタンス購入オプション
  - オンデマンドインスタンス
    - 時間課金
  - リザーブドインスタンス
    - 低額の一括前払い
  - スポットインスタンス
    - 特定のインスタンスタイプを使用したい最大料金を設定
    - 最大料金を超える場合インスタンスはシャットダウン
    - オークション制
- 各タイプによってインスタンス数の制限がある
  - https://aws.amazon.com/jp/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2
- インスタンスタイプ
  - T2
    - バーストパフォーマンスインスタンス
    - 暇な時にCPU クレジットを蓄積
    - アクティブなときに CPU クレジットを使用
    - CPUクレジットを使い切るとベースラインまで下がるため性能が落ちる
  - M4
    - 最新世代の汎用インスタンス
    - デフォルトでEBS最適化
    - 専有EBS帯域有り
  - M3
    - 汎用インスタンス
    - 高速な I/O パフォーマンスのための SSD ベースのインスタンスストレージ
    - 専有EBS帯域なし
  - C4
    - コンピューティング最適化インスタンス
    - 高性能プロセッサー持ち
    - デフォルトでEBS最適化
    - 専有EBS帯域有り
  - C3
    - コンピューティング最適化インスタンス
    - SSDベースのインスタンスストレージ
    - 専有EBS帯域なし
  - メモリ最適化
    - X1, R4, R3
  - Accelerated Computing インスタンス
    - P2, G2, F1
  - ストレージ最適化
    - I3, D2
  - key pair
    - 基本的にキーペアでのみ(SSH)ログイン可能

## Amazon Linux
### 特徴
- AWS APIツールとCloudInitがインストール
- SSH キーペアの使用およびリモートルートログインの無効化
- Message of the Dayでアップデート通知

### CentOsとの違い
- RedHat系のディストリビューション
- バージョニングはYYYY.MM
- 常に最新版にアップデートされる
  - yumの設定ファイルを修正することで固定可能
- /opt/aws 以下にツールを一通りインストールされている
- aws cliもyumでupdateできる
- cloud-initが初期インストール
  - UserDataに設定したOSの設定情報を自動で起動時に反映してくれるツール
  - (CentOS7もcloud-init初期インストール) 
- SELinux無効
- デフォルトユーザー
  - AmazonLinux: ec2-user
  - CentOS7: centos

## EBS
- ブロックレベルのストレージボリューム
- データにすばやくアクセスする必要があり、長期永続性が必要な場合に推奨
- スナップショット
  - 増分バックアップ
  - 最初は時間がかかる
  - 暗号化可能
  - 最新のスナップショットさえあればボリュームを復元

## UserData
- EC2起動時に実行するコマンド群を定義できる
- 定義方法
  - シェルスクリプト
  - cloud-init
- シェルスクリプト実行
  - rootユーザーで実行される
  - `sudo` は不要
  - ユーザーフィードバックが必要なコマンドは実行できない
  - `/var/log/cloud-init-output.log` にログが出力される
  - `#!/bin/bash` で始める必要あり
- cloud-init
  - .ssh/authorized_keysec2-user ファイルを設定
  - [cloud-init Documentation][http://cloudinit.readthedocs.io/en/latest/index.html]
  - `#cloud-config` と先頭に記載

## ロードバランサ
- Elastic Load Balancing
  - Classic Load Balancer
  - Application Load Balancer
- 複数のインスタンスと複数のアベイラビリティーゾーン内で自動的にトラフィックをルーティングする
- 異常なインスタンスを検出する
- availability zoneの片方がとまったら、もう片方のazに割り振る
- 要求処理容量を自動的に縮小/拡大
- 内部ロードバランサーと外部ロードバランサーを利用する
- AWS Certificate Manager と統合することでSSL/TLS を有効化できる
- Route53と併用することでDNS fail overすることが可能
  - 代替ロードバランサーを設定する
  - sorry pageに振り分けたりできる
- Auto Scalingを利用してEC2の最小台数を設定可能
- ヘルスチェックを設定可能
- CLBはクロスゾーン負荷分散がデフォルトで無効
- ALBはインスタンスをターゲットとしてターゲットグループにルーティングを行う
- クロスゾーン負荷分散
  - 有効: 各インスタンスに均等にリクエストを振り分ける(azは関係なし)
  - 無効: azごとに均等にリクエストを振り分ける
- ラウンドロビンのルーティングアルゴリズム
- アイドルタイムアウト値設定可能
  - default: 60秒
- 時間 + 転送量で課金される
- Layer4
  - ネットワークプロトコルのレベルで動作
  - パケットの中身は見ない
  - HTTPやHTTPSの機能は無視
- Layer7
  - パケットを見る
  - HTTPやHTTPSのヘッダーを参照できる
- ALB
  - URLベースのルールを10個まで定義できる
  - コンテナベースのアプリケーションを認識しサポート
  - ポートレベルでヘルスチェックを実施可能
  - 削除保護可能
  - 異なるアプリケーション(ターゲットグループ)にバランシング可能
  - ひとつのドメインで複数サービスが利用可能となった
- 作業流れ
1. ターゲットグループを作成する
2. ターゲットグループにインスタンスを登録する
3. ALBを作成する
4. ルールを作成してターゲットグループへのルーティングを設定する
- 動的ポートマッピング
  - ひとつのEC2上の複数ポートに対してバランシング可能
  - ECSでは自動でターゲットグループにインスタンスIDとポートを登録可能
- Connection Draining
  - 一定時間サービスのリクエストは受け付けず、ただ残った処理は続行される
  - 表示はInService
- ELB自体もスケールする
  - ただIPアドレスが変わるためDNS名で必ずアクセスする
  - 間に合わなければHTTP 503を返す
  - pre-warning等で回避
- リクエスト振り分け
  - DNSラウンドロビンでAZごとに振り分け
  - ELBが均等になるようEC2に振り分け
- ELB上でSSL処理が可能 →　後ろの人達は気にしなくて良い
  - ELBにサーバー証明書をアップロードする
- stickey session
  - 同一ユーザーリクエストを同一EC2に分散
  - ELBで有効期限を設定可能
- [ELBのすごくわかりやすい資料][https://www.slideshare.net/AmazonWebServicesJapan/aws-black-belt-online-seminar-2016-elastic-load-balancing]

## CloudWatch
- リアルタイミングログモニタリングサービス
- CloudWatch アラーム
  - 通知送信
  - リソースに変更を加える
  - Auto scaling
- データポイントは時間の経過に伴うその変数の値
- メトリックス
  - 15 か月後に自動的に有効期限切れ(新しいデータが入ってこない場合)
  - メトリクスは監視対象の変数
  - 名前、名前空間、1 つ以上のディメンションで一意に定義
  - データは2週間保存
  - 1 メトリクス当たり最大 10 ディメンション割当可能
- 名前空間
  - 命名規則 AWS/<service>
- ディメンション
  - メトリクスの一意の識別子
- データ保存日数
  - 1分ごとのデータポイントを 15 日間保存
  - 5 分ごとのデータポイントを 63 日間保存
  - 1 時間ごとのデータポイントを 455 日間保存
- アラーム
  - OK – メトリックスの値が、しきい値を下回っている
  - ALARM – メトリックスの値が、しきい値を上回っている
  - INSUFFICIENT_DATA – アラームが開始直後であるか、メトリックスが利用できないか、データが不足していてアラームの状態を判定できない

## Amazon SNS(Simple Notification Service)
- Client
  - Publisher
    - topicにメッセージを送信する
    - Subscriberと `Asynchronous` に通信する
  - Subscriber
    - 例
      - サーバー
      - Mail
      - SQS
      - Lambda
- 手順
1. Topicの作成
2. Topicのポリシー設定(通信可能なPublisherとSubscriberを決定)
- senderとreceiverの仲立ちをする立ち位置
- receiverがsenderに依存しているモデル
- [Push通知だけじゃない][http://dev.classmethod.jp/cloud/aws/non-mobile-sns/]
- pub-sub messaging
  - publiserとsubscriberは1対多の関係
  - Asynchronous通信
  - publisherがsubscriberに依存していない(= senderがrecieverに依存していない)
- ポリシー
  - 各ポリシーは、1 つのトピックだけを対象とする必要あり
  - 固有のポリシ-IDが必要
  - ポリシーを構成するステートメントには固定IDが必要
- SNSポリシーアクション
  - sns:AddPermission
    - topicポリシーへのアクセス許可の追加を許可
  - sns:ListSubscriptionsByTopic
    - 特定のトピックへのサブスクリプションを全て取得することを許可
  - sns:Publish
  - sns:Subscribe
    - そのままの意味
- SNSキー
  - sns:Endpoint
  - sns:Protocol