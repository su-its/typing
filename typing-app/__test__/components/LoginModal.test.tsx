import { LoginModalPresenter } from "@/components/molecules/LoginModal";
import { describe, expect, it } from "@jest/globals";
import { render, screen, fireEvent } from "@testing-library/react";

describe("LoginModal", () => {
  const mockDispatchAction = jest.fn();;
  const mockState = {};
  const mockPending = false;

  it("renders input form in LoginModal", () => {
    const loginModal = render(
      <LoginModalPresenter
        isOpen={true}
        onClose={() => {}}
        state={mockState}
        dispatchAction={mockDispatchAction}
        pending={mockPending}
      />
    );

    const textbox = screen.getByRole("textbox");
    expect(textbox).toBeInTheDocument();
  });

  it("renders button in LoginModal", () => {
    const loginModal = render(
      <LoginModalPresenter
        isOpen={true}
        onClose={() => {}}
        state={mockState}
        dispatchAction={mockDispatchAction}
        pending={mockPending}
      />
    );

    const submitButton = screen.getByRole("submit");
    expect(submitButton).toBeInTheDocument();
  });

  it("Do not calls dispatchAction when submit button is clicked", () => {
    const loginModal = render(
      <LoginModalPresenter
        isOpen={true}
        onClose={() => {}}
        state={mockState}
        dispatchAction={mockDispatchAction}
        pending={mockPending}
      />
    );

    const submitButton = screen.getByRole("submit");
    fireEvent.click(submitButton);
    expect(mockDispatchAction).not.toHaveBeenCalled();
  });

  it("Do not calls dispatchAction when user input is not valid", () => {
    const loginModal = render(
      <LoginModalPresenter
        isOpen={true}
        onClose={() => {}}
        state={mockState}
        dispatchAction={mockDispatchAction}
        pending={mockPending}
      />
    );

    const textbox = screen.getByRole("textbox");
    fireEvent.change(textbox, { target: { value: "1" } });
    fireEvent.click(screen.getByRole("submit"));
    expect(mockDispatchAction).not.toHaveBeenCalled();
  });
});
