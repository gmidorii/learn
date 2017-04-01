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

## Package Naming
- lowercase only
- 短く、ユニークで代表的な名前
  - パッケージ名から目的を把握できるくらいに
- util等の広範なパッケージ名は避ける
- 悪い名前しかつけられない場合は、全体構成が誤りな可能性が高い
- import pathをきれいに保つ
  - src,pkg等は使わない
- 複数形にはしない
- goのスタンダードパッケージをrenameする場合は、goプレフィックスをつける

## Package Documentation
- packageの前にpackageの説明を記載する
  - `“Package {pkgname}”`
-  `doc.go` にpackage情報を記載する手法もある