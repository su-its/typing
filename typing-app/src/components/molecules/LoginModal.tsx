"use client";

import React, { useActionState } from "react";
import { login } from "@/app/actions";
import { showWarningToast } from "@/utils/toast";
import styles from "@/assets/sass/molecules/LoginModal.module.scss";

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
  state: { error?: string };
  dispatchAction: (payload: FormData) => void;
  pending: boolean;
}

const LoginModalPresenter: React.FC<LoginModalProps> = ({ isOpen, onClose, state, dispatchAction, pending }) => {
  return (
    <>
      {isOpen && (
        <div className={styles.modal}>
          <div className={styles.overlay}></div>
          <div className={styles.content}>
            <form
              action={async (formData: FormData) => {
                await dispatchAction(formData);
                state.error && showWarningToast(state.error);
              }}
            >
              <div className={styles.header}>続けるにはログインが必要です</div>
              <div className={styles.body}>
                <input
                  required
                  name="student-number"
                  placeholder="学籍番号を入力してください"
                  pattern="[0-9A-Z]{8}"
                  title="学籍番号"
                  role="textbox"
                />
              </div>
              <div className={styles.footer}>
                <button className={styles.button} role="submit">
                  ログインして続行
                </button>
                <button className={styles.button} onClick={onClose}>
                  閉じる
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </>
  );
};

export { LoginModalPresenter };

interface LoginModalContainerProps {
  isOpen: boolean;
  onClose: () => void;
}

const LoginModalContainer: React.FC<LoginModalContainerProps> = ({ isOpen, onClose }) => {
  const [state, dispatchAction, pending] = useActionState(login, {});

  return (
    <LoginModalPresenter
      isOpen={isOpen}
      onClose={onClose}
      state={state}
      dispatchAction={dispatchAction}
      pending={pending}
    />
  );
};

export default LoginModalContainer;
