# 💻 開発環境のセットアップ

このガイドでは、`typing` プロジェクトの環境をセットアップする手順を説明します。

---

## 1. リポジトリのクローン

まず、リポジトリをクローンし、`typing-app` ディレクトリに移動します。

```bash
git clone https://github.com/su-its/typing.git
cd typing/typing-app
```

---

## 2. パッケージのインストール（Yarn を使用）

このプロジェクトでは **npm ではなく Yarn** を使用します。

### ✅ 依存パッケージのインストール（`typing-app` ディレクトリ内で実行）

```bash
corepack enable yarn
yarn
```

### ⚠️ 注意事項（Windows環境）
Windows環境では、`cmd` または `PowerShell` を **管理者権限で実行** する必要がある場合があります。正常に動作しない場合は、管理者権限でターミナルを実行してください。

### 🛠 WSL ユーザー向けの追加設定
WSL 環境では、Node.js のバージョンによって `corepack` が正常に動作しない場合があります。問題が発生する可能性がある場合は、以下のコマンドで `nvm` をインストールし、Node.js をセットアップしてください。

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
nvm install node
```

---

## 3. API & データベースの起動（`typing-server` ディレクトリで実行）

バックエンドの API サーバーとデータベースを起動します。

```bash
docker compose --file docker-compose.dev.yml up --build
```

---

## 4. フロントエンドの起動（`typing-app` ディレクトリで実行）

開発環境を起動するには、以下のコマンドを実行してください。

```bash
yarn dev
```

---

この手順で、開発環境が準備完了です！ 🎉

