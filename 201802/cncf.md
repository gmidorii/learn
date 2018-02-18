# CNCF v1.0

## 概要
- 特定の実装を推奨したりはしない

## Serverless とは?
- サーバーのマネジメントなしにアプリケーションをbuild/runすること
- コードを実行するサーバーを必要としないという意味ではない
- 以下を考えなくて良い(提供者が対応してくれる)
  - provisioning
  - maintenance
  - updates
  - scaling
  - capacity planning

### Platform
- Function as a Service (Faas)
  - event driven
  - アプリケーションコードを関数としてマネージする
  - event or HTTP req等をtriggerに実行する
- Backend as a Service (Baas)
  - third-party API based service
  - APIを通して実行されるためサーバーレスのように見える

### Benefit
- Zero Server Ops
  - No provisioning, updating, managing server infra
    - サーバー管理は重大な経費(人件費)
  - 柔軟なスケーラビリティ
    - pre-planned capacityをしなくて良くなる
    - 逆にauto scalingのルールを設定しておく
    - コードが実行されていないときはお金はかからない

### Use Cases
- 非同期、並行、独立したタスクでの並列化が容易
- 頻繁ではなく散発的なリクエストがありスケーリング要件で予測できない変動がある
- ステートレスであり、コールドスタートを必要としない
- ビジネス要件がdynamcであり、開発速度を加速させる必要がある時

### Trade Off
- 動作が止まっていた状態からスタートした場合、パフォーマンスが低下する場合がある
  - 起動時間が加算されるため
- 機能面と非機能面の両方共を考慮する必要がある
- Non HTTP もしくは Non Elasticなワークロードではサーバーレスにアドバンテージがある
  - 例:
    - DBの変更に伴うロジック実行
    - ストリームプロセスのハンドリング

### 解決したこと
- オンデマンドができなかった問題を効率的に解決
- 伝統的なCloudでの問題を効率的に解決
- データサイズやリクエストサイズでない"largeness" という次元を示した
- 低エラーレートで自動でスケーリング
- ずっと早く問題を解決した

### Example

#### Multimedia processing
- File Uploadに伴う変換プロセスの実行
  - image -> thumbnail作成
- 頻繁に実行されないが要求に応じてスケールする必要があるシステム
  - atomic
  - 並列処理可能

#### Database changes or change data capture(CDC)
- Like Traditional SQL Trigger
  - main処理と並列に走らせることが可能に
- atomicでconsistency(一貫性)が重要
- 例
  - データの翻訳処理

#### IoT sensor input messages
- MQTT ProtocolでのIotデバイスからのメッセージ
- ネットワークに機器がつながることで大量スケールが必要

#### Stream processing at scale
- ストリーム処理でスケールする必要があるシステム

#### Chat bots
- 最初の応答でコールドスタート時間が許容されるケース

#### Batch jobs or scheduled tasks
- 非同期的に1日に数分間だけ計算を必要とするシステム
  - I/O, ネットワークアクセス

#### HTTP REST APIs and web app
- 従来のreq/resは十分適している
  - 静的サイト
  - オンデマンドでレスポンスを生成する

#### Mobile backends
- サーバーサイドロジックが小さい場合

#### Business Logic
- 特定のBusiness Logicを実行する機能を関数化
- Statefulなマネージャーにより実行される

#### Continuous Integration Pipeline