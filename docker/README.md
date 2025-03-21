# docker ディレクトリ

- **docker-bake.hcl** Docker イメージをビルドするための Bake file
- **compose.yaml** デプロイ先のサーバで実際にサービスを起動するための構成を記述したファイル

手元でやるときはこのディレクトリで以下のようにすると本番サーバの様子をそこそこ再現できます。

```bash
docker buildx bake --allow=fs.read='*' --file docker-bake.hcl
docker compose -f compose.yaml up -d
# docker compose -f compose.yaml logs -f # ログを tail
```

1つ目のコマンドはイメージをビルドします。`TAG=v1.0.0 docker...` のように、あらかじめ環境変数 `TAG` をセットしておくとビルドされるイメージにそのタグがつきます。環境変数 `TAG` の有無にかかわらず latest タグのついたイメージはいつもビルドされます。2つ目のコマンドはコンテナを起動します。
