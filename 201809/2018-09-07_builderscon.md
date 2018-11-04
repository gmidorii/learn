# builderscon
2018/09/07

## Session
1日目
### Envoyについて (by Envoy作者)
* LyftのモノリスからMicroServiceへの移行
* モノリスでも様々な問題が発生していた
  * ネットワーク
  * caching
* 3年前
  * モノリス + MicroService (Pythonが数種類)
  * 問題点
    * multiple language
    * many protocol
    * black boxのLB
    * 観測部分 (stats, tracing, logging)
      * 一貫性を持って観測できる必要がある
    * 分散処理
      * retry, cb, timeouts
      * Implementの一貫性が重要
    * 認証
    * 言語ごとにライブラリ内にベストプラクティスが入り込んでしまっている
      * 他の言語では？再実装?
    * デバッグのやり方がわからない
    * MicroServiceへの信頼性の低下
* これらの問題を解決するためにEnvoyを開発
* Envoy
  * 問題が起きたときに、ちゃんと観測できるようにしよう
  * できるだけインフラを隠したい
    * アプリケーションコード100%が理想
  * design goal
    * side car solution
      * 各機能を共通で実装
    * L3/L4 proxy
      * byte to byte
      * 多様なprotocolに対応する
    * HTTP2 first
    * configをAPI経由でできるようにした
      * ファイルをhostに投げて、ロードする必要があった←つらい
    * health checking
      * 各言語ごとの再実装をなくす
    * 観測可能性
      * ここを重要視している
    * Hot restart
      * k8sの世界
      * connectionを落とさずにrestart
  * architecture
    * filterをプラグイン的に追加できるように実装
    * xDs API
      * universal data plane
      * API郡
      * **proto3** で実装されている
      * Global bootstarp config
      * Eventual consistency
        * 結果整合性を持つこと
    * threading model(c10k)
      * c10k
        * x connectionごとのthread
          * メモリを食う
        * o 非同期プログラミングモデル
          * 複数のコネクションがthreadごとにある
          * connections -> event loop -> thread
        * Main Thread
          * etc..
        * Workder Thread
          * 1core 1worker
          * 100% non-blocking
          * lockなし
        * concurrentに動くがプログラマは意識しなくて良い
        * 何百ものthreadで動作するようにした
      * RCU (Read Copy Update)
        * Linux カーネルで使われているらしい
        * むずかしい..
        * workder threadへpostingする方法
      * Thread Local Storage
    * hot restart
      * full binary reload without dropping connection
      * stats/locks in shared memory
        * statsを旧/新 で共有するため
        * UDSでプロセス間通信
        * コンテナの世界ではプロセスはimmutable
          * プロセスforkしてのrestartはNG
          * そのためUDSで通信して整合性を担保する
      * stats
        * ...
    * histogram
      * histogram A/B がある
      * histogramをスワップさせる
      * スワップ完了後に、ロック無しで読み込む
* 質問
  * k8s Istio/envoy
    * ネットワークに関する問題がわかっている必要がある
  * envoy goal
    * コミュニティを伸ばしていくこと

### ランチ サイボウズ

### Algorithm in React
* 資料
  * https://speakerdeck.com/koba04/algorithms-in-react?slide=84
* Reactの内部実装のお話
* Reconciler
  * 役割
    * 更新管理
  * stack -> fiberへ変わった
  * stack reconcilerはソースから削除されている
* stack reconciler
  * traverseして、再帰呼出しをしてパッチを当てる
  * React Element
    * DOMと1対1のオブジェクト
    * no state
  * Public Instance
    * React.Component
  * Internal Instance
    * internal instanceをtraverseして処理
    * public instacleのfieldとしてset
  * mount componentを再帰的に呼び出している
  * 問題
    * 同期的な処理
      * ツリーを再帰的に呼び出すこと自体
