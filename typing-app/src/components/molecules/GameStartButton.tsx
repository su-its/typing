"use client";

import React from "react";
import { Image, useDisclosure } from "@chakra-ui/react";
import LoginModal from "./LoginModal";
import gameButton from "@/assets/images/home/game.png";

const GameStartButton = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Image mb={2} src={gameButton.src} onClick={onOpen} cursor="pointer" />
      <LoginModal isOpen={isOpen} onClose={onClose} />
    </>
  );
};

export default GameStartButton;
