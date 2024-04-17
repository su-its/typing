"use client";

import React from "react";
import { Image, useDisclosure } from "@chakra-ui/react";
import LoginModal from "./LoginModal";
import gameButton from "@/assets/images/home/game.png";
import { getCurrentUser } from "@/app/actions";
import { useRouter } from "next/navigation";


const GameStartButton = () => {
  const router = useRouter();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const onClick = async () => {
    const user = await getCurrentUser();
    if (user) {
      router.push("/game");
    } else {
      onOpen();
    }
  }

  return (
    <>
      <Image mb={2} src={gameButton.src} onClick={onClick} cursor="pointer" />
      <LoginModal isOpen={isOpen} onClose={onClose} />
    </>
  );
};

export default GameStartButton;
