# Serverless

## サーバーレス化について
- AWS Lambdaでの利用を想定

### サーバーレスアーキテクチャについて?

### 知っておくこと
- AWS Lambda
  - ストリーム or 同期実行
- API GW
- Dynamo DB?

### ツール
- ツール一覧

#### Swagger
- API GWの仕様をSwaggerで連携可能
- Swagger
  - API仕様書を作成できるツール
    - API仕様書作成の共通化を目的にxxx
- swagger 2.0を利用
  - open api 3.0はまだ未対応

#### AWS SAM
- AWS CloudFormationの拡張
- **説明追加**
- 構築
  - API GW
  - Lambda
  - DynamoDB(→要調査)
- Swaggerテンプレートと連携可能
- [aws-sam-local](https://github.com/awslabs/aws-sam-local)
  - β版?
  - ローカル環境でAPI GW + Lambda環境を構築
    - Dynamo/S3等は未対応
    - Swagger連携未対応
  - `template.yml` 1つで作成可能
  - デプロイも可能
    - CloudFormationが起動する
    - Access key/Passが必要
    - Jenkinsで利用
  - Goの拡張は10日以内に作成
  - コマンド
    - `sam local invoke xxx`
    - `sam local start-api`

#### Node.js環境構築
- ツール
  - Typescript
  - webpack
  - tslint
- パッケージ構成
  - xxx
- ソースコードは一般のAPIと変わらない

#### build/deploy
- webpackで各関数ごとにbundle
- S3でのデプロイ
  - 各関数のdir以下をzip化
  - S3UP
  - aws aws-lambda xxx
- SAMを利用
  - S3へUP
  - SAMテンプレートを実行