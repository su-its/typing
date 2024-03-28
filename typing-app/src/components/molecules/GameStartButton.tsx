"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button, useDisclosure } from "@chakra-ui/react";
import LoginModal from "./LoginModal"; // LoginModalコンポーネントをインポート

const GameStartButton = () => {
  const router = useRouter();
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Button colorScheme="green" size="lg" onClick={onOpen}>
        Game Start
      </Button>

      <LoginModal isOpen={isOpen} onClose={onClose} />
    </>
  );
};

export default GameStartButton;
