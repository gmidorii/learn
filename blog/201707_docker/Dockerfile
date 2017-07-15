# ベースイメージの指定
# Docker Hub(https://hub.docker.com/)上に様々なイメージが公開されています。
# そちらから利用したいベースイメージを選択します。
FROM golang:1.8.3-alpine3.6

# 作業を行うディレクトリを選択
WORKDIR /go/src

# コマンドを実行します
# golangイメージは Alpine Linuxをもとに作られている(軽量のためDockerでよく利用されます)
# linuxコマンド(ls,mkdir)等が実行可能
RUN mkdir hello

# 作業を行うディレクトリを変更
WORKDIR /go/src/hello

# 作業ディレクトリからコンテナ内にファイルをコピーします
# COPY [ホスト側] [コンテナ内]
COPY . .

# 今回は、Go Appをbuildします
RUN go build -o hello main.go

# コンテナ起動時に実行されるコマンド
CMD ["/go/src/hello/hello"]