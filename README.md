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

4. .envファイルをもらうか作成。[.envファイルの説明](#env)

## API 仕様書

エンドポイント、リクエストレスポンスの形式、その他情報のAPIの仕様書です。

### エンドポインツ

#### ユーザが所属する会社のユーザ一覧取得エンドポイント

- **URL:** `/users/list`
- **メソッド:** GET
- **説明:** パラメーターで"id"を指定、そのユーザが所属する会社のユーザ一覧を返す。
- **リクエスト:**
  - パラメーター:
    - `id`: (int)ID。トークンと合わせて本人のものか確認、所属している会社を特定する。
  - ヘッダー:
    - `Authorization`: (string) 認証トークン
- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
        "srvResCode": 1002,                            // コード
        "srvResMsg":  "Successfully retrieved list of users matching the criteria.", // メッセージ
        "srvResData": {
          "userList": [
            {
              "name": "hogeta piyonaka",
              "groupName": "人事部",
              "situation": "安否確認中"
            },
            {
              "name": "fugako nakapiyo",
              "groupName": "情報技術部",
              "situation": "支援必要"
            },,,
          ]
        },                         // データ
      }
      ```

  - ステータスコード: 400 Bad Request
    - ボディ:

      ```json
      {
        "srvResCode": 7002,                            // コード
        "srvResMsg":  "Incorrect request.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

  - ステータスコード: 404 Not Found
    - ボディ:

      ```json
      {
        "srvResCode": 7003,                            // コード
        "srvResMsg":  "The condition specification may be correct, but the specified resource cannot be found.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

#### ユーザの詳細情報取得エンドポイント

- **URL:** `/users/user`
- **メソッド:** GET
- **説明:** パラメーターで"id"を指定、そのユーザの詳細情報を返す。
- **リクエスト:**
  - パラメーター:
    - `id`: (int)ID。トークンと合わせて本人のものか確認、ユーザ情報を取得。
  - ヘッダー:
    - `Authorization`: (string) 認証トークン
- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
        "srvResCode": 1003,                            // コード
        "srvResMsg":  "Successful acquisition of user information.", // メッセージ
        "srvResData": {
          "userInfo": {
            "name": "hogeta piyonaka",
            "id": 1,
            "groupName": 1,  // "人事部"
            "situation": "支援必要", 
            "mailAddress": "hogeta@gmail.com",
            "address": "にほんのどこか",
            "company_no": 1,  // "AComp"
          }
        },                         // データ
      }
      ```

  - ステータスコード: 400 Bad Request
    - ボディ:

      ```json
      {
        "srvResCode": 7002,                            // コード
        "srvResMsg":  "Incorrect request.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

  - ステータスコード: 404 Not Found
    - ボディ:

      ```json
      {
        "srvResCode": 7003,                            // コード
        "srvResMsg":  "The condition specification may be correct, but the specified resource cannot be found.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

  - ステータスコード: 500 Not Found
    - ボディ:

      ```json
      {
        "srvResCode": 7014,                            // コード
        "srvResMsg":  "Failure to retrieve user data.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

#### ユーザ作成エンドポイント

- **URL:** `/users/register`
- **メソッド:** POST
- **説明:** 新規ユーザをDBに登録します。
- **リクエスト:**
  - ヘッダー:
    - `Authorization`: (string) 認証トークン
    - `Content-Type`: application/json
  - ボディ:

    ```json
    {
      "CompanyNo": 1,
      "name": "hogeta piyonaka",
      "id": 1,
      "mailAddress": "hogeta@gmail.com",
      "address": "にほんのどこか",
      "password": "C@h",
      "groupNo": 1,  // ここまで？
    }
    ```

- **レスポンス:**
  - ステータスコード: 201 Created
    - ボディ:

      ```json
      {
        "srvResCode": 1004,                            // コード
        "srvResMsg":  "Successful user registration.", // メッセージ
        "srvResData": {
          "authenticationToken": "token@h",
        },                         // データ
      }
      ```

  - ステータスコード: 400 Bad Request
    - ボディ:

      ```json
      {
        "srvResCode": 7004,                            // コード
        "srvResMsg":  "Failed to bind request JSON data.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

  - ステータスコード: 409 Conflict
    - ボディ:

      ```json
      {
        "srvResCode": 7005,                            // コード
        "srvResMsg":  "The user is already registered.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

  - ステータスコード: 500 Internal Server Error
    - ボディ:

      ```json
      {
        "srvResCode": 7006,                            // コード
        "srvResMsg":  "Some problems with db registration of new users.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

  - ステータスコード: 500 Internal Server Error
    - ボディ:

      ```json
      {
        "srvResCode": 7007,                            // コード
        "srvResMsg":  "There is already a user with the same primary key. Uniqueness constraint violation.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

  - ステータスコード: 500 Internal Server Error
    - ボディ:

      ```json
      {
        "srvResCode": 7010,                            // コード
        "srvResMsg":  "Failed to generate authentication token.", // メッセージ
        "srvResData": {},                         // データ
      }
      ```

#### ログイン認証エンドポイント

- **URL:** `/users/login`
- **メソッド:** POST
- **説明:** ログイン処理をする。
- **リクエスト:**
  - ヘッダー:
    - `Content-Type`: application/json
  - ボディ:

    ```json
    {
      "mailAddress": "hogeta@gmail.com",
      "password": "C@h"
    }
    ```

- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
        "srvResCode": 1005,                            // コード
        "srvResMsg":  "Successful login.", // メッセージ
        "srvResData": {
          "authenticationToken": "token@hogeta"
        },                         // データ
      }
      ```

  - ステータスコード: 500 Internal Server Error
    - ボディ:

      ```json
      {
        "srvResCode":7009,                    // コード
        "srvResMsg":  "User not found.", // メッセージ
        "srvResData": {}// データ
      }  
      ```

  - ステータスコード: 500 Internal Server Error
    - ボディ:

      ```json
      {
        "srvResCode":7010,                    // コード
        "srvResMsg":  "Password does not match.", // メッセージ
        "srvResData": {}// データ
      }  
      ```

  - ステータスコード: 500 Internal Server Error
    - ボディ:

      ```json
      {
        "srvResCode":7011,                    // コード
        "srvResMsg":  "Failure to obtain user ID.", // メッセージ
        "srvResData": {}// データ
      }  
      ```

  - ステータスコード: 500 Internal Server Error
    - ボディ:

      ```json
      {
        "srvResCode":7012,                    // コード
        "srvResMsg":  "Failed to generate authentication token.", // メッセージ
        "srvResData": {}// データ
      }  
      ```

### その他のエンドポインツ

#### トップサイトを返すエンドポイント

- **URL:** `/`
- **メソッド:** GET
- **説明:** トップサイトを返す。
- **リクエスト:**
  - パラメーター:
  - ヘッダー:
- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:
      トップサイト。

#### テスト用のJSONを返すデバッグエンドポイント

- **URL:** `/test/json`
- **メソッド:** GET
- **説明:** 疎通実装テスト用のJSONを返す。
- **リクエスト:**
  - パラメーター:
  - ヘッダー:
- **レスポンス:**
  - ステータスコード: 200 OK
    - ボディ:

      ```json
      {
        "srvResCode": 1001,                                 // コード
        "srvResMsg":  "JSON for testing.",      // メッセージ
        "srvResData": {"message": "hello go server!"}, // データ
      }
      ```

#### テスト用のPOST送信を鯖側が受信できてるかテストするためのデバッグエンドポイント

- **URL:** `/test/cfmreq`
- **メソッド:** POST
- **説明:** テスト用のPOST送信をサーバー側で受信確認する。
            ボディはJSONを送るが、サーバー側でデバッグ出力して確認するだけなので内容は何でもいい。
- **リクエスト:**
  - ヘッダー:
  - ボディ:

    ```json
    {

    }
    ```

- **レスポンス:**

### 認証

このAPIエンドポイントは認証トークンが不必要です。  
ただし、以降のアプデでログイン時にトークンを発行し、以降のリクエストに付与する仕様に変更する可能性があります。  
(認証トークンは `Authorization` ヘッダーに含まれる必要があります。)

認証が失敗した場合のレスポンス:  

- ステータスコード: 401 Unauthorized
  - ボディ:

    ```json
    {
      "srvResCode": 7001,                            // コード
      "srvResMsg":  "Authentication unsuccessful.", // メッセージ
      "srvResData": {},                         // データ
    }
    ```

## エラー処理

APIがエラーを返す場合、詳細なエラーメッセージが含まれます。エラーに関する情報は[サーバーエラーコード](#server-error-code)を参照してください。　　

## SERVER ERROR CODE

サーバーレスポンスコードとして"srvResCode"キーで数値を返す。  
以下に意味を羅列。  

- 成功関連
  - 1001: JSON for testing.  
    テスト用JSON。
  - 1002: Successfully retrieved list of users matching the criteria.  
    条件に合うユーザー一覧の取得に成功。
  - 1003: Successful acquisition of user information.  
    ユーザー情報の取得に成功。
  - 1004: Successful user registration.  
    ユーザー情報の登録に成功。
  - 1005: Successful login.  
    ログインに成功。

- エラー関連
  - 7001: Authentication unsuccessful.  
    認証トークンが不正。
  - 7002: Incorrect request.  
    リクエストが正しくない。
  - 7003: The condition specification may be correct, but the specified resource cannot be found.  
    条件指定は正しい可能性があるが、指定されたリソースが見つからない。
  - 7004: Failed to bind request JSON data.  
    POSTリクエストボディのGO構造体へのバインドが失敗。
  - 7005: The user is already registered.  
    すでにユーザーが登録されている。
  - 7006: Some problems with db registration of new users.
    新規ユーザのDB登録になんらかの問題が発生した。
  - 7007: There is already a user with the same primary key. Uniqueness constraint violation.
    同じ主キーを持つユーザがすでに存在します。一意性制約違反。  
  - 7008: Authentication unsuccessful. Failed to parse token.
    認証に失敗。トークンの解析に失敗。  
  - 7009: User not found.  
    ユーザーが見つからない。  
  - 7010: Password does not match.
    パスワードが一致しない。
  - 7011: Failure to obtain user ID.
    ユーザIDの取得に失敗。
  - 7012: Failed to generate authentication token.
    認証トークンの生成に失敗。
  - 7013: The id is not stored.  
    トークンから取得したidが保存されていない。
  - 7014: Failure to retrieve user data.
    ユーザデータの取得に失敗。

## .ENV

.evnファイルの各項目と説明

```env:.env
MYSQL_USER=DBに接続する際のログインユーザ名
MYSQL_PASSWORD=パスワード
MYSQL_HOST=ログイン先のDBホスト名。dockerだとサービス名。
MYSQL_PORT=ポート番号。dockerだとコンテナのポート。
MYSQL_DATABASE=使用するdatabase名
JWT_SECRET_KEY="openssl rand -base64 32"で作ったJWTトークン作成用のキー。
TOKEN_LIFETIME=JWTトークンの有効期限
```

## 開発者

- Author:[unSerori]
- Mail:[x]
