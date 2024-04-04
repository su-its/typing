import { UserCardPresenter } from "@/components/molecules/UserCard";
import { describe, expect, it } from "@jest/globals";
import { render, screen } from "@testing-library/react";
import { User } from "@/types/user";

describe("UserCard", () => {
  const mockUser: User = {
    handleName: "しずっぴー",
    studentNumber: "B1234567",
    id: "1",
  };

  it("renders UserCard", () => {
    const userCard = render(UserCardPresenter(mockUser));

    const avatar = screen.getByRole("img");
    expect(avatar).toBeInTheDocument();
  });

  it("renders UserCard with user data", () => {
    const userCard = render(UserCardPresenter(mockUser));

    const name = screen.getByText(/名前:/);
    const studentNumber = screen.getByText(/学籍番号:/);
    expect(name).toHaveTextContent(mockUser.handleName);
    expect(studentNumber).toHaveTextContent(mockUser.studentNumber);
  });

  it("renders UserCard with default data", () => {
    const userCard = render(UserCardPresenter(null));

    const name = screen.getByText(/名前:/);
    const studentNumber = screen.getByText(/学籍番号:/);
    expect(name).toHaveTextContent("ログインしていません");
    expect(studentNumber).toHaveTextContent("未ログイン");
  });
});
