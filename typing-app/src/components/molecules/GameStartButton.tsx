"use client";

import React from "react";
import { Button, useDisclosure } from "@chakra-ui/react";
import LoginModal from "./LoginModal";

const GameStartButton = () => {
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
