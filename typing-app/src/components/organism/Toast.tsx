"use client";
import { useEffect, useRef, useState } from "react";
import styles from "@/assets/sass/organism/Toast.module.scss";
import infoImage from "@/assets/images/toast/info.svg";
import warningImage from "@/assets/images/toast/warning.svg";
import successImage from "@/assets/images/toast/success.svg";
import errorImage from "@/assets/images/toast/error.svg";

const STATUS_LIST = ["info", "warning", "success", "error"] as const;
export type ToastStatus = (typeof STATUS_LIST)[number];

type ToastData = {
  title: string;
  status?: string;
};

const isValidStatus = (status: unknown): status is ToastStatus => {
  return typeof status === "string" && STATUS_LIST.includes(status as ToastStatus);
};

export const Toast = () => {
  const [toast, setToast] = useState<ToastData | null>(null);
  const [isClosing, setIsClosing] = useState(false);
  const closingTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const removeTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  useEffect(() => {
    const handler = (event: CustomEvent<ToastData>) => {
      if (closingTimerRef.current) clearTimeout(closingTimerRef.current);
      if (removeTimerRef.current) clearTimeout(removeTimerRef.current);
      setIsClosing(false);
      setToast(event.detail);

      closingTimerRef.current = setTimeout(() => {
        setIsClosing(true);
        removeTimerRef.current = setTimeout(() => {
          setToast(null);
        }, 500);
      }, 3000);
    };
    window.addEventListener("app-toast", handler as EventListener);
    return () => {
      window.removeEventListener("app-toast", handler as EventListener);
      if (closingTimerRef.current) clearTimeout(closingTimerRef.current);
      if (removeTimerRef.current) clearTimeout(removeTimerRef.current);
    };
  }, []);

  if (!toast) return null;

  const status: ToastStatus = isValidStatus(toast.status) ? toast.status : "info";

  return (
    <div
      className={`
        ${styles.toast}
        ${styles[status]}
        ${isClosing ? styles.hide : styles.show}
      `}
    >
      <div className={styles.image}>
        {status === "info" && <img src={infoImage.src} alt="Info" />}
        {status === "warning" && <img src={warningImage.src} alt="Warning" />}
        {status === "success" && <img src={successImage.src} alt="Success" />}
        {status === "error" && <img src={errorImage.src} alt="Error" />}
      </div>
      <div className={`${styles.title} ${styles[status]}`}>
        <div className={styles.text}>{toast.title}</div>
      </div>
    </div>
  );
};

export default Toast;
