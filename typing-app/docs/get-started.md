# 開発環境構築

## 1. リポジトリのクローン

```bash
$ git clone https://github.com/su-its/typing.git
$ cd typing/typing-app
```

## 2. パッケージのインストール(npm ではなく bun を採用しています)

### Bun をインストールする (まだインストールしていない場合)

```bash
$ curl -fsSL https://bun.sh/install | bash
```

### プロジェクトの依存パッケージをインストール(typing-app ディレクトリで実行)

```bash
$ bun install
```

## 3. API & DB の起動(typing-server ディレクトリで実行)

```bash
$ docker-compose --file docker-compose.dev.yml up --build
```

## 4. フロントエンドの起動(typing-app ディレクトリで実行)

```bash
$ bun dev
```
