interface showToastProps {
  title: string;
  description?: string;
  status?: "info" | "warning" | "success" | "error";
}

const showToast = ({ title, description = "", status }: showToastProps) => {
  const event = new CustomEvent("app-toast", {
    detail: {
      title,
      description,
      status,
    },
  });
  window.dispatchEvent(event);
};

export const showSuccessToast = (title: string, description?: string) => {
  showToast({ title, description, status: "success" });
};

export const showWarningToast = (title: string, description?: string) => {
  showToast({ title, description, status: "warning" });
};

export const showErrorToast = (title: string, description?: string) => {
  showToast({ title, description, status: "error" });
};

export const showInfoToast = (title: string, description?: string) => {
  showToast({ title, description, status: "info" });
};
