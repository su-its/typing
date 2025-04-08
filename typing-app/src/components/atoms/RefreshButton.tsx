import React from "react";
import styles from "@/assets/sass/atoms/RefreshButton.module.scss";

const RefreshIcon = () => (
  <svg
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="2"
    strokeLinecap="round"
    strokeLinejoin="round"
  >
    <polyline points="23 4 23 10 17 10"></polyline>
    <polyline points="1 20 1 14 7 14"></polyline>
    <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10"></path>
    <path d="M20.49 15a9 9 0 0 1-14.85 3.36L1 14"></path>
  </svg>
);

interface RefreshButtonProps {
  onClick: () => void;
  isDisabled?: boolean;
}

const RefreshButton: React.FC<RefreshButtonProps> = ({ onClick, isDisabled = false }) => {
  if (isDisabled) {
    return (
      <div className={`${styles.button} ${styles.disabled}`}>
        <RefreshIcon />
      </div>
    );
  } else {
    return (
      <div className={styles.button} onClick={onClick}>
        <RefreshIcon />
      </div>
    );
  }
};

export default RefreshButton;
