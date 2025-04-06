"use client";

import React, { useState } from "react";
import LoginModal from "./LoginModal";
import gameButton from "@/assets/images/home/game.png";
import { getCurrentUser } from "@/app/actions";
import { useRouter } from "next/navigation";
import styles from "@/assets/sass/molecules/GameStartButton.module.scss";

const GameStartButton = () => {
  const router = useRouter();
  const [isOpen, setOpen] = useState(false);
  const onClick = async () => {
    const user = await getCurrentUser();
    if (user) {
      router.push("/game");
    } else {
      setOpen(true);
    }
  };

  return (
    <>
      <div className={styles.button} onClick={onClick} aria-label="start game">
        {/* eslint-disable-next-line @next/next/no-img-element */}
        <img src={gameButton.src} alt="" />
      </div>
      <LoginModal isOpen={isOpen} onClose={() => setOpen(false)} />
    </>
  );
};

export default GameStartButton;
