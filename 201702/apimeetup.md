# API MeetUp 2017/02/03

## 日経電子版を支える基盤API

### 登壇者
Takayasu さん

### アーキテクチャ
* API Gatewayを通過する
	* 独自に実装?
	* djangoで作成
	* APIごとに認可認証をしている
	* アノテーションからドキュメント作成(swagger)
		* 仕様書から直接リクエストできる
* ElasticBeanStalkで各APIが実装
	* 役割ごとに別れている
		* search,flag,paper...
	* EBS = ELB + EC2
	* DockerコンテナがEC2に載っている
	* Docker.run.json→ コンテナ数等を設定
	* Log Sender(OSS)
* 各APIの共通部分をテンプレート化
	* テンプレートから作成したいAPIをすぐに作成できる
	* 各APIにテンプレートの差分を反映させられる
* 良かったこと
	* ビジネスロジックに注力でき開発スピードアップ
* Batchファイル
	* EB + Docker
	* RUNDECK
	* Docker pullして実行
* Log流れ
	* Log Sender -> Log Aggregator
* Deploy
	* ブランチごとに環境に対応
	* マージタイミングでデプロイ
	* 現在は10minでdeploy
	* Prod/Stagingをswapして作成
* Error監視
	* エラーの階層ごとに利用しているサービスが異なる
		* インフラ CloudWatch
		* 500エラー SENTRY
	* Jenkinsでe2eテスト
	* kibanaでダッシュボードを事前に作成

### その他
* AWS
	* 責任範囲
		* AWS設定
		* DockerContainer
	* Support契約を結んでいる
	* 要望を聞く文化がある
	* 1つのイベントから1つのlambdaしか発火できない
	* 無駄にlambdaを発火させないように(お金の問題)
* SaaSの活用
	* NewRelic
	* ImageAPIのSaas化
		* luaでValidation
	* ライブラリと同様の感覚での利用
* 今後
	* lambdaのエラーログが探しづらい
* 6名程度で運用されている


## レコチョクのサービス群を支えるAPIたち

### 登壇者
山本　耕琢さん

### アーキテクチャ
* 共通のAPIが存在する
	* そこから、データ、ロジック層へアクセス
* 層
	* クライアント
	* フロント
	* API
	* サービス
	* データ
* サービス層からIDを返却してもらう
	* IDをもとにデータ層にアクセスしてメタ情報を付与
* AWS環境に載っている
* Swaggerの導入
* Athena,X-Ray導入
	* 利用状況/ボトルネックを可視化
	* apacheログをAthenaで検索


## Swaggerを利用したフロントサービス開発

### 登壇者
松木さん

### Swaggerについて
- CodeGeneratorでPythonクライアントを作成できる
- definitionで定義したオブジェクトをそのままJSON化
- JSON Web Tokenを利用
- SwaggerCodeGeneratorで下記を生成できる
	- html
	- client
	- server
- 良かったこと
	- API仕様レビューがYAMLファイル上で行える
	- 書式が決まっている
	- API開発者とクライアント側で並行作業できる
	- コードとドキュメントで差分が出ない
	- swagger-specのリリースに合わせてAPI仕様に変更がある
	- 互換性管理がswagger-specsによって管理される
- 引っかかりどころ
	- operationIdの指定方法
		- pathの指定ルールと変更する必要あり
	- OpenAPI3.0への移行
		- 互換性がないらしい

## API公開について

### 登壇者
酒井さん

### 概要
- APIの外部公開を進めていく
- 公開用APIを作成する

## JavaEE8によるREST+Microservices

### 登壇者

### JavaEE8
- JAX-RS 2.1
	- Reactive Client API
		- ここ最近決まってきたばかり
	- Non Blocking I/O
		- AsyncResponseの利用
- MVC 1.0
 	- Bean Validation
- JCache 1.0
	- @Cache~でオブジェクトのキャッシュが可能となる
- Microservicesの実装が整備されている感じ

### Swagger-codegen
- JAX-RSのコードからswagger-specsを自動生成
- Java8でJAX-RS/Swaggerの問題点
 	- ISO3339の拡張形式に対応していない(Gson -> joda-time)
 	- javax.timeでデシリアライズする
