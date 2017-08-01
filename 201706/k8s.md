# kubernetes

## 概要
WEB DB Pressの記事を参考にk8sを利用してみました。

## コマンド
```sh
% gcloud config set project xxxx
% gcloud config set compute/zone asia-northeast1-a
% gcloud auth login

# クラスタ起動
% gcloud container clusters create one \
--cluster-version=1.6.7 \
--machine-type=g1-small
Creating cluster one...done.
Created [https://container.googleapis.com/v1/projects/xxxxx/zones/asia-northeast1-a/clusters/one].
kubeconfig entry generated for one.
NAME  ZONE               MASTER_VERSION  MASTER_IP       MACHINE_TYPE  NODE_VERSION  NUM_NODES  STATUS
one   asia-northeast1-a  1.6.7           35.190.233.232  g1-small      1.6.7         3          RUNNING

# push gcr
% gcloud container builds submit --config=manifest/cloudbuild.yaml .

# create pod
% kubectl create -f manifest/pod.yaml

# create replicaset
% kubectl create -f manifest/replicaset.yaml

# replace replicaset
% kubectl replace -f manifest/replicaset.yaml

# create deployment
% kubectl create -f manifest/deployment.yaml

# check deployment
% kubectl get deployment
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
goneup    2         2         2            2           50s

# check replicasets
% kubectl get replicasets
NAME                DESIRED   CURRENT   READY     AGE
goneup-1100937273   2         2         2         57s

# create service
% kubectl create -f manifest/service.yaml

# create ingress
% kubectl create -f manifest/ingress.yaml
```