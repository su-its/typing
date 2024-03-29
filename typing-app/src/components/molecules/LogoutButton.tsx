"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button, useDisclosure } from "@chakra-ui/react";
import LogoutModal from "./LogoutModal"; // 正しいパスに変更してください

const LogoutButton: React.FC = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const router = useRouter();

  const handleLogout = async () => {
    // TODO:ログアウト処理を実装
    onOpen();
  };

  return (
    <>
      <Button colorScheme="blue" size="lg" onClick={handleLogout}>
        Logout
      </Button>

      <LogoutModal isOpen={isOpen} onClose={onClose}/>
    </>
  );
};

export default LogoutButton;
