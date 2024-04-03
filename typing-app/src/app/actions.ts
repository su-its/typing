"use server";
import { client } from "@/libs/api";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { User } from "@/types/user";

type LoginActionState = any;

export async function login(_prevState: LoginActionState, formData: FormData) {
  const studentNumber = formData.get("student-number")!.toString();

  const { data, error } = await client.GET("/users", {
    params: {
      query: {
        student_number: studentNumber,
      },
    },
  });

  if (error) {
    return error;
  }

  const expires = new Date(Date.now() + 3 * 60 * 60 * 1000);

  const user: User = {
    id: data.id!,
    handleName: data.handle_name!,
    studentNumber: data.student_number!,
  };

  cookies().set("user", JSON.stringify(user), { expires, httpOnly: true });

  redirect("/game");
}

export async function logout() {
  cookies().set("user", "", { expires: 0 });
}

export async function getCurrentUser() {
  const userStr = cookies().get("user")?.value;
  if (!userStr) return null;

  function isValidUser(o: any): o is User {
    return o && typeof o.id === "string" && typeof o.studentNumber === "string" && typeof o.handleName == "string";
  }

  const user = JSON.parse(userStr) as User;
  return isValidUser(user) ? user : null;
}
