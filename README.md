# Ginについて

ginのTutorial: https://go.dev/doc/tutorial/web-service-gin

# それぞれのファイルについて

## main.go
ルーティングとかしている

## service/init.go
DB接続関連をやってくれてる

## service/book.go
controllerから振られたDB関連の処理をやってくれる

## model/book.go
* bookの構造体の定義
* マイグレーション的なこともここでやっている

## controller/book.go
* main.goから振られたリクエストをserviceに振る
* レスポンスを返す

## middleware/bookMiddleware.go
リクエストのログをとっておく

# ライブラリについて
* GORMはGo言語でよく使われているORM
* zapはuberが作ったopensourceのやつ(ログとかとってくれる)

# public とprivateについて
packageを横断した時の変数について、
頭文字が大文字のものはPublicとなり、外部packageから参照が可能に、
頭文字が小文字のものはPrivateとなり、外部packageから参照が不可能になります。