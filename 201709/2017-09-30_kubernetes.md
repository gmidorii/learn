# 雰囲気でgRPC,GKE+kubernetes使ってマイクロサービス作る
(非公開)

## 概要
[雰囲気でgRPC,GKE+kubernetes使ってマイクロサービス作る](http://gong023.hatenablog.com/entry/2017/09/21/174819) の記事が分かりやすかったので、自分も同じように手を動かしてみました。  
備忘録的なので、あまり内容は重要ではないです。(↑の記事がすべてです。)

### アプリ作成
```sh
% cat ping.proto
syntax = "proto3";

service Ping {
        rpc Ok(OkRequest) returns (OkResponse) {}
}

message OkRequest {
        // type name = tag
        string quetion = 1;
}

message OkResponse {
        // type name = tag
        string answer = 1;
}

# proto file 作成
% protoc --go_out=plugins=grpc:protoc ping.proto
```
#### 実装
https://github.com/midorigreen/test-kube

- front [リンク](https://github.com/midorigreen/test-kube/blob/master/front/main.go)
- back [リンク](https://github.com/midorigreen/test-kube/blob/master/back/main.go)


### 構築
docker
```sh
# back
% docker build -t gcr.io/[project]/test-kube-back:v0.1 .
% gcloud docker -- push gcr.io/[project]/test-kube-back:v0.1

# front
% docker build -t gcr.io/[project]/test-kube-front:v0.2 .
% gcloud docker -- push gcr.io/[project]/test-kube-front:v0.2
```

claster 作成
```sh
% gcloud container clusters create test-kube --num-nodes=2 --machine-type=g1-small
```

#### front
Deployment
```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: test-kube-front
spec:
  replicas: 2
  template:
    metadata:
      labels:
        name: test-kube-front
    spec:
      containers:
        - image: gcr.io/[project]/test-kube-front:v0.2
          imagePullPolicy: Always
          name: test-kube-front
```

実行
```sh
% kubectl apply -f deployment.yaml
```

Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: test-kube-front
spec:
  type: LoadBalancer
  selector:
      name: test-kube-front
  ports:
      - port: 80
        targetPort: 8888
```

実行
```sh
% kubectl apply -f service.yaml
```

#### backend
Deployment
```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: test-kube-back
spec:
  replicas: 2
  template:
    metadata:
      labels:
        name: test-kube-back
    spec:
      containers:
        - image: gcr.io/[project]/test-kube-back:v0.1
          imagePullPolicy: Always
          name: test-kube-back
```

実行
```sh
% kubectl apply -f deployment.yaml
```

Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: test-kube-back
spec:
  selector:
      name: test-kube-back
  ports:
      - port: 8080
```

実行
```sh
% kubectl apply -f service.yaml
```

### 結果
```sh
% curl "http://[test-kube-front_EXTERNAL_IP]/ping"
{"answer":"pong"}%
```

### 全部削除
```
% kubectl delete svc test-kube-front
% kubectl delete svc test-kube-back
% gcloud container clusters delete test-kube
```

### 引っかかったところ
- gRpcは接続の確認はどうすればいいんだろう
  - curlで投げられる口が欲しい
- ingress,service,deployment 間の関係性を正確に理解していないと躓く
  - port指定がまぜこぜになってしまう
- manifestfileの更新は、 `kubectl replace -f xxx.yaml` で行う
  - `kubectl update` はdeprecatedで自動でreplace実行される
- serviceのreplace時には、spec.ClusterIPの指定を求められる
  - manifestfileに追記で解消
