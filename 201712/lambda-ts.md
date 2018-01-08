# AWS Lambda + Typescript

## 記事
- https://qiita.com/goto63/items/1374ae0c1f1696266177
- https://qiita.com/ryusaka/items/8778e3b2afcffc8e7ca6
- https://qiita.com/horike37/items/b295a91908fcfd4033a2
- [serverless](https://serverless.com/framework/docs/providers/aws/guide/intro/)
- [TypeScript Doc](https://github.com/Microsoft/TypeScript/tree/master/doc)
- [DynamoDB Sample](https://github.com/serverless/examples/blob/master/aws-node-typescript-rest-api-with-dynamodb/todos/create.ts)
- [AWS LambdaをTypeScriptで開発する](http://dream-of-electric-cat.hatenablog.com/entry/2016/11/06/231255)
- [LocalStack](https://dev.classmethod.jp/cloud/aws/localstack-lambda/)

## 対応メモ
- Typescript compile
  - Javascriptへコンパイルする必要がある
```sh
% npm install typescript -g
```
- webpack.config.tsを作成
  - `touch webpack.config.ts`
  - resolve.extensionsに空文字は指定できない
```js
  resolve: {
    extensions: ['.ts', '.js']
  },
```

### 実行
- local: ローカル環境
- -f: 関数名
- -p: input JSON
```
% sls invoke local -f hello -p event.json
```

### Deploy
- aws configureの設定が必要
- serverless.ymlを修正しないと `us-east-1` へdeployされる
```yml
provider:
  name: aws
  runtime: nodejs6.10
  region: ap-northeast-1
```