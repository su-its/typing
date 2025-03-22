# docker ディレクトリ

- **docker-bake.hcl** Docker イメージをビルドするための Bake file
- **compose.yaml** デプロイ先のサーバで実際にサービスを起動するための構成を記述したファイル

手元でやるときはこのディレクトリで以下のようにすると本番サーバの様子をそこそこ再現できます。

```bash
# typing-appおよびtyping-serverを本番用にビルドするための環境変数等は
# セットされているものとします。
# 細かい内容はそれぞれのドキュメントやソースコードを参照してください。
docker buildx bake --allow=fs.read='*' --file docker-bake.hcl
docker compose -p typing -f compose.yaml up -d
docker compose -p typing -f compose.yaml logs -f # ログを tail
```

1つ目のコマンドはイメージをビルドします。あらかじめ `export TAG=v1.0.0` のようにして環境変数 `TAG` をセットしておいてから実行するとイメージにそのタグがつきます。環境変数 `TAG` の有無にかかわらず latest タグのついたイメージはいつもビルドされます。最新のタグは `git describe --tags --abbrev=0 main` で確認できます。2つ目のコマンドはコンテナを起動します。
