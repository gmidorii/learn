# Java Testing

## メモ

### テストコード基本
- 命名
  - テストクラス名 = TargetClassTest
  - テストメソッド名 = method_日本語で処理について記載
  - テスト対象 = sut (System Under Test)
  - テストごとの前処理 -> `@Before` (setUp())
  - テストごとの後処理 -> `@After` (tearDown())
  - テストクラスごとの前後処理 -> `@BeforeClass`, `@AfterClass`
- テストのフェーズ
1. 事前準備 = set up
2. 実行 = exercise
3. 検証 = verify
4. 後処理 = tear down
- Assertは、`org.junit.Assert.assertThat()` を利用する
  - static importを行っておく
- Matcherは、2種類利用する
  - `org.hamcrest.CoreMatchers`
    - `is()`  equalsメソッドによる比較
    - `nullValue()`  nullであることを検証する
      - 利用する際は、 `is(nullValue())` と利用する
    - `not()` 評価値を反転させる
  - `org.junit.matchers.JUnitMatchers`
    - `hasItem()` Iterableインタフェースを実装したクラスに、期待値が含まれているか検証
    - `hasItems()` 複数指定可能
  - BaseMatcher<> を継承することで、カスタムMatcherが作成できる
- @RunWith(Enclosed.class)
  - テストケースごとにclassを分割してテストを行う
  - 共通の初期化処理を複数パターン行いたい場合に利用する

```java
public class TargetClassTest {
  @Before
  public void setUp() {

  }

  @After
  public void tearDown() {

  }

  @Test
  public void hasUserName_ユーザー名称をもつ場合_trueを返却() throws Exception {
   // set up
   /** テスト対象のオブジェクトの変数名は sut とする **/
   TargetClass sut = new TargetClass)();

   // exercise
   boolean actual = sut.hasUserName();

   // verify
   /** 英語の構文 "assert that actual is expected" を意識する**/
   asseertThat(actual, is(true));
  }

  @Test(expected = NullPointerException.class)
  public void fetchUserName_ユーザー名称を取得に失敗_exception発生() {
   // set up
   /** テスト対象のオブジェクトの変数名は sut とする **/
   TargetClass sut = new TargetClass)();

   // exercise
   /** 下記メソッド実行時にException発生 **/
   Result actual = sut.fetchUserName();
  }
}
```

### Fixture
- 事前準備にて、設定する情報のこと
- inline set up
  - 各テストメソッドごとに、fixtureのset upを行う
  - simpleに設定を記述すれば良い
  - コードが長くなり、可読性が悪くなりがち
- implicit set up
  - @Beforeアノテーションをつけたメソッドにて設定を行う手法

### パラメータ化テスト
- Theories
  - `@RunWith(Theories.class)` を利用
  - `@Theory` テストメソッドのアノテーション
  - `@DataPoint` パラメータ
```java
@RunWith(Theories.class)
public class TargetClassTest {
  @DataPoint
  public static int PARAM_1 = 1;

  @DataPoint
  public static int PARAM_2 = 2;

  public TargetClassTest() {
    // 初期化処理
  }

  @Theory
  public void testCase(int x) throws Exception {

  }
}
```

### Rule
- テストをプラグイン的に拡張できる機能
- publicなfiledにアノテーションをつけて利用する
```java
@Rule
public Timeout timeout = new Timeout(100);
```
- 処理はテストごとに実行される → テストごとに実行したい共通処理をRuleとしてまとめられる
- `@ClassRule` にてテストClassごとに１回のRuleも作成可能
- カスタムルールの作成
  - `org.junit.rules.TestRule` インターフェイスを継承する
  - `Statement apply(final Statement base, Description description)` をOverrideする
    - 引数で渡される `Statement base` が各テストメソッドのイメージ
    - `base.evaluate()` を実行することでテストが実行される
  - これを利用することでテストに共通の前後処理を定義することができる
