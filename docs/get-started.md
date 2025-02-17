# 💻 Get Started
## 1. リポジトリのクローン

```bash
git clone https://github.com/su-its/typing.git
cd typing/typing-app
```

## 2. パッケージのインストール(npm ではなく yarn を採用しています)

### プロジェクトの依存パッケージをインストール(typing-app ディレクトリで実行)

Windows環境の場合、`cmd`または`PowerShell`を**管理者権限で**実行しなければ実行できない場合があります  
正常に動作しない際は、ターミナルを**管理者権限**で実行することを検討してください

**WSLを使っている人は以下も参照**
<details>
  nodeのバージョンの影響でcorepackが使えないことがあるので，自信がない人は以下のコマンドを実行することを推奨
  ```bash
  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
  nvm install node
  ```
</details>

```bash
corepack enable yarn
yarn
```


## 3. API & DB の起動(typing-server ディレクトリで実行)

```bash
docker compose --file docker-compose.dev.yml up --build
```

## 4. フロントエンドの起動(typing-app ディレクトリで実行)

```bash
yarn dev
```
