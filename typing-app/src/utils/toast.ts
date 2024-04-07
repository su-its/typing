import { createStandaloneToast } from "@chakra-ui/react";

const { toast } = createStandaloneToast();

const showToast = (title: string, description?: string, status?: "info" | "warning" | "success" | "error") => {
  toast({
    title,
    description,
    status,
    position: "top",
  });
};

export const showSuccessToast = (title: string, description?: string) => {
  showToast(title, description, "success");
};

export const showWarningToast = (title: string, description?: string) => {
  showToast(title, description, "warning");
};

export const showErrorToast = (title: string, description?: string) => {
  showToast(title, description, "error");
};
