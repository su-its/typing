import { createStandaloneToast } from "@chakra-ui/react";

const { toast } = createStandaloneToast();

interface showToastProps {
  title: string;
  description?: string;
  status?: "info" | "warning" | "success" | "error";
}

const showToast = ({ title, description = "", status }: showToastProps) => {
  toast({
    position: "top",
    title,
    description,
    status,
  });
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
