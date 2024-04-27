# kakuninkun_server

## 概要

安否状況確認アプリのAPIサーバ。  

### 環境

Visual Studio Code: 1.88.1  
golang.Go: v0.41.4  
image Golang: go version go1.22.2 linux/amd64  

## 環境構築

1. 以下のDocker環境を作成  
[リポジトリURL](https://github.com/unSerori/docker_kakuninkun_server)  
SSH URL:  

    ```SSH:SSH URL
    git@github.com:unSerori/docker_kakuninkun_server.git
    ```

2. ここまでが1の内容（フォルダーをVScodeで開きgo_serverをVScodeアタッチ。）
3. shareディレクトリ内で以下のコマンド。

    ```bash:Build an environment
    # vscode 拡張機能を追加
    cat vscode-ext-base.txt | while read line; do code --install-extension $line; done
    # Goのライブラリインストール
    go install -v github.com/go-delve/delve/cmd/dlv@latest
    ```

4. .envファイルをもらうか作成。

## API仕様

工事予定  

## SERVER ERROR CODE

サーバーレスポンスコードとして"srv_res_code"キーで数値を返す。  
以下に意味を羅列。  

- 成功関連
  - 1001: Successful user registration.  
    ユーザー登録成功
- エラー関連
  - 7001: Failed to bind request JSON data.:  
    POSTリクエストボディのGO構造体へのバインドが失敗
  