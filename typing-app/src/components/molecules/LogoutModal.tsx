import React from "react";
import styles from "@/assets/sass/molecules/LogoutModal.module.scss";

interface LogoutModalProps {
  isOpen: boolean;
  onClose: () => void;
}

const LogoutModal: React.FC<LogoutModalProps> = ({ isOpen, onClose }) => {
  return (
    <>
      {isOpen && (
        <div className={styles.modal}>
          <div className={styles.overlay}></div>
          <div className={styles.content}>
            <div className={styles.header}>ログアウトしました</div>
            <div className={styles.body}>ご利用ありがとうございました。</div>
            <div className={styles.footer}>
              <button className={`${styles.button} ${styles.blue}`} onClick={onClose}>
                OK
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default LogoutModal;
