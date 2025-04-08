import styles from "@/assets/sass/atoms/CustomButton.module.scss";

interface ButtonProps {
  onClick: () => void;
  isDisabled: boolean;
  children: React.ReactNode;
}

export const CustomButton = ({ onClick, isDisabled, children }: ButtonProps) => {
  return (
    <button className={styles.button} disabled={isDisabled} onClick={onClick}>
      {children}
    </button>
  );
};
