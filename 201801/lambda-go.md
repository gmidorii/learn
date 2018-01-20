# AWS Lambda for Go

## 概要
Lambda で Go 言語が公式サポートされました。  
ちょっと触ってみたので、記事にしました。

## サンプルコード
ひとまずサンプルコードです。
```go
package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Body struct {
	Word string
	Code string
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(request)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello " + request.Body,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
```

### 要点
* `main` 関数に `lambda.Start(Handler)` を書く
  * `func Start(handler interface{})`
* Handler の引数は 0~2 個まで指定可能
  * 2 個の引数の場合、第 1 引数は `context.Context` を実装する
* `aws/aws-lambda-go/events` 配下に各種eventのtypeが宣言されている
  * 例 ([リンク](https://github.com/aws/aws-lambda-go/tree/master/events))
    * API GW
    * S3
    * SNS
    * Kinesis
    * DynamoDB
    * CloudWatch
  * 各TypeごとにREADMEが用意されている

### デプロイと実行
1. architectureを `linux` に設定してbuild
```sh
% GOOS=linux go build -o main
```
2. buildしたバイナリをzip化
```sh
% zip deployment.zip main
```
3. 作成したzipをLambaに追加

4. ハンドラをバイナリ名と同じにする
ここで微妙につまりました。

### デプロイコマンド
zipアップロードが面倒なので、Makefileでデプロイ完了するようにしました。  
```sh
zip:
  GOOS=linux GOARCH=amd64 go build -o $(BUILDFILE)
  zip $(ZIPFILE) $(BUILDFILE)

upload:
   make zip
   aws s3 cp $(ZIPFILE) s3://$(S3BUCKET)/$(S3KEY)

deploy:
   make upload
   aws lambda update-function-code \
           --function-name $(FUNCTIONNAME) \
           --s3-bucket $(S3BUCKET) \
           --s3-key $(S3KEY)
```

## 参考

* [Announcing Go Support for AWS Lambda](https://aws.amazon.com/jp/blogs/compute/announcing-go-support-for-aws-lambda/)
* [aws-lambda-go](https://github.com/aws/aws-lambda-go)
* [deploy makefile](http://nwpct1.hatenablog.com/entry/lambda-makefile)