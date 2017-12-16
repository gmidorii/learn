# JDBC

## 概要
- Java Database Connectivity
  - Java標準
  - java.sqlインターフェイスを利用
   - 各DBごとに実装をできる
- X/Open SQL CLIに基づいている
- JDBCドライバ
- version
  - JDBC 1.2API
  - JDBC 2.0API

### JDBCドライバ
- データベースに接続する部分
- JDBC Driver APIインターフェイスの実装
- 4つのタイプ
- ドライバ
  - `jdbc:oracle:thin`
- データベースURL
  - jdbc:<サブプロトコル>:<サブネーム>
  - "jdbc:oracle:thin:@localhost:1521:ORCL"

#### 実行
1. クラスロード
```java
Class.forName("oracle.jdbc.driver.OracleDriver")
```
2. Connection
```java
// Oracle8iに接続
Connection conn =DriverManager.getConnection("jdbc:oracle:thin:@localhost:1521:ORCL", "scott", "tiger");
```

### マッピング
- Javaの型とDBの型をマッピングさせている
- マッピングはドライバ側で実装

### MySQL
- MySQL Protocolを利用して通信している
- MySQL Packet
- TCP接続

## 資料
- https://docs.oracle.com/cd/E16338_01/java.112/b56281/overvw.htm
- http://www.atmarkit.co.jp/ait/articles/0106/26/news001.html
- https://docs.oracle.com/javase/jp/1.5.0/tooldocs/windows/classpath.html