```java
public class HogeRule implements TestRule {

  private void before() {}
  private void after() {}

  @Override
  public Statement apply(final Statement base, Description description) {
    // new Statement()することで実際のテストを拡張している
    return new Statement() {
      // 前処理
      before()
      
      // テスト実行 (@Before -> テスト実行 -> @After)
      base.evaluate();

      // 後処理
      after()
    }
  }
}
```
- カスタムルールの作成ver2 (← こちらのほうが効率よく作れる)
  - ExternalResourceを継承して作成
```java
// 上記のサンプルコードをすでに実装した抽象クラス
public abstract class ExternalResource implements TestRule {
    public Statement apply(Statement base, Description description) {
        return statement(base);
    }

    private Statement statement(final Statement base) {
        return new Statement() {
            @Override
            public void evaluate() throws Throwable {
                before();
                try {
                    base.evaluate();
                } finally {
                    after();
                }
            }
        };
    }

    /**
     * Override to set up your specific external resource.
     *
     * @throws if setup fails (which will disable {@code after}
     */
    protected void before() throws Throwable {
        // do nothing
    }

    /**
     * Override to tear down your specific external resource.
     */
    protected void after() {
        // do nothing
    }
}
```
  - 利用する際は、`before()`と`after()`をOverrideして、前後処理を記載する
  - Rule内で利用する変数はprivate fieldに宣言しておき、コンストラクタで受け取る

### モック
- `org.mockito.Mockito` ライブラリを利用する
  - モックの作成 `mock(TargetClass.class)`
  - 返り値の設定 `when(sut.exec(anyInt(), any(Date.class))).thenReturn(obj)`

### データベースのテスト
- H2 Databaseの利用
  - ピュアJavaのSQLデータベース
  - インメモリ or ファイルの２種類の作成方法がある
  - 組み込みモード、サーバーモード、ミックスモードで動作
  - jarファイル１つ(1.5MB)動作
  - JDBCサポート
- DBUnit
  - @Ruleとして、作成することでDB接続周りを一手に引き受けられる
  - `org.dbunit.AbstractDatabaseTester` を継承する

### コードカバレッジ
#### カバレッジの種類
- C0(命令網羅)
  - プログラム中に定義された命令が１回以上実行されたかを測定 
  - line coverageと同等
  - 基準
    - ソースレベルのステートメント
    - バイトコード上のイントラクション
- C1(分岐網羅)
  - すべての分岐(if)を１回以上実行したかを測定
  - `if`条件の `true` or `false` のどちらも通す必要がある
- C2(条件網羅)
  - すべての条件を１回以上実行したかを測定
  - `if`条件の`true`になる全ての条件と`false`になるすべての条件を見る必要がある
- カバレッジ測定ツール
  - [EclEMMA(http://www.eclemma.org/)
    - Eclipseのプラグイン
    - カバレッジ測定エンジン`Jacoco`を利用
  - [Jacoco](http://www.eclemma.org/jacoco/)
    - カバレッジ測定エンジン
    - Mavenにライブラリとして追加 [Maven Repository](http://www.eclemma.org/jacoco/trunk/doc/maven.html)
```sh
# テスト実行 + カバレッジ測定
% mvn clean jacoco:prepare-agent test jacoco:report

# カバレッジレポート参照
% target/site/jacoco/index.html
```

## 参考文献
[JUnit実践入門 ~体系的に学ぶユニットテストの技法 (WEB+DB PRESS plus)](https://www.amazon.co.jp/JUnit%E5%AE%9F%E8%B7%B5%E5%85%A5%E9%96%80-%E4%BD%93%E7%B3%BB%E7%9A%84%E3%81%AB%E5%AD%A6%E3%81%B6%E3%83%A6%E3%83%8B%E3%83%83%E3%83%88%E3%83%86%E3%82%B9%E3%83%88%E3%81%AE%E6%8A%80%E6%B3%95-WEB-PRESS-plus/dp/477415377X/ref=sr_1_1?ie=UTF8&qid=1495546152&sr=8-1&keywords=junit+%E5%AE%9F%E8%B7%B5%E5%85%A5%E9%96%80)
[](http://qiita.com/hitomatagi/items/42fbb031ca95af50bb7e)