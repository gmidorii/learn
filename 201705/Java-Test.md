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

