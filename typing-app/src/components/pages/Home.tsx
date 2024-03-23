import React from "react";
import { Text, VStack } from "@chakra-ui/react";
import GameStartButton from "../buttons/gameStartButton";
import RankingButton from "../buttons/rankingButton";
import LogoutButton from "../buttons/logoutButton";

const HomePage: React.FC = () => {
  return (
    <>
      <VStack>
        <Text fontSize="2xl">Hello, World!</Text>
        <Text fontSize="xl">Welcome to the ITS Room</Text>
        <GameStartButton />
        <RankingButton />
        <LogoutButton />
      </VStack>
    </>
  );
};

export default HomePage;
