# vim-goの便利コマンド一覧

## 概要
Goを開発している際に、vim-goを利用しています。

https://github.com/fatih/vim-go

最低限のコマンドしか利用できていなかったため、  
便利なコマンドを再洗い出ししてみます。  
(個人的なまとめの意味合いが強いです。)

## Commands
下記を参照しております。  
https://github.com/fatih/vim-go/blob/master/doc/vim-go.txt

### :GoRun
`go run` コマンドに相当します。  
vim上から実行できるところが便利です。  

### :GoBuild
`go build` コマンドに相当します。build後のバイナリは排出しないです。  
buildが成功するかどうかを確認する際に利用すると便利です。  
成功すると、下記のように出力されます。
```
vim-go: SUCCESS
```

失敗すると、quickfix windowに一覧が出力されます。
```
1 main.go|16| syntax error: unexpected semicolon or newline, expecting comma or }
Quickfix  go build -i . errors
vim-go: FAILED
```

### :GoDef
カーソル以下の、宣言元にjumpできます。  
実際の実装がどうなっているか確認したりする際に利用します。  
`gd` で同等の動作をします。  
(基本的には、 `gd` で移動することが多い印象です。)


### :GoCallers
カーソル以下のfuncの、呼び出し元を一括検索できます。  
(ファイル全検索していたのが、馬鹿らしく思えます..)  
*注意点* として、複数packageで検索したい場合は、 `:GoGuruScope` でスコープを設定します。  
(どのディレクトリ以下で検索をかけたいかを設定するイメージです。)
```
:GoGuruScope github.com/midorigreen/gprof
```

Sample: selectPeco()のカーソル上で
```
:GoCallers
```
quickfix window
```
1 main.go|93 col 6| github.com/midorigreen/gmd.selectPeco is called from these 3 sites:
2 /Users/midori/src/golang/src/github.com/midorigreen/gmd/exec.go|28 col 28| static function call from github.com/midorigreen/gmd.cmdExec
3 /Users/midori/src/golang/src/github.com/midorigreen/gmd/hist.go|38 col 28| static function call from github.com/midorigreen/gmd.cmdHist
4 /Users/midori/src/golang/src/github.com/midorigreen/gmd/del.go|29 col 25| static function call from github.com/midorigreen/gmd.cmdDel
```

### :GoCallstack
カーソル以下のfuncのcallstackがquickfix windowで見れます。  
(こちらも `:GoGuruScope` の指定が必要です。)  
```
1 main.go|120 col 7| Found a call path from root to github.com/midorigreen/gmd.run
2 main.go|120 col 6| github.com/midorigreen/gmd.run
3 main.go|138 col 13| static function call from github.com/midorigreen/gmd.main
```

### :GoTest
テストを実行します。  

### :GoTestFunc
カーソル以下のtest funcのみテストを実行します。

### :GoDoc
カーソル以下のGoDocを別windowで参照することができます。
(Shift+k と同等の認識です。)

### :GoDocBrowser
カーソル以下のGoDocをブラウザ上で参照することができます。  
わざわざGoDocを検索する手間が省けます。  

### [range]:GoAddTags
range指定した、structにタグを自動で追加してくれます。
`Before`
```golang
type Prof struct {
    Cores     []Core
    Model     string
    ModelName string
    CacheSize int32 
}
```
`After` (defaultはjsonタグ)
```golang
type Prof struct {
    Cores     []Core `json:"cores"`
    Model     string `json:"model"`
    ModelName string `json:"model_name"`
    CacheSize int32  `json:"cache_size"`
}
```
- dbタグ = `[range]:GoAddTags db`
- omitempty = `[range]:GoAddTags json,omitempty`

### :GoFillStruct
structの宣言時に、literalをdefault値で埋めてくれます。  
初期化時に、各literalを打たなくて良いのが楽です。  
`Before`
```
prof := Prof{}
```
`After`
```golang
prof := Prof{
    Cores:     nil,
    Model:     "",
    ModelName: "",
    CacheSize: 0,
}
```

### :GoRename [to]
カーソル以下の文字を、一括でrenameしてくれます。  
呼び出し元も、合わせて修正してくれろところが便利です。  
`func`, `struct`, `変数名` のそれぞれ動作可能です。  

### :GoPlay
開いているファイルをGo Playgroundで見ることができます。  
サンプルコード書いて、展開する場合とかに便利そうです。  

## 所感
他にもまだまだ眠ってそうですが、追加され次第、追記していく形を取ろうかと思います。  
ここまで揃っていると、vimで十分な効率で開発できそうです。
