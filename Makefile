#!/usr/bin/make

APPNAME="ys-back"
OUTDIR="./dist"

# 初回セットアップ & 開発サーバー起動
all: migrate_up gen dev

# ディレクトリのクリーンアップ
clean:
	@rm -rf $(OUTDIR)
	@mkdir $(OUTDIR)

# 依存パッケージのインストール
depend:
	@go mod tidy

# Twitter APIの疎通確認
search_test: depend fmt
	@go run ./src/cmd/search_test/main.go

# データベースのマイグレーション
migrate_up: depend fmt
	@go run ./src/cmd/migrate_up/main.go

# cosmtrek/airを使用したライブリロードの開始。コードを変更すると自動で再ビルドされる
dev: depend fmt
	@air

# WEBサーバーの起動
run: depend fmt
	@go run ./src/cmd/start/main.go

# コードをフォーマット
fmt:
	@go fmt ./...

# ビルド
build: clean depend fmt test
		@go build -o $(OUTDIR)/$(APPNAME) ./src/cmd/start/main.go
		@echo "done build build"

# ユニットテスト
test: depend fmt
	@go test -short -cover ./...
	@echo "done unittest"

# コードの自動生成
gen: gen_mock

# モック生成
gen_mock: depend fmt
	@go install github.com/vektra/mockery/v2@latest
	@mockery --all --recursive=true
	@go mod tidy
	@echo "done generate mocks"
