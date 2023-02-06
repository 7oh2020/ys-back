# ys-back

これは夢色水車(https://ys.7oh.dev/)のバックエンドです。
Twitter API v2 を利用して要望ツイートをデータベースに収集します。

## 特徴

- DevContainer を使用して Go + PostgreSQL 環境をコード化しています
- アーキテクチャにレイヤードアーキテクチャを採用しています

## 環境変数の設定

起動には動作環境に合わせたデータベース接続情報やBearer Tokenなど環境変数が必要です。
.devcontainer/.env.example を参考にして.devcontainer/.env ファイルを作成してください。

- 開発時はPostgreSQLへの接続情報はデフォルトのままで動作します。本番環境の場合は適宜書き換えてください
- TWITTER_BEARER(Twitter API v2 の Bearer Token) は[Twitter の開発者向けページ](https://developer.twitter.com/en/apps)で作成できます。

## コンテナの起動

前提として VS Code に DevContainer 拡張機能がインストールされている必要があります。

1. VS Code でこのプロジェクトを開きステータスバーのリモートボタン > Reopen in Container と進みます
2. コンテナがビルドされます
3. コンテナが起動します

## 開発 サーバーの起動

初回は以下のコマンドを実行するだけで開発サーバーが起動します:
起動後はホストマシンの localhost:8080 でアクセス可能になります。

```
make
```

開発サーバーの起動のみを行うコマンドは以下の通りです。
起動後はコードの変更字に自動で再ビルドされます。

```
make dev
```

ビルドを実行するコマンドは以下の通りです。
dist ディレクトリに実行ファイルが出力されます。

```
make build
```

## API ルート

CRON 等を使用して GET /needs/update に定期的にアクセスすることでツイートがデータベースに蓄積されます。
GET /tweet/meta で各種件数が取得できるので Next.js の GetStaticPaths 関数に利用できます。

- GET /needs/update -> 要望ツイートの収集を行います
- GET /tweet/index?tag_id=[tagID]&page=[page] -> タグ ID が一致するツイートを取得します。ページングにより一定数毎に取得されます
- GET /tweet/meta -> 総件数、タグ毎の件数などのメタ情報が取得できます
