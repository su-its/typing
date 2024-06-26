# デプロイ手順

## GitHub Actions を利用したデプロイ

手動で workflow を実行すると(学内のサーバで) [deploy.yml](../.github/workflows/deploy.yml) に書かれた内容が実行されます。

**前提条件:** 必要な権限のあるユーザで GitHub にログインしている

1. リポジトリの Actions タブへ移動する
2. 左のメニューから Deploy Containers というワークフローを選択、main ブランチで実行(Run Workflow)する

<!-- put image here -->

## SSH でサーバにログインして行うデプロイ

できれば[#GitHub Actions を利用したデプロイ](#GitHub Actions を利用したデプロイ)を実施してください。
どうしてもうまくいかない場合はこちらの方法を試してください。

**前提条件:** github ユーザに ssh でログインしている。以下、全て SSH 先のシェルでの操作です。
> github ユーザには安全のため sudo 権限を与えていません。注意してください。

移動
```bash
cd ~/actions-runner/_work/typing/typing/docker
```

ソースコードを pull

```bash
git switch main
git fetch origin main
git reset --hard origin/main
```

[`COMPOSE_PROJECT_NAME`](https://docs.docker.com/reference/cli/docker/compose/#use--p-to-specify-a-project-name) を `typing` に設定

```bash
export COMPOSE_PROJECT_NAME=typing
```

(オプション)正常に動いてることを確認

```bash
curl http://localhost:8080/health
curl http://localhost
```

(オプション)コンテナの様子を確認(UP になってるか)

```bash
podman ps
```

コンテナを作成し直すために一度コンテナを削除

```bash
podman-compose down app
podman-compose down api
```

(オプション)無事に削除できたことを確認(先ほどあったコンテナが表示されなくなってるか)

```bash
podman ps
```

最新のイメージを pull し、コンテナを作成

```bash
podman-compose pull
podman-compose up -d app
podman-compose up -d api
```

> (余談)イメージは su-its org の [packages](https://github.com/orgs/su-its/packages) にある

(オプション)サービスが起動したことを確認

```bash
curl http://localhost:8080/health
curl http://localhost
```

(オプション)コンテナの様子を確認(STATUS が UP になってるか)

```bash
$ podman ps

CONTAINER ID  IMAGE                             COMMAND               CREATED       STATUS                 PORTS                   NAMES
b14f0aa19ef9  docker.io/library/mysql:8.3.0     mysqld                2 months ago  Up 28 hours (healthy)                          typing_db_1
f71389e573c4  ghcr.io/su-its/typing-app:dev     node server.js        2 months ago  Up 28 hours            0.0.0.0:3000->3000/tcp  typing_app_1
73178f42db34  ghcr.io/su-its/typing-server:dev  ./server -seed fa...  2 months ago  Up 28 hours (healthy)  0.0.0.0:8080->8080/tcp  typing_api_1
```

> イメージのタグ(`ghcr.io/su-its/typing-app:` 以降の部分)は場合によっては上記と違うかもしれない