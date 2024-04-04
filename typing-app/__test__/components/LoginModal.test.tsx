import { LoginModalPresenter } from "@/components/molecules/LoginModal";
import { describe, expect, it } from "@jest/globals";
import { render, screen } from "@testing-library/react";

describe("LoginModal", () => {
  it("renders a modal", () => {
    const mockDispatchAction = () => {};
    const mockState = {};
    const mockPending = false;

    const tree = render(
      <LoginModalPresenter
        isOpen={true}
        onClose={() => {}}
        state={mockState}
        dispatchAction={mockDispatchAction}
        pending={mockPending}
      />
    );
    const textbox = screen.getByRole("textbox");
    const submitButton = screen.getByRole("submit");
    expect(textbox).toBeInTheDocument();
    expect(submitButton).toBeInTheDocument();
  });
});
