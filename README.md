# カレンダーアプリ　backend



サマーインターンを終えて、以前作成したカレンダーアプリのバックエンド(apiサーバー)に修正したい部分がたくさん見つかったので、学習もかねて改善したいと思ってはじめました。

修正前→(https://github.com/f-masanori/calendar_app)　[readmeが汚いです]

### 目的

学習をメインの目的として開発し始めました。

### 技術選定

学習のため、これまで自分が使ったことがなかった言語やツールを使用して開発しました。

- 言語
  - golang
    - 静的型付き言語(好み、読みやすい)
    - 標準ライブラリだけでサーバーを構築できるので、学習しやすそう
- 開発環境
  - Docker
    - ローカル環境を汚さない
    - デプロイが簡単になる
- 設計
  - クリーンアーキテクチャライクなもの
    - 挑戦してみたかった

### セットアップ

1. Makefile内の`GOOGLE_APPLICATION_CREDENTIALS=$(HOME)/.config/gcloud/calendar_service_account.json`を自身のfirebase プロジェクトのjsonファイルのパスに書き換える

2. `make dcu`で実行



### 前回からの改善点[改善予定のものも含む]

1. reflexを導入して開発の効率を上げた
2. サマーインターン先でのMakefile文化を真似して、MakefileにDB接続先情報やfirebase設定情報を書いた（設定情報がプロジェクトのルートに来るので読みやすくなった）
3. ミドルウェアをライブラリaliceを使ってまとめて、可読性アップ
4. プレゼンテーター機能をinterface層に記述していたのを、infrastructure層に移動して、責務の分散する[改善予定]
5. マイグレーション機能を一つのコンテナとして独立させる。[改善予定]（前回はカレンダーアプリの中でマイグレーションするようにしていたが、毎回コンテナの中に入って実行するのが面倒だった。）
6. Rest API化する。[改善予定]（（規則を作ることで可読性アップ）