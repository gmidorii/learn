# dicon

## 概要
利用方法メモ+プルリク方法メモ

## プルリク方法
http://deeeet.com/writing/2014/07/23/golang-pull-request/

1. forkする
2. go get -u github.com/xxx/xxx
3. cd github.com/xxx/xxx
4. git ch -b "fix-generate-code"
5. git commit
6. git remote add origin-midori git@github.com:midorigreen/dicon.git

### ホーム
```sh
[midori@midori:~/src/golang/src/github.com/akito0107/dicon on fix-generate-code]
% make main
go build -ldflags "-X 'main.version=0.0.1' -X 'main.revision='" -o bin/dicon main.go
[midori@midori:~/src/golang/src/github.com/akito0107/dicon on fix-generate-code]
% ./bin/dicon generate --pkg sample
[midori@midori:~/src/golang/src/github.com/akito0107/dicon on fix-generate-code]
% cd sample
[midori@midori:~/src/golang/src/github.com/akito0107/dicon/sample on fix-generate-code]
% go build
# github.com/akito0107/dicon/sample
./dicon_gen.go:17:8: cannot use dicontainer literal (type *dicontainer) as type DIContainer in return argument:
        *dicontainer does not implement DIContainer (missing Sample2Component method)


```