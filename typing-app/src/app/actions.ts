"use server";
import { client } from "@/libs/api/server-side-client";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { User } from "@/types/user";

type LoginActionState = {
  error?: string;
};

export async function login(_: LoginActionState, formData: FormData): Promise<LoginActionState> {
  const studentNumber = formData.get("student-number")!.toString();

  try {
    const { data, error } = await client.GET("/users", {
      params: {
        query: {
          student_number: studentNumber,
        },
      },
    });
    if (error || !data) {
      if (/not found/.test(`${error}`.toLowerCase())) {
        return { error: "見つかりませんでした" };
      }
      return { error: "もう一度お試しください" };
    }

    const expires = new Date(Date.now() + 3 * 60 * 60 * 1000);

    const user: User = {
      id: data.id,
      handleName: data.handle_name,
      studentNumber: data.student_number,
    };

    (await cookies()).set("user", JSON.stringify(user), { expires, httpOnly: true });
  } catch (error) {
    return { error: "通信に失敗しました" };
  }

  redirect("/game"); // NOTE: try-catch外でリダイレクトする
}

export async function logout() {
  (await cookies()).set("user", "", { expires: 0 });
}

export async function getCurrentUser() {
  const userStr = (await cookies()).get("user")?.value;
  if (!userStr) return undefined;

  function isValidUser(o: any): o is User {
    return o && typeof o.id === "string" && typeof o.studentNumber === "string" && typeof o.handleName == "string";
  }

  const user = JSON.parse(userStr) as User;
  return isValidUser(user) ? user : undefined;
}
