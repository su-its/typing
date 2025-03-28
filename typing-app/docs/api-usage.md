# API サーバとのやりとりについて

どのように API サーバとやり取りする**べき**かを説明します。

本アプリケーションでは [`libs/api`](../src/libs/api/) ディレクトリから利用できる API サーバの型定義ファイルを利用して API サーバとやり取りする**べき**です。

例えば、

```tsx
import { client } from "@/libs/api";

async function SomePageToGetRanking() {
  const { data, error } = await client.GET("/scores/ranking");

  if (error) {
    return <div>Error</div>;
  }

  return (
    <div>
      ranking data: {JSON.stringify(data)}
    </div>
  );
}
```

のようにします。

> :memo: [OpenAPI について](../../typing-server/docs/about-openapi.md) も合わせて読んでください。
