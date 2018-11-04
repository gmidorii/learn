# Go

## テスト
* 基本的にはTable Driven Testで記載
* assertは利用せずに、ifで判定する
* エラーは下記のように記載
  * 例: `%n Hoge(%v) = %v; want %v`

## Effective Go
* gofmt
  * standarad packageはすべてフォーマットされている
  * Indentation > tab
  * Parentheses(かっこ) > 少なくしている
* Commentary
  * block /* */
  * line //

## beta
```
go get golang.org/x/build/version/go1.11beta1
go1.11beta1 download
```
or
```
brew upgrade goenv
goenv install go1.11beta1
```

## go.mod
https://roberto.selbach.ca/intro-to-go-modules
```bash
# GOPATH以外のディレクトリ
midori@midori ~/s/m/testmod> go mod init github.com/midorigreen/testmod
go: creating new go.mod: module github.com/midorigreen/testmod

midori@midori ~/s/m/testmod> cat go.mod
module github.com/midorigreen/testmod

# 利用側
midori@midori ~/s/m/usemod> go mod init mod
go: creating new go.mod: module mod

midori@midori ~/s/m/usemod> go build
go: finding github.com/midorigreen/testmod v0.0.1
go: downloading github.com/midorigreen/testmod v0.0.1

# module update (いずれか)
go get -u
go get -u=patch
go get github.com/rselbach/testmod@1.0.1

# この配下に配置される
midori@midori ~/s/m/usemod> ls $GOPATH/pkg/mod/cache/download/github.com/midorigreen/testmod/@v
list           v0.0.1.info    v0.0.1.mod     v0.0.1.zip     v0.0.1.ziphash v0.0.2.info    v0.0.2.mod     v0.0.2.zip     v0.0.2.ziphash

# v0 -> v1
midori@midori ~/s/m/testmod> echo "module github.com/midorigreen/testmod/v1" > go.mod
# go get -u
midori@midori ~/s/m/usemod> go get -u
go: finding github.com/midorigreen/testmod v1.0.1
go: github.com/midorigreen/testmod@v1.0.1: go.mod has post-v1 module path "github.com/midorigreen/testmod/v1" at revision v1.0.1
go get: error loading module requirements
# build
midori@midori ~/s/m/usemod> go build -o usemod
go: downloading github.com/midorigreen/testmod v1.0.1
main.go:7:2: unknown import path "github.com/midorigreen/testmod/v1": cannot find module providing package github.com/midorigreen/testmod/v1


# v1 -> v2
midori@midori ~/s/m/testmod> echo "module github.com/midorigreen/testmod/v2" > go.mod
# build
midori@midori ~/s/m/usemod> go build -o usemod
go: finding github.com/midorigreen/testmod/v2 v2.0.0
go: downloading github.com/midorigreen/testmod/v2 v2.0.0
```