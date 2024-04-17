# docker ディレクトリ

- **compose.ci.yaml** Docker イメージをビルドするためのレシピ
- **compose.yaml** デプロイ先のサーバで実際にサービスを起動するための構成を記述したファイル

手元でやるときはこのディレクトリで以下のようにすると本番サーバの様子をそこそこ再現できます。
```bash
docker compose build -f compose.ci.yaml
docker compose -f compose.yaml up -d
# docker compose -f compose.yaml logs -f # ログを tail
```