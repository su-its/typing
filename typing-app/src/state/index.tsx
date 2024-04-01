"use client";
import { ReactNode, createContext, useContext } from "react";

// TODO: あとで直す。適当な型であることを明示するためにあえてへんな名前にしてる
type UserType001 = {
  student_number: string;
  handle_name: string;
};

export const LoginContext = createContext<UserType001 | undefined>(undefined);

export function useUser() {
  return useContext(LoginContext);
}

type LoginProviderProps = {
  children: ReactNode;
  user?: UserType001;
};

/**
 * Server Component で Context が使えないのでその children の Client Component で `useUser()` できない。
 * これをできるようにするための Client Component。`LoginContext.Provider` を挟むだけ。
 * - https://nextjs.org/docs/app/building-your-application/rendering/composition-patterns#using-context-providers
 * - https://future-architect.github.io/articles/20231214a/
 */
export function LoginProvider({ children, user }: LoginProviderProps) {
  return <LoginContext.Provider value={user}>{children}</LoginContext.Provider>;
}
