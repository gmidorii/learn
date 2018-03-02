<!-- $theme: default -->

<!-- page_number: true -->
<!-- $size: 4:3 -->

# Perfect Java 7章 Interface

----
# この章の見方/観点
- オブジェクトがどんな振る舞いをするか
- どんな機能を提供するか
- 異なるクラスが同一のインターフェースを提供する意味

----

# 境界と抽象化

- 規模が大きくなると "機能" と "役割" で分割する
- 全体を組み立てるときに詳細を意識しないようにする
	- 詳細まで見きれない
    - *部品として抽象化することが大事*

→ プログラミングの原理原則

----

# Interface
- Interfaceは型の一つ
  - Javaでは、変数/オブジェクトに型を使える
- 参照型変数の型に使う
  - 参照先のオブジェクトに対して呼べるメソッドを確定する
  - 型→メソッドの流れ

----

# クラスとインターフェース
- クラス
  - 雛形
  - 型定義
- オブジェクトの振る舞いの共通性を担う役割は同じ

## 違い
- Interfaceは雛形ではない

- 型定義に特化した言語機能

```java
interface HogeInterface {
  void print(); //抽象メソッド
}
public class Hoge implements HogeInterface {
  @Override void print() {
    System.out.println("print");
  }
}
```

----

# Interface型の変数
```
HogeInterface h = new Hoge();
```

- 参照型変数を通じて呼べるメソッドは型で決まる
- メソッドの本体は、参照先のオブジェクトで決まる

----

# メソッドの引数
```java
class FugaString implements CharSequence{
}

void method(FugaString s) {
  System.out.println("length" + s.length());
}
```
↓ あ、Stringクラス渡したくなったからFugaStringを拡張継承
```java
void method(String s) {
  System.out.println("length" + s.length());
}
```
→ StringBuilder渡したくなった..終わった..


----

# メソッドの引数

実際は？ これでOK
```java
void method(CharSequence s) {
  System.out.println("length" + s.length());
}
```

----

# メソッドの引数
- 今回の例は、該当クラスがちょうどCharSequenceを継承してるから？
  - 確かに..
- ただ、よく見るとlength()メソッドにしか依存してない!!

#### 要するに
引数の方をInterfaceにすることで、
*「特定の振る舞いをするメソッドをもつオブジェクトを受けるよ！」*
と表明できる

----

# メソッドの引数
3つがlength()メソッドを持ってればよくない?
→ Javaの思想的にNG

## 特別な関係
- 変数の型が、オブジェクトのクラス型と一致
- 変数の型が、オブジェクトのクラスの拡張元クラスと一致
- 変数の型が、オブジェクトのクラスが実装するインターフェースと一致

----

# 多態性 (Polymorphism)
```java
class FugaString implements CharSequence{
  public int indexOf(String str) {}
}

class Hoge {
  public void print(CharSequence c) {
    // エラー
    c.indexOf();
  }
}

// エラー
new Hoge().print(new FugaString());
```
→ 呼び出すメソッドは型に依存
→ ダウンキャストでやる方法は危険なのでNG

----

# 多態性 (Polymorphism)
メソッドの引数
### CharSequenceで渡す
- CharSequenceを実装した任意のクラスのインスタンスを参照できる
→ これこそ多態性

![](./img/polymorphism.png)

----

# 多態性 (Polymorphism)

### クラスの拡張継承
- 多態性
- 実装コードの共有

### インターフェース
- 多態性

----

# 依存性

実装(クラス)と振る舞い(インターフェース)では経験的に実装が変化しやすい
→ 逆に、変化しづらい振る舞いをインターフェースとして切り出す

### クラスとインターフェースの適切な使い分け
- 変化しづらい振る舞いをインターフェースにまとめる
- クラスはインターフェースを継承することで変化しづらい部分を表明する
- クラスはインターフェースに依存したコードにより、変化しづらい部分のみへの依存を保証する

----

# インターフェースと抽象クラス
## 同じ
- インスタンス化できない型定義であること

## 違う
- 抽象クラスは雛形としての役割も担う
- インターフェイス: 振る舞いの規定
- 抽象クラス: 実装の拡張の役割

----

# インターフェースと抽象クラス
![](./img/abstract.png)

- クラスの拡張継承は、実装の継承のため
- インターフェイスの継承は、振る舞いの継承のため

----

# インターフェースと抽象クラス
## 現実的な使い分けの指針
変数の引数について
### インターフェイス
- 境界を意識するコード
- APIとして公開するコード

### クラス
- それ以外