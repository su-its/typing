"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { useDisclosure } from "@chakra-ui/react";
import LogoutModal from "./LogoutModal";
import { logout } from "@/app/actions";
import logoutButton from "@/assets/images/home/logout.png";
import styles from "@/assets/sass/molecules/LogoutButton.module.scss";

const LogoutButton: React.FC = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const router = useRouter();

  const handleLogout = async () => {
    await logout();
    onOpen();
  };

  //Note: ログアウト時にページをリレンダリングするためにリダイレクトする
  const reLoad = () => {
    onClose();
    router.push("/");
  };

  return (
    <>
      <div className={styles.button} onClick={handleLogout}>
        <img src={logoutButton.src} />
      </div>
      <LogoutModal isOpen={isOpen} onClose={reLoad} />
    </>
  );
};

export default LogoutButton;
