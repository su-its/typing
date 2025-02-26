# タイピング練習アプリ バックエンドサーバー

このディレクトリには、静岡大学の新入生向けタイピング練習アプリケーションのバックエンドサーバーのコードが含まれています。

## 技術スタック

- **言語**: Go
- **O/Rマッパー**: ent
- **データベース**: MySQL
- **API定義**: OpenAPI (openapi.yaml)
- **コンテナ化**: Docker, Docker Compose

## バックエンドの開発環境セットアップ

### 前提条件

- Git
- Go (go.modで指定されているバージョン)
- Docker、Docker Compose
- Make

### セットアップ手順

1. go.modで指定されているGolangのバージョンをインストールしてください。

2. 必要なパッケージのインストール:

```bash
go mod tidy
```

3. entの生成ファイルを作成:

```bash
make generate
```

4. Dockerコンテナのビルドと起動:

```bash
make up-with-build
```

そのほかのコマンドは `Makefile` を参照するか `make help` を実行してください。

## ディレクトリ構造

```
typing-server/
├── api/            # API関連のコード
├── cmd/            # アプリケーションのエントリーポイント
├── config/         # 設定ファイル
├── docs/           # ドキュメント
├── internal/       # 内部パッケージ（プロジェクト内でのみ使用）
│   ├── handler/    # HTTPハンドラー
│   ├── repository/ # データベースとのやり取りを行うレイヤー
│   ├── usecase/    # ビジネスロジック
│   └── ent/        # ORマッパーの生成コード
└── pkg/            # 外部からインポート可能なパッケージ
```

クリーンアーキテクチャとディレクトリ構造の詳細は `docs/directory-strategy.md` を参照してください。

## API仕様

API仕様は `openapi.yaml` ファイルに定義されています。これを使用してクライアントとサーバー間のインターフェースを確認できます。

APIリクエストの例や詳細な説明は `docs/about-openapi.md` を参照してください。

## テスト

テストの作成方法や実行方法の詳細は `docs/how2test.md` を参照してください。