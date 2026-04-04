"use client";
import { useEffect, useState } from "react";
import styles from "@/assets/sass/organism/Toast.module.scss";
import infoImage from "@/assets/images/toast/info.svg";
import warningImage from "@/assets/images/toast/warning.svg";
import successImage from "@/assets/images/toast/success.svg";
import errorImage from "@/assets/images/toast/error.svg";

export type ToastStatus = "info" | "warning" | "success" | "error";

interface ToastData {
  title: string;
  status?: ToastStatus;
}

export const Toast = () => {
  const [toast, setToast] = useState<ToastData | null>(null);
  const [isClosing, setIsClosing] = useState(false);

  useEffect(() => {
    const handler = (event: CustomEvent<ToastData>) => {
      setIsClosing(false);
      setToast(event.detail);
      setTimeout(() => {
        setIsClosing(true);
        setTimeout(() => {
          setToast(null);
        }, 500);
      }, 3000);
    };
    window.addEventListener("app-toast", handler as EventListener);
    return () => {
      window.removeEventListener("app-toast", handler as EventListener);
    };
  }, []);

  if (!toast) return null;

  return (
    <div
      className={`
        ${styles.toast}
        ${styles[toast.status ?? "error"]}
        ${isClosing ? styles.hide : styles.show}
      `}
    >
      <div className={styles.image}>
        {toast.status === "info" && <img src={infoImage.src} alt="Info" />}
        {toast.status === "warning" && <img src={warningImage.src} alt="Warning" />}
        {toast.status === "success" && <img src={successImage.src} alt="Success" />}
        {toast.status === "error" && <img src={errorImage.src} alt="Error" />}
      </div>
      <div className={`${styles.title} ${styles[toast.status ?? "error"]}`}>
        <div className={styles.text}>{toast.title}</div>
      </div>
    </div>
  );
};

export default Toast;
