# cover.outというファイルテストのカバー内容を吐き出す
go test -cover ./... -coverprofile=cover.out

# go toolを用いてcover.htmlを作成する
go tool cover -html=cover.out -o cover.html

# cover.htmlを開く
open cover.html