# 機械学習初心者

## 概要
Machine LearningをTensorFlowのTutorialsを通して触れてみました。  
解いた問題は、手書き数値の認識です。

## Machine Learning
**おこなっていることは、任意のグラフに対して近似する関数(=Model)を見つけること** 

### 用語
- データ
  - (Training) データとそのラベルのセットを持つ必要がある
    - ex) 手書き文字 <-> 書いてある文字
  - (Test) accuracyを計測するためのデータ
    - 同様にデータと正解を持つ必要がある
- weight
  - 入力値への重みテンソル(行列)
- bias
  - weightと入力値の積をずらす
- x
  - 入力値
- y
  - 出力
- 活性化関数
```
evidence = W * x + b

-------------------
W = weight
x = input
b = bias
-------------------
```
- softmax function
  - evidence(活性化関数=activation functionの結果)を確率へと変換する
  - 確率より 0 <= y <= 1
- loss function
  - 期待値と実際の結果の差分を計測する関数
  - ※ loss functionの結果を0に近づけることが目標(=期待値と実際の値が一致する)

### Machine Learningの流れ
1. Trainingデータセットを用意する
2. activation function(`W * x + b`)を定義
3. loss functionを定義
4. TrainingするOptimizerを決めて、loss functionをセットする  
   ex) Gradient Decent (最小勾配法)
5. Trainingを実施  
   テストデータを利用して、関数を何度も実行する。  
   実行するたびに、`W`と`b`をOptimizeする。(TensorFlow利用時は勝手にOptimizeしてくれる)
6. Modelのaccuracyを測定する

### 肝となる部分
- 大量のTrainingデータを用意すること
- accuracyを向上させるためのチューニング

### サンプルコード
https://github.com/midorigreen/pyTensor

## 参考
- [TensorFlow - Getting Started](https://www.tensorflow.org/get_started/)
- [Softmaxって何をしてるの？](http://hiro2o2.hatenablog.jp/entry/2016/07/21/013805)

