import styles from "@/assets/sass/atoms/CustomButton.module.scss";

interface ButtonProps {
  onClick: () => void;
  isDisabled: boolean;
  children: React.ReactNode;
}

export const CustomButton = ({ onClick, isDisabled, children }: ButtonProps) => {
  if (isDisabled) {
    return <div className={`${styles.button} ${styles.disabled}`}>{children}</div>;
  } else {
    return (
      <div className={styles.button} onClick={onClick}>
        {children}
      </div>
    );
  }
};