* fiber reconciler
  * unit of work
  * render と commit phaseがある
  * ReactElementに対応して存在
  * fixed data structure
  * alternate Fiber
  * Linked List
    * ReactElement Tree => fiberのLinked List
    * parent/child の双方向の参照を持つ
    * treeをchildren -> parent -> (sibling) -> childre .. とたどりながら参照する
  * Double Buffer
    * current と work in progress
    * commitをすることでswapする
    * `Game Programming Pattarn` に記載されている
  * render phase
    * 副作用を抽出するだけ => 副作用を持たない
    * 非同期処理
  * commit pahse
    * commit side effect
  * begin workで下って complete workで上る
  * Side Effect
    * extract時は、何が起きているか程度を抽出するだけ
    * commit時に、diffを取る作業をする
  * Starvation
  * Expire Time
    * timeが近い順に実行される
    * timeが切れそうな場合は、同期的に実行される
  * Interactive Event Type
    * clicke等の動作がHigh Priorityとして処理される
    * High/Lowを明示的に設定することで、すぐ出したいものと、遅れて出すものを切り替えられる
    * 最新の挙動は注意
  * Bitwise Operator
    * 0b => bit列
    * Maskをかけて、LifeCycleEffectだけを抽出できる(ビット演算で)
    * createContextでbit演算結果を返して、renderの切り分けができる
  * Suspense
    * throw promise -> React.placeholderがキャッチ
    * 時間がたったときだけ、Loadingメッセージを出せる
    * 次のReact Versionで投入
    * ↑便利そう


### 静的検査のいろは
* 資料
  * https://speakerdeck.com/orgachem/introduction-about-static-analysis-without-previous-knowledge
* Vim Scriptの静的開発ツール作者(Vint)
* 目標
  * 特定のコメントを静的解析できる
* 抽象構文木
  * コードを木構造で作成したもの
  * 各葉?をノードと呼ぶ
  * 各言語ごとに標準実装があり簡単に見られる(Goでもいける)
* 構文解析
  * プログラム → 文 → 式/コメント ...
    * 各検査ごとに文法規則をあてはめる
    * なければ構文エラーになる
  * 文法解析
    * 始まる位置と終わった位置を関数の引数と戻り値で受け取れる
      * 解析できなかった場合は、indexのいちを進めない
    * 順々に処理ができる
    * バックトラックあり
      * 単純だが計算量が多い
  * Parser Combinator (参照B)
* 走査関数 traverse
  * 深さ優先探索する関数
  * 走査関数を実装することで様々なことができる


### 数学で需要予測
* 4つの分析
  * 記述的
    * ダッシュボード
  * 診断的
    * 因果関係
  * 予測的
    * 未来を知る
  * 処方的
    * 未来をコントロール
* 分析は数%の成長を目指す
* 分析の重要性
  * 長期運用になればなるほど
  * 様々なツール等が登場
    * BIツール -> 可視化
    * 機械学習 -> 機械学習
    * 因果関係は..?
* 問題点
  * 因果関係を履き違えると適切なアクションがとれない
    * データは解釈までは与えてくれない
* 正しくデータを解釈するためには？ => 数学
* 数学
  * 純粋数学の考え方を応用
  * 純粋数学
    * 実際に計算をするわけではない
    * 2種類
      * 問題解決
      * 理論構築
    * やっていること (かつ重要なこと)
      * なぜの追求
      * 数値と言語を行き来
    * 客観的な正しさを求める
    * 分析に戻すと..
      * 認知バイアスを取り除く
* ゲーム内の需要予測
  * 思考のフレームワーク
    * 前提条件とゴール
    * データを眺める
      * 統計的にどれくらいのサンプル数を見れば十分か
    * 仮説をたてる
      * ドメイン知識が重要
      * 書けるだけ書き出してみる
    * 仮説の数値化
      * ゴールと相関関係がありそうなものから順番に処理する
        * なるべく少ない係数で処理したい
      * 最初は最大値/最小値があっているような線形近似をする
        * 間のズレは、ズレにあった関数を作ってかければ良い
      * グループ化
        * 各相関があるグループごとに式をかける
    * 例外処理
      * 2,3割はかけ離れている
        * 理由を探していく
        * 少数
          * データが少ない場合は、外すもしくは影響を小さくする関数を作る
* 機械学習
  * データの前処理を先にやらないといけない
* 質問
  * 因果関係をもとめるには?
    * 誰に説明しても納得してもらえることをゴールとする

### Java Card
* 目標
  * Java Card Appletが書けるようになる
* Java Card
  * クレジットカード
  * SIMカード
  * etc..
* Java Card Quiz
  * float/charは扱えない
  * byteで基本的には扱う
* Smart Card (ISO 7816)
  * 2種類
    * Intellijent Smart Card (CPUあり←Java Card)
    * MemoryCard
  * Interface
    * 金のチップの部分
    * プロトコル: APDU
  * APDU
    * 外部とのprotocol
    * byte列でやりとり
    * Command(Request) APDU
      * レスポンスのContent Lengthをリクエスト時に指定する
    * Respones APDU
      * 90x系が成功
  * UICC
    * ISCの一種
    * Java Cardの実行基盤
    * Architecture
      * Core O.S
      * ファイルシステムあり
