# React + TypeScript + Css-Loader

## 概要
めちゃめちゃハマったので書き残しておきます。

### 問題点
1. Module build failed
```
./src/styles/form.css
  0:0  error  Module build failed (from ./node_modules/css-loader/index.js):
```
-> style-loader → css-loaderの順にすると治った


2. Cannot find module
```
[at-loader] ./src/components/Form.tsx:3:24
  0:0  error  TS2307: Cannot find module '../styles/form.css'.
```
-> cssファイルの型定義を作成すると治った
-> tcsを利用する `npm -g install typed-css-modules`


