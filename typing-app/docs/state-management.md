# 状態管理

## ログイン状態

- ログインしているかどうか判断したいとき

  ```tsx
  "use client";
  import { useUser } from "@/state";

  export function LoginStatus() {
    const user = useUser()

    return <div>User is{user ? " " : " not "}logged in!</div>
  }
  ```
