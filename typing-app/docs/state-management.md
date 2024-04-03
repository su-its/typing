# 状態管理

## ログイン状態

- ログインしているかどうか判断したいとき

  ```tsx
  import { getCurrentUser } from "@/app/actions";

  export async function LoginStatus() {
    const user = await getCurrentUser();

    return <div>User is{user ? " " : " not "}logged in!</div>
  }
  ```
