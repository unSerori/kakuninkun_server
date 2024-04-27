# kakuninkun_server


## SERVER ERROR CODE
サーバーレスポンスコードとして"srv_res_code"キーで数値を返す。  
以下に意味を羅列。  
"  
srv_res_code: srv_res_mgs:  
mean  
"  の形式で記述。  

- 成功関連
    - 1001: Successful user registration.  
    ユーザー登録成功
- エラー関連
    - 7001: Failed to bind request JSON data.:  
    POSTリクエストボディのGO構造体へのバインドが失敗
