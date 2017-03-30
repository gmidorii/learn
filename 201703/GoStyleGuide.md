# GoStyleGuide

## Package Organaization
- 複数のファイルに分割する
- typeは利用場所の近くに配置する
  - 必ずtypeを定義する必要はない
  - ただfileのトップのcoreのタイプを置くことがgood practice
- 機能的な責任によってパッケージを構成する
  - modelパッケージを作るのでなく利用されるレイアーの近くで定義する
- godocに最適化させる
- gapは例を用いることで埋める
  - [Go Example]{https://blog.golang.org/examples}
- mainパッケージからエクスポートしない
