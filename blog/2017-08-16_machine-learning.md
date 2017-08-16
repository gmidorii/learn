# 機械学習初心者

## 概要
お盆休みを利用して、Machine LearningをTensorFlowのTutorialsを通して触れてみました。  
解いた問題は、手書き数値の認識です。

## Machine Learning
**おこなっていることは、任意のグラフに対して近似する関数(=Model)を見つけること** 

### 用語
- データ
  - (Training用) データとそのラベルのセットを持つ必要がある
    - ex) 手書き文字 <-> 書いてある文字
  - (Test用) accuracyを計測するためのデータ
    - 同様にデータと正解ラベルを持つ必要がある
- weight
  - 入力値への重みテンソル(行列)
- bias
  - weightと入力値の積をずらす
- x
  - 入力値
- y
  - 出力
- activation function
```
evidence = W * x + b

-------------------
Tensor
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

## 手書き数値の特徴量抽出Model
### Convolutional Neural Network
- 畳み込みニューラルネットワーク
- 面を一定の大きさのフィルタで覆い、領域で特徴量を抽出する方法
  - ex) 32x32の画像を5x5のフィルタでスライド1の場合 => 28x28の画像となる
- Filter
  - パラメータ
    - filterの数
    - filterの大きさ
    - filterの移動幅
    - padding
      - 画像の端の領域を0で埋める
- Layer
  - Convolutional Layer
  - Pooling Layer
    - サイズを圧縮する層
    - max poolingが利用される
      - 領域内の最大値を取る手法
  - Fully Connected Layer

## 参考
- [TensorFlow - Getting Started](https://www.tensorflow.org/get_started/)
- [Softmaxって何をしてるの？](http://hiro2o2.hatenablog.jp/entry/2016/07/21/013805)
- [Convolutional Neural Networkとは何なのか](http://qiita.com/icoxfog417/items/5fd55fad152231d706c2)
