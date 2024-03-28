"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button, useDisclosure } from "@chakra-ui/react";
import LoginModal from "./LoginModal"; // LoginModalコンポーネントをインポート

const GameStartButton = () => {
  const router = useRouter();
  const { isOpen, onOpen, onClose } = useDisclosure();

  const handleLogin = async (studentId: string) => {
    // TODO:ログイン済みかどうかを判別
    // TODO:学籍番号を使用したログイン処理
    console.log(studentId);
    router.push("/game");
  };

  return (
    <>
      <Button colorScheme="green" size="lg" onClick={onOpen}>
        Game Start
      </Button>

      <LoginModal isOpen={isOpen} onClose={onClose} onLogin={handleLogin} />
    </>
  );
};

export default GameStartButton;
