"use client";

import { useRouter } from "next/navigation";1
import { Button } from "@chakra-ui/react";

const GameStartButton = () => {
  const router = useRouter();

  const handleRouteGame = async () => {
    //TODO:ログインの維持を実装
    router.push("/game");
  };

  return (
    <Button colorScheme="green" size="lg" onClick = {handleRouteGame}>
      Game Start
    </Button>
  );
};

export default GameStartButton;
