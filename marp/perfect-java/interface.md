<!-- page_number: true -->
<!-- $size: 4:3 -->

# Perfect Java 7章 Interface

----
# 見方
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