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

