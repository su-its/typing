import { Button } from "@chakra-ui/react";

interface ButtonProps {
  onClick: () => void;
  isDisabled: boolean;
  children: React.ReactNode;
}

export const CustomButton = ({ onClick, isDisabled, children }: ButtonProps) => {
  return (
    <Button onClick={onClick} isDisabled={isDisabled}>
      {children}
    </Button>
  );
};
