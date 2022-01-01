# Ginについて

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