"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import LogoutModal from "./LogoutModal";
import { logout } from "@/app/actions";
import logoutButton from "@/assets/images/home/logout.png";
import styles from "@/assets/sass/molecules/LogoutButton.module.scss";

const LogoutButton: React.FC = () => {
  const router = useRouter();
  const [isOpen, setIsOpen] = useState(false);

  const handleLogout = async () => {
    await logout();
    setIsOpen(true);
  };

  //Note: ログアウト時にページをリレンダリングするためにリダイレクトする
  const reLoad = () => {
    setIsOpen(false);
    router.push("/");
  };

  return (
    <>
      <div className={styles.button} onClick={handleLogout} aria-label="Logout">
        {/* eslint-disable-next-line @next/next/no-img-element */}
        <img src={logoutButton.src} alt="" />
      </div>
      <LogoutModal isOpen={isOpen} onClose={reLoad} />
    </>
  );
};

export default LogoutButton;
