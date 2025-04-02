"use client";

import React from "react";
import { useDisclosure } from "@chakra-ui/react";
import LoginModal from "./LoginModal";
import gameButton from "@/assets/images/home/game.png";
import { getCurrentUser } from "@/app/actions";
import { useRouter } from "next/navigation";
import styles from "@/assets/sass/molecules/GameStartButton.module.scss";

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
  };

  return (
    <>
      <img className={styles.button} src={gameButton.src} onClick={onClick} />
      <LoginModal isOpen={isOpen} onClose={onClose} />
    </>
  );
};

export default GameStartButton;
