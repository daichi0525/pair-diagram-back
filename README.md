# pair-diagram-back

## goポイント
https://qiita.com/k-penguin-sato/items/62dfe0f93f56e4bf9157

## go init

https://qiita.com/TakanoriVega/items/6d7210147c289b45298a

```
go mod init github.com/daichi0525/pair-diagram-back.git

```

## DB接続クライアントツール
https://qiita.com/okuzumi_slj/items/e2fabf5b642d8889ed3f

## フレームワークのインストール

### gorm

```
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

https://zenn.dev/yasu2122/articles/045183cd36add5

(インストールの内容も載ってる)
https://qiita.com/katsuomi/items/d1e6625ae9a5b663e11f

### gin
https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a

```
go get github.com/gin-gonic/gin
```
### CORS

```
go get github.com/rs/cors
go get github.com/gin-contrib/cors

```

### 起動方法
Docker
```shell
# PostgreSQLの起動
docker-compose up -d

# Goの開始
go run main.go
```