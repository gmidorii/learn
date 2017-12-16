# AWS Lambda + Typescript

## 記事
https://qiita.com/goto63/items/1374ae0c1f1696266177
https://qiita.com/ryusaka/items/8778e3b2afcffc8e7ca6

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