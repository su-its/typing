# 💻 Get Started
## 1. リポジトリのクローン

```bash
$ git clone https://github.com/su-its/typing.git
$ cd typing/typing-app
```

## 2. パッケージのインストール(npm ではなく yarn を採用しています)

### プロジェクトの依存パッケージをインストール(typing-app ディレクトリで実行)

```bash
$ corepack enable yarn
$ yarn
```

## 3. API & DB の起動(typing-server ディレクトリで実行)

```bash
$ docker-compose --file docker-compose.dev.yml up --build
```

## 4. フロントエンドの起動(typing-app ディレクトリで実行)

```bash
$ yarn dev
```
