"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Image, useDisclosure } from "@chakra-ui/react";
import LogoutModal from "./LogoutModal";
import { logout } from "@/app/actions";
import logoutButton from "@/assets/images/home/logout.png";

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
      <Image src={logoutButton.src} onClick={handleLogout} cursor="pointer" />
      <LogoutModal isOpen={isOpen} onClose={reLoad} />
    </>
  );
};

export default LogoutButton;
