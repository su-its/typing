"use client";

import React, { useActionState } from "react";
import { login } from "@/app/actions";
//import { showWarningToast } from "@/utils/toast"; // FIXME: showWarningToast がちゃんと出ないことがある
import styles from "@/assets/sass/molecules/LoginModal.module.scss";

interface LoginModalProps {
  isOpen: boolean;
  onClose: () => void;
  state: { error?: string };
  dispatchAction: (payload: FormData) => void;
  pending: boolean;
}

const LoginModalPresenter: React.FC<LoginModalProps> = ({ isOpen, onClose, state, dispatchAction }) => {
  return (
    <>
      {isOpen && (
        <div className={styles.modal}>
          <div className={styles.overlay}></div>
          <div className={styles.content}>
            <form
              action={(formData: FormData) => {
                return dispatchAction(formData);
              }}
            >
              <div className={styles.header}>続けるにはログインが必要です</div>
              <div className={styles.body}>
                <input
                  required
                  type="text"
                  name="student-number"
                  placeholder="学籍番号を入力してください"
                  pattern="[0-9A-Z]{8}"
                  title="学籍番号"
                  role="textbox"
                />
                {/* FIXME: 一度 state.error に値が入ると次にモーダルを開いたときに前の state.error が表示されてしまう */}
                {state.error && <span><sub>{state.error}</sub></span>}
              </div>
              <div className={styles.footer}>
                <button className={`${styles.button} ${styles.blue}`} role="submit">
                  ログインして続行
                </button>
                <button className={`${styles.button} ${styles.gray}`} onClick={onClose}>
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
