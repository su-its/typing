// components/MenuContainer.js
import { Flex, VStack } from "@chakra-ui/react";
import GameStartButton from "../atoms/GameStartButton";
import RankingButton from "../atoms/RankingButton";
import LogoutButton from "../atoms/LogoutButton";

const HomeMenuContainer = () => {
  return (
    <Flex justify="center" align="center" h="65vh">
      <VStack spacing={8} align="stretch" width="50%" maxWidth="md" mx="auto">
        <GameStartButton />
        <RankingButton />
        <LogoutButton />
      </VStack>
    </Flex>
  );
};

export default HomeMenuContainer;
