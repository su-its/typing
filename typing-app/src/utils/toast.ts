interface showToastProps {
  title: string;
  status?: "info" | "warning" | "success" | "error";
}

const showToast = ({ title, status }: showToastProps) => {
  const event = new CustomEvent("app-toast", {
    detail: {
      title,
      status,
    },
  });
  window.dispatchEvent(event);
};

export const showSuccessToast = (title: string) => {
  showToast({ title, status: "success" });
};

export const showWarningToast = (title: string) => {
  showToast({ title, status: "warning" });
};

export const showErrorToast = (title: string) => {
  showToast({ title, status: "error" });
};

export const showInfoToast = (title: string) => {
  showToast({ title, status: "info" });
};
