"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button, useDisclosure } from "@chakra-ui/react";
import LoginModal from "./LoginModal"; // LoginModalコンポーネントをインポート

const GameStartButton = () => {
  const router = useRouter();
  const { isOpen, onOpen, onClose } = useDisclosure();

  const handleLogin = async (studentId: string) => {
    // 学籍番号を使用したログイン処理
    console.log(studentId); // 例: ログイン処理
    router.push("/game"); // ログイン成功後の遷移
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