* Java Card
  * 2.0.1がスタート(理由は謎)
  * Java Runtime: JCRE
    * 通常のJVMのサブセット
    * 16bitが上限
  * Lifecycle
    * 変数はリファレンスがなくなったらlostかGC
      * GCはカードに実装されているかどうか次第
    * shutdownなし
    * 電源供給がないと「無限のクロックサイクル」
    * Java Card Lifecycle
      * Initialization phase
        * カード購入時に済んでいる
      * Personalization phase
  * Security
    * firewall (applet間でのデータやり取り)
  * Memory
    * インスタンス変数はEEPROM?に書かれる
      * インスタンス変数のデータは残り続ける
      * 電源が切れてもJVMは止まらない
    * RAM
      * 8byteしかない..
  * Applet
    * Java Card上の各アプリケーションのこと
    * 複数のappletを同時実行不可
  * Environment
    * ローカル変数はRAMに(native typeだけ)
    * JDK1.6が安定している..
  * Compile
    * compile => convert
* ただし、カードに書き込むことは基本的にはできない..
  * 各種鍵 + IDが必要なため
* 質問
  * Java Cardって結局どういうアプリが動いてる?
    * 認証系の処理
    * IDを特定する処理

### lld (by lld作者)
* lld
  * LLVMのサブプロジェクト
  * clang -> LLVM
  * とにかく速い
  * cross buildが簡単にできる
  * 事例
    * Chrome
    * Rust/ARM
  * 行数
    * 2.9万行程度
* リンカ
  * オブジェクトファイルを一つの実行ファイルにまとめる
  * 仕事
    * 名前を解決
    * リロケーション
* unixではelfと呼ばれる実行形式に統一
* 速さを考えるときは、重要な部分にフォーカスする
  * ちゃんとオーダーレベルで考える
* 開発手順
  * 空のアセンブラ→リンク→16進ダンプ→バイナリを読み解く
  * だんだん機能を足していく
* 最初は遅かった
  * ファイルフォーマットの違いが原因
    * コード統一
      * 抽象化のためのコードを大量に書かないといけない
      * 中間フォーマットコストが高い
  * 書き直しを決意
* リデザイン
  * 難しくしない
  * 自然と速くなるように作る ( **データ構造が最重要** )
  * データ構造デザイン
    * 遅い処理は最小数まで減らす
    * 数千万回行う処理はポインタのデリファレンスを1~2回にする
  * マルチターゲット
    * デザインは共有するがコードは共有しない
      * 4つの異なるリンカが1つの実行ファイルに同梱されている
        * mainで分岐している
      * アンチパターンっぽいが、大幅にシンプルになる
* 速くてシンプルなコードを書くためには?
  * データ構造が最重要
    * コードで最適化を目指す必要はない
  * 2回書く
  * 最適化箇所を最小に留める
    * ほかは読みやすさを重視する
* 動的メモリ管理
  * はじから順にメモリを使う
  * bump pointer allocator
    * さいしょからメモリを埋めていくだけ
    * 途中の穴埋めはしない
* 並列
  * 基本シングルスレッド
  * 必要なとこだけ、マルチスレッドにするけどスレッド間でやりとりしない


2日目

### OSの中身
* 資料
  * https://speakerdeck.com/ariaki/os-that-we-should-know
* 問題提議
  * 昔に比べて認知度が下がっている?
  * OS自体がロストテクノロジーになっていくのでは..?
* 原因
  * サーバーレスの台頭
  * 価値の変化
  * ハード/ソフト両面の高度化が進んだ=>理解難易度高
* `setenforce 0` => SELinuxの無効化
* ブートシーケンスからOSの動きを紐解く
* BIOS起動時にモードが何種類かある..
  * リアル
    * メモリどこでも読める..
  * プロテクト
    * メモリ保護ができる
    * 管理
      * 特権管理
        * call gate方式
        * Segmentation Fault
          * SYSSEGVシグナルに紐づく例外が発生
      * メモリ管理
        * PDE > PTE > page => ページング
        * リニアアドレス
          * PDE PTE OFFSET をまとめたアドレス
        * リニアアドレスをまとめた空間=>セグメント
          * 要するに何がどこにあるかを保存している空間?
        * セグメンテーション
      * タスク管理
        * マルチタスク
        * プリエンプティブ
          * Task State Segmentによってタスクの状態が管理
* Intel 8086
  * x86 => 16bit or 32bit CPU
  * AD => Data input/output 
  * A => 20bit output
  * レジスタ
    * 32bit=>E, 64bit=>R が先頭につく
