# [Java] ResolverStyle.LENIENTを利用して時間をいい感じに扱う

## 概要
Java8より利用可能な、 `java.time.format.ResolverStyle` を使って、  
扱いづらい時間をいい感じに利用できます。  
`25:16` ってどうやって扱えばよいのかわからないといった方に便利です。  

## java.time.format.ResolverStyleとは？
Java8から、 `java.time.format.DateTimeFormatter` を利用して時間をパースします。  
このパースする際の、スタイルがEnumとして `ResolverStyle` に定義してあります。  
`DateTimeFormatter` へ渡すことで、スタイルを切り替えることができます。  

### モード (3種類)
(STRICTとSMARTに関しては、実証できていないのでさらっと記載してます)

#### STRICT
> 日付と時間を厳密に解決するスタイル

#### SMART
> 日付と時間をスマートな(賢い)方法で解決するスタイル。

#### LENIENT
> 日付と時間を厳密でない方法で解決するスタイル。
- 厳密でない解決 = 良い感じに解釈してくれる
- 例
  - `2018-01-32` → `2018-02-01`
  - `2018-01-01 25:00` → `2018-01-02 01:00`
  - `2018-01-01 00:67` → `2018-01-01 01:17`

## 実装
### 実装クラス
`DateTimeFormatter#withResolverStyle()` にてスタイルを指定する  
[コード](https://github.com/midorigreen/resolver-style/blob/master/src/main/java/SampleResolver.java#L12)  

(一部抜粋)
```java
public class SampleResolver {
    private DateTimeFormatter formatter;

    public SampleResolver(String format, ResolverStyle style) {
        this.formatter = DateTimeFormatter
                .ofPattern(format)
                // ここが重要
                .withResolverStyle(style);
    }

    public LocalDateTime formatLocal(String s) {
        return LocalDateTime.parse(s, formatter);
    }
}
```

### テスト
下記のテストが全て通ります。  
[コード](https://github.com/midorigreen/resolver-style/blob/master/src/test/java/SampleResolverTest.java#L25)  
(一部抜粋)
```java
@DataPoints
public static Fixture[] fixtures = {
        // 一般
        new Fixture("2018-01-01 10:00", "2018-01-01 10:00"),
        // 月が13
        new Fixture("2018-13-01 00:00", "2019-01-01 00:00"),
        // 日が32
        new Fixture("2018-01-32 00:00", "2018-02-01 00:00"),
        // うるう年
        new Fixture("2020-02-29 00:00", "2020-02-29 00:00"),
        // うるう年 + 1日
        new Fixture("2020-02-30 00:00", "2020-03-01 00:00"),
        // 時刻が24
        new Fixture("2018-01-01 24:00", "2018-01-02 00:00"),
        // 時刻が48
        new Fixture("2018-01-01 48:00", "2018-01-03 00:00"),
        // 分が70
        new Fixture("2018-01-01 00:70", "2018-01-01 01:10"),
};

@Theory
public void formatLocal_LENIENT(Fixture f) throws Exception {
    SampleResolver r = new SampleResolver(FORMAT, ResolverStyle.LENIENT);

    LocalDateTime actual = r.formatLocal(f.input);
    assertThat(f.toString(), actual.format(actualFormatter), is(f.expected));
}
```

## 所感
利用した感じは、大方期待通りの動作をしてくれています。  
「日付またぎ等の仕様で25時xx分を扱わないといけない」となった場合に、  
利用してみるのはいかがでしょうか。

## 参考
- [ResolverStyle (Java Platform SE 8 )](https://docs.oracle.com/javase/jp/8/docs/api/java/time/format/ResolverStyle.html)
- [DateTimeFormatter (Java Platform SE 8 )](https://docs.oracle.com/javase/jp/8/docs/api/java/time/format/DateTimeFormatter.html)
- [JUnit | JUnit4のDataPointsによるテストとFixtureと流れるようなインターフェース](http://tbpgr.hatenablog.com/entry/20121003/1349285882)
