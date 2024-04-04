"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button, useDisclosure } from "@chakra-ui/react";
import LogoutModal from "./LogoutModal"; 
import { logout } from "@/app/actions";

const LogoutButton: React.FC = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const router = useRouter();

  const handleLogout = async () => {
    await logout();
    onOpen();
  };

  //Note: ログアウト時にページをリレンダリングするためにリダイレクトする
  const handleRedirect = () => {
    onClose();
    router.push("/");
  }

  return (
    <>
      <Button colorScheme="blue" size="lg" onClick={handleLogout}>
        Logout
      </Button>

      <LogoutModal isOpen={isOpen} onClose={handleRedirect} />
    </>
  );
};

export default LogoutButton;