* ブートシーケンス
  * Bootstrap Loader 初期化 512B
  * MBR(Master Boot Recorder) が0x7c00へ書き込む..?
  * INT = 割り込み
    * 例外割り込み
  * システムコール
    * Ring-3を超えた機能を呼び出す
    * User Landのライブラリを通してシステムコールを発行
      * User Landと Kernel
    * EAXによって機能が切り分けられている
    * 3種類の呼び出し方
      * SYSCALL
      * SYSENTER
      * INT 0x80
    * sys_call_tableを介して発行する
* 最近のCPU
  * マルチコアはL3 Cacheだけ共通化
  * L1/L2は各CPUごと
* Linux Kernel
  * Linux1.0から読むと良さそう
  * ほかは複雑

### コンテナホスティング
* 資料
* STNS
  * https://github.com/STNS/STNS
* ロリポップ!クラウド
* コンテナエンジン
  * haconiwa(https://github.com/haconiwa/haconiwa)
    * mrubyベース
* コンテナ
  * 構成要素
    * namespace
      * `ip a` で確認できる
      * unshareによって分離
    * chroot
      * ディレクトリ?自体を分離してリソースを見えなくする
      * chroot container によって分離
    * cgroup
      * プロセスごとにCPUやメモリを制限
    * capability
      * 権限管理の仕組み
      * getpcaps [process id] で確認
  * コンテナはプロセス
  * 参考資料: LXCで学ぶコンテナ入門
* Fast Container
  * リクエスト契機でContainerを立ち上げる (Lambdaに近い)
  * 一定時間経過でContainerを落とす
  * FastCGIが元イメージ
* netns追加に時間がかかる問題
  * ipコマンドで追加
  * 調査
    * strace
      * システムコールをトレース
      * -fcで統計を取ることができる
    * perf
      * シンボルレベルまでボトルネックを特定
      * slabが原因
      * echo 3 をdrop_cacheに書き込む
  * 原因
    * slab自体に過剰メモリを利用すると初期化自体がボトルネックになる
* Linux Bridgeの上限に引っかかった
* ネットワークコマンド起因
  * cadvisor
    * コンテナの情報を取得するツール
  * メモリが枯渇=>openが支配的
  * gdbを利用してデバッグ
    * バックトレースをとることができる
  * Linuxのソースコードは配布されているものと乖離している可能性がある
* ボトルネックを追いかけ続けて、問題を見つけたあと、解決する必要があるのかを再度考えたほうが良い
* 手に馴染むツールを作ることで、調査を進めることができる
  * pref
  * starce
  * gdb
  * valgrind
    * メモリのアロケーションを見ることができる
* abかけたら関係ないサイトが落ちた問題
  * PHPで(NAS)小さいファイルを大量にオープンするのは相性が悪い
  * OPcacheを利用して解消

### ランチ voyage

### ブログサービスのHTTPS化
* 資料
  * https://speakerdeck.com/aereal/the-construction-of-large-scale-tls-certificates-management-system-with-aws
* はてなブログの常時HTTPS配信
* 複雑なバッチの構築
* はてなブログPro
  * 独自ドメインを持てる
* Let's Encrypt
  * 大量のTLS発行ができるようになる
  * ACMI challenge
* 問題点
  * 万単位のTLS証明書の管理運用
* 解決案
  * proxyに万単位の証明書を読み込ませる
    * メモリ使用量が増加
    * proxyの再起動にも時間がかかる
  * SAM (Subject Alternative Names)
    * 1つの証明書に複数ドメインを紐付ける
    * NG
    * DNS設定は各ユーザーごと
* 方針
  * リクエストごとに都度証明証を読み込む
  * 複数台proxyに対応
  * ローカルキャッシュ
* 非機能要件
  * 発行に失敗し続けるとブログが閲覧できなくなる
  * ドメインの削除
    * LEにtime windowごとの失敗上限がある
    * 放置するとAPI limitにあたってしまう
  * ドメイン数に対してスケールする
* 配信システム
  * はてなブログ -> ngx_mruby -> cache gateway -> DynamoDB
  * cache gateway
    * AWS API→ HTTP API
    * memcachedからの読み込み、書き込み
      * DynamoDBへのリクエストを減らしてレイテンシを下げた
* 証明書発行
  * AWS StepFunctions
    * JSONでstateを記述できる
    * Lambdaをつないで実行したりする
    * リトライ/エラー処理もできる
  * AWS Lambda
    * cert-updater-function: 発行 + DynamoDBへ書き込み
    * 結果をはてなブログ側へ返す
* 証明書更新
  * DynamoDBのTTL Triggerを利用
    * このTriggerでLambdaを経由して、Step Functionsを起動
    * TTL Trigger
      * 期限切れで削除した結果をstreamとして流している
  * Step Functions
    * Choicesを利用して分岐させられる
* 巨大なバッチの難しさ
  * 実行ステップ全容を把握することの難しさ
  * 処理単位が大きくなりがち
* ピタゴラスイッチ
  * ワークフローエンジンの導入
  * pub/subモデルの利用
  * 分割統治
    * なぜできないのか?
      * composableに作れていない
  * 局所状態を持たない
    * グローバルなただ1つの状態を持つ => グローバル変数?
    * 各ステップは同じ状態をうけて、同じ状態を返す?(中身は変える)


### 業務でOSS
* OSSの原則
  * 不具合や脆弱性は自分で直す
* 書いたパッチは本家にマージ
* 趣味OSS
  * 業務時間を趣味OSSに使ってよいのか..
  * 会社のコードを混ぜたくない
* OSSポリシー
  * [oss-policy](https://cybozu-oss-policy.readthedocs.io/ja/latest/)
  * [blog](https://blog.cybozu.io/entry/oss-policy)
  * 基本的には上司の支持がなければ自分のもの
  * 趣味のOSSのパッチを取り込む際に著作権は個人の著作物になる
  * 新規OSSの許可は不要
  * パッチ提供は誰の許可も不要 (基本自分で書く)
* OSS定義
  * http://www.opensource.jp/osd/osd-japanese.html
* [伽藍とバザール](https://cruel.org/freeware/cathedral.html)

### Building selfhosted kubernetes
* 資料
  * https://gitpitch.com/nasa9084/slides/builderscon18#/
* k8s
  * インフラの抽象化基盤であることが重要
  * k8s is an microservice app.
  * components
    * etcd
      * 分散KVS
    * kubelet 
      * 各ノードに配置されてContainerを立ち上げる役割
* k8s on k8s
  * 下側のk8sは結局人間が管理?
* Self Hosted k8s
  * 例) gccはCで書かれているけど、gccのコンパイルは?
    * 小さなサブセットを作って、それをコンパイルして...
  * API Driven Update/Rollback
  * Cluster logs
  * 構成
    * k8sをサーバ上に立ち上げる
    * static podを利用してk8sを立ち上げる
    * 元k8s master nodeを消す
  * 管理は基本k8s APIでやれる
* static pod
  * kubeletが直接管理するpod
  * kube-apiserverは不要
* k8s cluster deploytool
  * self host対応
    * Tectonic
      * CoreOS系?
    * kubeadm
      * 公式が出しているがまだβ版


### Extending Kubernetes
* Problems
  * dynamic, self-healing
    * 障害が起きた時にやり直さないといけない
    * スケールについてもある
    * 各アプリについて最適なわけではない
  * new api and construct
* Memcached
  * memcachedのデプロイ
  * シャーディングやキーのハッシュ化が行われている
  * client側でIPアドレスリストが必要
    * client load balancingのため
  * proxyがConfigMapをみる
  * スケールダウンした場合
    * v2 ConfigMapを作成
    * Deploymentを更新
    * ローリングアップデート(再起動)
* Quick K8s API
  * REST API or gRPC
  * API Object
    * API Version
    * Kind
    * MetaData
      * Owener References
* The Spoke and the Wheel
  * ↑車輪のようなロゴ
    * API Serverを中心にまわりをClientが囲んでいるイメージ
  * Controllers
    * stateを宣言する
    * controllerは望ましいstateを実現する役割
    * observe(watch) -> diff -> act(update) のループを回す
* Kubernetes Controller
  * 各Controllerが連携して実行される
    * 本当にこれ
  * Deployment Controller -> ReplicaSet Object作成
  * DeploymentとReplicasetは違う
    * アップデート形式などの情報をDeploymentは持つ
  * Deployment Controller -> ReplicaSet Controller -> Scheduler (Node) -> kubelet
* Extending k8s
  * Need
    * store state - Data
    * something - Logic
  * Custom Resource Definition (Data)
  * Controller (Logic)
    * REST API的にやりとりできれば何でもOK
    * client-go
      * Controller Library
* operator-sdk
  * demo
  * https://github.com/operator-framework/operator-sdk

### LT
* QMK
  * firmwareを書き込めるキーボードが必要
* WebReplay
* https://nedi.app/