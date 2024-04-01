.

├── api
│   ├── cmd        - アプリケーションのエントリーポイントとなるmainパッケージを格納するディレクトリ
│   ├── handler    - APIのエンドポイントハンドラーを実装するパッケージを格納するディレクトリ
│   ├── repository - データ永続化層とのインターフェースを定義するパッケージを格納するディレクトリ
│   ├── router     - APIルーティングを定義するパッケージを格納するディレクトリ
│   └── service    - アプリケーションのビジネスロジックを実装するパッケージを格納するディレクトリ
├── domain
│   ├── model
│   └── repository     - データ永続化層の実装を格納するディレクトリ
│       └── ent
│           └── schema - データベーススキーマの定義を格納するディレクトリ
├── Dockerfile
├── Dockerfile.dev
├── docker-compose.dev.yml
├── docker-compose.yml
├── go.mod
├── go.sum
└── openapi.yaml