import { UserCardPresenter } from "@/components/molecules/UserCard";
import { describe, expect, it } from "@jest/globals";
import { render, screen } from "@testing-library/react";

describe("UserCard", () => {
  const mockUser = {
    handleName: "しずっぴー",
    studentNumber: "B1234567",
    id: "1",
  };

  it("renders UserCard", () => {
    const userCard = render(<UserCardPresenter user={mockUser} />);

    const avatar = screen.getByRole("img");
    expect(avatar).toBeInTheDocument();
  });

  it("renders UserCard with user data", () => {
    const userCard = render(<UserCardPresenter user={mockUser} />);

    const name = screen.getByText(/名前:/);
    const studentNumber = screen.getByText(/学籍番号:/);
    expect(name).toHaveTextContent(mockUser.handleName);
    expect(studentNumber).toHaveTextContent(mockUser.studentNumber);
  });

  it("renders UserCard with default data", () => {
    const userCard = render(<UserCardPresenter user={null} />);

    const name = screen.getByText(/名前:/);
    const studentNumber = screen.getByText(/学籍番号:/);
    expect(name).toHaveTextContent("ログインしていません");
    expect(studentNumber).toHaveTextContent("未ログイン");
  });
});
