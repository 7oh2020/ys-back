FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye

# WEBサーバーのライブリロードのためにcosmtrek/airコマンドをインストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
